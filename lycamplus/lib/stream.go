package lib

import "fmt"
import "encoding/json"

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
func (that *Stream) Create(streamRequestModel StreamRequestModel) (*StreamResponseModel, error) {

	path := fmt.Sprintf("%s/%s/%s", DefaultAPIURL, DefaultAPIVersion, "streams")
	params, err := Struct2Map(streamRequestModel)

	if err != nil {
		return nil, err
	}

	data, err := that.client.Post(path, params)

	if err != nil {
		return nil, err
	}

	response := new(StreamResponseModel)

	err = json.Unmarshal(data, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Update Stream By StreamID.
func (that *Stream) Update(streamID string,
	StreamRequestModel StreamRequestModel) (*StreamResponseModel, error) {

	path := fmt.Sprintf("%s/%s/%s/%s", DefaultAPIURL, DefaultAPIVersion, "streams", streamID)
	params, err := Struct2Map(StreamRequestModel)
	if err != nil {
		return nil, err
	}

	data, err := that.client.Put(path, params)

	if err != nil {
		return nil, err
	}

	response := new(StreamResponseModel)
	err = json.Unmarshal(data, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Show  Stream by streamID.
func (that *Stream) Show(streamID string) (*StreamResponseModel, error) {
	path := fmt.Sprintf("%s/%s/%s/%s", DefaultAPIURL, DefaultAPIVersion, "streams", streamID)

	data, err := that.client.Get(path)
	if err != nil {
		return nil, err
	}

	response := new(StreamResponseModel)
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// List query all video stream.
func (that *Stream) List() (*StreamResponseModelList, error) {
	path := fmt.Sprintf("%s/%s/%s", DefaultAPIURL, DefaultAPIVersion, "streams")

	data, err := that.client.Get(path)

	if err != nil {
		return nil, err
	}

	responseList := new(StreamResponseModelList)

	err = json.Unmarshal(data, &responseList)
	if err != nil {
		return nil, err
	}

	return responseList, nil
}

// ListSince get video stream from timestamp.
func (that *Stream) ListSince(timestamp int64) (*StreamResponseModelList, error) {

	path := fmt.Sprintf("%s/%s/%s/%s/%d", DefaultAPIURL, DefaultAPIVersion,
		"streams", "since", timestamp/1e6)

	data, err := that.client.Get(path)
	if err != nil {
		return nil, err
	}

	responseList := new(StreamResponseModelList)
	err = json.Unmarshal(data, &responseList)
	if err != nil {
		return nil, err
	}
	return responseList, nil
}

// SearchByKeyword search video stream by keyword.
func (that *Stream) SearchByKeyword(keywordModel KeywordModel) (*StreamResponseModelList, error) {

	path := fmt.Sprintf("%s/%s/%s", DefaultAPIURL, DefaultAPIVersion, "search")
	params, err := Struct2Map(keywordModel)
	if err != nil {
		return nil, err
	}

	data, err := that.client.Post(path, params)

	if err != nil {
		return nil, err
	}

	response := new(StreamResponseModelList)
	err = json.Unmarshal(data, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// SearchByLocation search video stream by location.
func (that *Stream) SearchByLocation(locationModel LocationModel) (*StreamResponseModelList, error) {

	path := fmt.Sprintf("%s/%s/%s/%s", DefaultAPIURL, DefaultAPIVersion, "search", "location")

	params, err := Struct2Map(locationModel)

	if err != nil {
		return nil, err
	}

	data, err := that.client.Post(path, params)

	if err != nil {
		return nil, err
	}

	response := new(StreamResponseModelList)
	err = json.Unmarshal(data, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Delete destroy a video stream.
func (that *Stream) Delete(streamID string) (*SuccessResponseModel, error) {
	path := fmt.Sprintf("%s/%s/%s/%s", DefaultAPIURL, DefaultAPIVersion, "streams", streamID)

	data, err := that.client.Delete(path)

	if err != nil {
		return nil, err
	}

	response := new(SuccessResponseModel)
	err = json.Unmarshal(data, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
