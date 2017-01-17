package lycamplus

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/lycam-dev/lycamplus-go-sdk/lycamplus/lib"
)

// Stream struct.
type Stream struct {
	client *lib.HTTPClient
}

// NewStream function.
func NewStream() *Stream {
	return &Stream{
		client: lib.NewHTTPClient(),
	}
}

// Create Stream
func (that *Stream) Create(streamRequestModel *StreamRequestModel) (*StreamResponseModel, error) {

	path := fmt.Sprintf("%s/%s/%s", lib.DefaultAPIURL, lib.DefaultAPIVersion, "streams")
	params, err := lib.Struct2Map(streamRequestModel)

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
	StreamRequestModel *StreamRequestModel) (*StreamResponseModel, error) {

	path := fmt.Sprintf("%s/%s/%s/%s", lib.DefaultAPIURL, lib.DefaultAPIVersion, "streams", streamID)
	params, err := lib.Struct2Map(StreamRequestModel)
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
	path := fmt.Sprintf("%s/%s/%s/%s", lib.DefaultAPIURL, lib.DefaultAPIVersion, "streams", streamID)

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
func (that *Stream) List(pageModel *PageModel) (*StreamResponseModelList, error) {
	path := fmt.Sprintf("%s/%s/%s", lib.DefaultAPIURL, lib.DefaultAPIVersion, "streams")
	path = paramsUtil(path, pageModel)

	data, err := that.client.Get(path)

	if err != nil {
		return nil, err
	}

	responseList := new(StreamResponseModelList)

	err = json.Unmarshal(data[:len(data)], &responseList)

	if err != nil {
		return nil, err
	}

	return responseList, nil
}

// ListSince get video stream from timestamp.
func (that *Stream) ListSince(timestamp int64, pageModel *PageModel) (*StreamResponseModelList, error) {

	path := fmt.Sprintf("%s/%s/%s/%s/%d", lib.DefaultAPIURL, lib.DefaultAPIVersion,
		"streams", "since", timestamp/1e6)
	path = paramsUtil(path, pageModel)

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
func (that *Stream) SearchByKeyword(keywordModel *KeywordModel) (*StreamResponseModelList, error) {

	path := fmt.Sprintf("%s/%s/%s", lib.DefaultAPIURL, lib.DefaultAPIVersion, "search")
	params, err := lib.Struct2Map(keywordModel)
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
func (that *Stream) SearchByLocation(locationModel *LocationModel) (*StreamResponseModelList, error) {

	path := fmt.Sprintf("%s/%s/%s/%s", lib.DefaultAPIURL, lib.DefaultAPIVersion, "search", "location")

	params, err := lib.Struct2Map(locationModel)

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
	path := fmt.Sprintf("%s/%s/%s/%s", lib.DefaultAPIURL, lib.DefaultAPIVersion, "streams", streamID)

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

// params to path
func paramsUtil(path string, pageModel *PageModel) string {

	if pageModel == nil {
		return path
	}

	paramsSlice := []string{}

	if pageModel.ResultsPerPage != 0 {
		resultsPerPage := fmt.Sprintf("resultsPerPage=%d", pageModel.ResultsPerPage)
		paramsSlice = append(paramsSlice, resultsPerPage)
	}

	if pageModel.Page != 0 {
		page := fmt.Sprintf("page=%d", pageModel.Page)
		paramsSlice = append(paramsSlice, page)
	}

	if pageModel.Order != "" {
		order := fmt.Sprintf("page=%s", pageModel.Order)
		paramsSlice = append(paramsSlice, order)
	}

	if pageModel.Sort != "" {
		sort := fmt.Sprintf("sort=%s", pageModel.Sort)
		paramsSlice = append(paramsSlice, sort)
	}

	params := strings.Join(paramsSlice, "&")

	return fmt.Sprintf("%s?%s", path, params)
}
