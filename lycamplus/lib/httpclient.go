package lib

import (
	"net/http"
	"net/url"
	"runtime"
	"time"
)

const (
	mediaType = "application/json"
)

var (
	userAgent = runtime.Version()
)

// RequestCompletionCallback defines the type of the request callback function.
type RequestCompletionCallback func(*http.Request, *http.Response)

// Rate contains the rate limit for the current client.
type Rate struct {
	Limit     int       `json:"limit"`
	Remaining int       `json:"remaining"`
	Reset     time.Time `json:"reset"`
}

// HTTPClient Object.
type HTTPClient struct {
	client *http.Client

	BaseURL *url.URL

	userAgent string

	Rate Rate

	onRequestCompleted RequestCompletionCallback
}
