package lib

import "fmt"

// Stream struct.
type Stream struct {
	client *HTTPClient
}

// NewStream function.
func NewStream() *Stream {
	return &Stream{
		client: new(HTTPClient),
	}
}

// Create Stream
func (that *Stream) Create(stream StreamRequest) (*StreamResponse, error) {

	path := fmt.Sprintf("%s/%s/%s", DefaultAPIURL, DefaultAPIVersion, "streams")
	params := Struct2Map(stream)

	_, err := that.client.Post(path, params)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
