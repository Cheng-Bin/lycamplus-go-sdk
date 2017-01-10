package lib

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"runtime"
	"strconv"
	"time"

	"io"

	"github.com/lycam-dev/lycamplus-go-sdk/lycamplus/utils"
)

const (
	mediaType = "application/json"

	headerRateLimit     = "RateLimit-Limit"
	headerRateRemaining = "RateLimit-Remaining"
	headerRateReset     = "RateLimit-Reset"
)

var (
	userAgent = runtime.Version()
)

// ErrorResponse reports the error caused by an API request
type ErrorResponse struct {
	Response  *http.Response
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
}

// RequestCompletionCallback defines the type of the request callback function.
type RequestCompletionCallback func(*http.Request, *http.Response)

// Rate contains the rate limit for the current client.
type Rate struct {
	Limit     int       `json:"limit"`
	Remaining int       `json:"remaining"`
	Reset     time.Time `json:"reset"`
}

// HTTPClient struct.
type HTTPClient struct {
	client *http.Client

	BaseURL *url.URL

	userAgent string

	Rate Rate

	onRequestCompleted RequestCompletionCallback
}

// ClientOpt are options for New.
type ClientOpt func(*HTTPClient) error

// New returns a new HTTPClient instance.
func New(base string, insecureSkipVerify bool, opts ...ClientOpt) (*HTTPClient, error) {
	baseURL, err := url.Parse(base)
	if err != nil {
		return nil, err
	}

	c := NewHTTPClient(baseURL, insecureSkipVerify)

	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

// NewHTTPClient returns a HTTPClient
func NewHTTPClient(baseURL *url.URL, insecureSkipVerify bool) *HTTPClient {
	var client *http.Client
	if baseURL.Scheme == "https" {
		tlsConfig := &tls.Config{RootCAs: x509.NewCertPool(), InsecureSkipVerify: insecureSkipVerify}
		transport := &http.Transport{TLSClientConfig: tlsConfig}
		client = &http.Client{Transport: transport}
	} else {
		client = http.DefaultClient
	}

	return &HTTPClient{client: client, BaseURL: baseURL, userAgent: userAgent}
}

// SetBaseURL is a client option for setting the base URL.
func SetBaseURL(baseURL string) ClientOpt {
	return func(c *HTTPClient) error {
		u, err := url.Parse(baseURL)
		if err != nil {
			return err
		}

		c.BaseURL = u

		return nil
	}
}

// SetUserAgent is a client option for setting the user agent.
func SetUserAgent(agent string) ClientOpt {
	return func(c *HTTPClient) error {
		c.userAgent = fmt.Sprintf("%s+%s", userAgent, c.userAgent)
		return nil
	}
}

// RequestOpt are options for New.
type RequestOpt func(*http.Request) error

// NewRequest create
func (c *HTTPClient) NewRequest(method string, urlStr string, body interface{}, opts []RequestOpt) (*http.Request, error) {
	rel, err := url.Parse(urlStr)

	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}

	req.Header.Add("Content-Type", mediaType)
	req.Header.Add("Accept", mediaType)
	req.Header.Add("User-Agent", c.userAgent)

	return req, nil
}

// SetHeader option for setting the a headder in the request
func SetHeader(key string, value string) RequestOpt {
	return func(r *http.Request) error {
		r.Header.Add(key, value)
		return nil
	}
}

// SetBearerToken to Header
func SetBearerToken(token string) RequestOpt {
	return SetHeader("Authorization", fmt.Sprintf("Bearer %s", token))
}

// setBasicAuthentication to Headder.
func setBasicAuthentication(username string, password string) RequestOpt {
	return func(r *http.Request) error {
		r.SetBasicAuth(username, password)
		return nil
	}
}

// OnRequestCompleted sets the API request completion callback .
func (c *HTTPClient) OnRequestCompleted(rc RequestCompletionCallback) {
	c.onRequestCompleted = rc
}

// Response of HTTP .
type Response struct {
	*http.Response
	Rate
}

// populateRate parses the rate related headers and populates the response Rate.
func (r *Response) populateRate() {
	if limit := r.Header.Get(headerRateLimit); limit != "" {
		r.Rate.Limit, _ = strconv.Atoi(limit)
	}

	if remaining := r.Header.Get(headerRateRemaining); remaining != "" {
		r.Rate.Limit, _ = strconv.Atoi(remaining)
	}

	if reset := r.Header.Get(headerRateReset); reset != "" {
		if v, _ := strconv.ParseInt(reset, 10, 64); v != 0 {
			r.Rate.Reset = utils.Timestamp{time.Unix(v, 0)}
		}
	}
}

// newResponse creates a new Response for the provided http.Response
func newResponse(r *http.Response) *Response {
	response := Response{Response: r}
	response.populateRate()
	return &response
}

// Do sends an API request and return the API response.
func (c *HTTPClient) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, resp)
	}

	defer func() {
		if rerr := resp.Body.Close(); err != nil {
			err = rerr
		}
	}()

	response := newResponse(resp)
	c.Rate = response.Rate

	err = CheckResponse(resp)
	if err != nil {
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err := io.Copy(w, resp.Body)
			if err != nil {
				return nil, err
			}
		} else {
			err := json.NewDecoder(resp.Body).Decode(v)
			if err != nil {
				return nil, err
			}
		}
	}

	return response, err
}

// CheckResponse checks the API response for errors
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	return nil
}
