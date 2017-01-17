package lib

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"golang.org/x/oauth2"

	"time"

	"github.com/mozillazg/request"
)

// Empty .
const (
	Empty         = ""
	Authorization = "Authorization"
	Bearer        = "Bearer"
)

// HTTPClient struct.
type HTTPClient struct {
	req *request.Request
}

// httpClient instance
var httpClient *HTTPClient

// httpHook struct
type httpHook struct{}

// lycamToken
var lycamToken *oauth2.Token

// BeforeRequest added access_token
func (h *httpHook) BeforeRequest(req *http.Request) (resp *http.Response, err error) {

	if appKey == Empty || appSecret == Empty || password == Empty {
		log.Fatal("appKey, appSecret and password is required.")
	}

	var token string

	if lycamToken == nil || isExpire(lycamToken.Expiry) {

		lycamPlusOAuth2 := NewLycamPlusOAuth2(appKey, appSecret)
		lycamToken, err = lycamPlusOAuth2.OAuth(username, password)

		if err != nil {
			log.Fatal(err.Error())
		}
	}

	token = lycamToken.AccessToken

	req.Header.Set(Authorization, fmt.Sprintf("%s %s", Bearer, token))

	return
}

func (h *httpHook) AfterRequest(req *http.Request, resp *http.Response, err error) (newResp *http.Response, newErr error) {
	resp.Header.Set("Content-Encoding", "gzip")
	return
}

// NewHTTPClient create httpClient.
func NewHTTPClient() *HTTPClient {
	m := new(sync.Mutex)

	if httpClient == nil {
		m.Lock()
		c := new(http.Client)
		q := request.NewRequest(c)
		hook := &httpHook{}
		q.Hooks = []request.Hook{hook}
		httpClient = &HTTPClient{
			req: q,
		}
		m.Unlock()
	}

	return httpClient
}

// Get helper method.
func (that *HTTPClient) Get(path string) ([]byte, error) {
	resp, err := that.req.Get(path)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, NewStatusError("Get()", resp.Reason())
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, NewStatusError("Get()", err.Error())
	}

	return body, nil
}

// Post helper method.
func (that *HTTPClient) Post(path string, data map[string]string) ([]byte, error) {
	that.req.Data = data
	resp, err := that.req.Post(path)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, NewStatusError("Post()", resp.Reason())
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, NewStatusError("Post()", err.Error())
	}

	return body, nil
}

// Put helper method.
func (that *HTTPClient) Put(path string, data map[string]string) ([]byte, error) {
	that.req.Data = data
	resp, err := that.req.Put(path)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, NewStatusError("Put()", err.Error())
	}

	return body, nil
}

// Delete helper method.
func (that *HTTPClient) Delete(path string) ([]byte, error) {

	resp, err := that.req.Delete(path)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, NewStatusError("Delete()", err.Error())
	}

	return body, nil
}

// token Expire
func isExpire(date time.Time) bool {
	return time.Now().After(date)
}
