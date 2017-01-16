package lib

import "fmt"

// Stream struct.
type Stream struct {
	client *HTTPClient
}

// NewStream function.
func NewStream() *Stream {
	return &Stream{
		client: NewHTTPClient(),
	}
}

// Create Stream
func (that *Stream) Create(stream StreamRequest) (*StreamResponse, error) {

	path := fmt.Sprintf("%s/%s/%s", DefaultAPIURL, DefaultAPIVersion, "streams")
	params, err := Struct2Map(stream)

	if err != nil {
		return nil, err
	}

	data, err := that.client.Post(path, params)

	if err != nil {
		return nil, err
	}

	response := new(StreamResponse)

	err = AdanceUnmarshal(data, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
