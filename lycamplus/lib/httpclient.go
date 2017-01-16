package lib

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

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

// BeforeRequest added access_token
func (h *httpHook) BeforeRequest(req *http.Request) (resp *http.Response, err error) {

	appKey := appKey
	appSecret := appSecret
	username := DefaultUsername
	password := password

	if appKey == Empty || appSecret == Empty || password == Empty {
		log.Fatal("appKey, appSecret and password is required.")
	}

	lycamPlusOAuth2 := NewLycamPlusOAuth2(appKey, appSecret)
	token, err := lycamPlusOAuth2.OAuth(username, password)

	if err != nil {
		log.Fatal(err.Error())
	}

	req.Header.Set(Authorization, fmt.Sprintf("%s %s", Bearer, token.AccessToken))

	return
}

func (h *httpHook) AfterRequest(req *http.Request, resp *http.Response, err error) (newResp *http.Response, newErr error) {
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
func (that *HTTPClient) Get(path string) (string, error) {
	resp, err := that.req.Get(path)
	if err != nil {
		return Empty, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return Empty, NewStatusError("Get()", resp.Reason())
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return Empty, NewStatusError("Get()", err.Error())
	}

	return string(body), nil
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
func (that *HTTPClient) Put(path string, data map[string]string) (string, error) {
	that.req.Data = data
	resp, err := that.req.Put(path)

	if err != nil {
		return Empty, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return Empty, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return Empty, NewStatusError("Put()", err.Error())
	}

	return string(body), nil
}

// Delete helper method.
func (that *HTTPClient) Delete(path string) (string, error) {

	resp, err := that.req.Delete(path)

	if err != nil {
		return Empty, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return Empty, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return Empty, NewStatusError("Delete()", err.Error())
	}

	return string(body), nil
}
