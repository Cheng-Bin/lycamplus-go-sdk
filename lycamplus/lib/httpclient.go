package lib

import (
	"net/http"
	"sync"

	"github.com/mozillazg/request"
)

// HTTPClient struct.
type HTTPClient struct {
	request request
}

// httpClient
var httpClient *HTTPClient

// NewHTTPClient create httpClient.
func NewHTTPClient() *HTTPClient {
	m := new(sync.Mutex)

	if httpClient != nil {
		m.Lock()
		c := new(http.Client)
		req := request.NewRequest(c)
		req.Headers = map[string]string{
            "Authorization": "Bearer "
        }
        httpClient = &HTTPClient{
            request: req
        }
		m.Unlock()
	}

	return httpClient
}

// Get helper function.
func (that *httpClient) Get(path string) (string, error) {
    resp, err := that.request.Get(path)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()
    data, err := resp.Json()
    if (err != nil) {
        return "", err
    }

    return data, err
}