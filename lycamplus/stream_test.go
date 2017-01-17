package lycamplus

import (
	"testing"
	"time"

	"github.com/lycam-dev/lycamplus-go-sdk/lycamplus/lib"
)

func init() {
	lib.InitKey("488ITUGN1G", "z1oyx55jNQEXeRUu1iltfINZegWuGx", "9O1MZJ5UJwnuZky3tUBiZFPAlDJNs2")
}

func TestStreamCreate(t *testing.T) {
	streamInstance := NewStream()
	requestModel := StreamRequestModel{
		Title:    "perfect",
		Privacy:  true,
		StartLat: 90.6,
		StartLon: 90.6,
		EndLat:   90.6,
		EndLon:   90.6,
	}

	response, err := streamInstance.Create(&requestModel)

	if err != nil {
		t.Error(err)
	} else {
		t.Log(response)
		t.Logf("streamId=%s, streamTitle=%s, privacy=%t",
			response.StreamID, response.Title, response.Privacy)
	}
}

func TestStreamUpdate(t *testing.T) {
	streamID := "2c0628e0-dbba-11e6-ac84-eb4c3b5430ae"

	streamInstance := NewStream()
	requestModel := StreamRequestModel{
		Title:   "666666",
		Privacy: false,
	}

	response, err := streamInstance.Update(streamID, &requestModel)

	if err != nil {
		t.Error(err)
	} else {
		t.Logf("streamId=%s, streamTitle=%s, privacy=%t",
			response.StreamID, response.Title, response.Privacy)
	}
}

func TestStreamShow(t *testing.T) {
	streamID := "2c0628e0-dbba-11e6-ac84-eb4c3b5430ae"
	streamInstance := NewStream()

	response, err := streamInstance.Show(streamID)

	if err != nil {
		t.Error(err)
	} else {
		t.Logf("streamId=%s, streamTitle=%s, privacy=%t",
			response.StreamID, response.Title, response.Privacy)
	}
}

func TestStreamList(t *testing.T) {
	streamInstance := NewStream()
	pageModel := PageModel{
		ResultsPerPage: 2,
	}
	response, err := streamInstance.List(&pageModel)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("streamSize=%d\n", len(response.Items))
		t.Log(response)
	}
}

func TestStreamListSince(t *testing.T) {

	streamInstance := NewStream()
	pageModel := PageModel{
		ResultsPerPage: 2,
	}

	response, err := streamInstance.ListSince(time.Now().UnixNano(), &pageModel)

	if err != nil {
		t.Error(err)
	} else {
		t.Logf("streamSize=%d\n", len(response.Items))
	}
}

func TestStreamSearchByKeyword(t *testing.T) {
	streamInstance := NewStream()
	keywordModel := KeywordModel{Keyword: "lycamplus"}
	response, err := streamInstance.SearchByKeyword(&keywordModel)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("streamSize=%d\n", len(response.Items))
		t.Log(response.Items)
	}
}

func TestStreamSearchByLocation(t *testing.T) {
	streamInstance := NewStream()
	locationModel := LocationModel{Lon: 90.6, Lat: 90.6, Radius: 1000}
	response, err := streamInstance.SearchByLocation(&locationModel)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("streamSize=%d\n", len(response.Items))
		t.Log(response.Items)
	}
}

func TestStreamDelete(t *testing.T) {
	streamID := "2c0628e0-dbba-11e6-ac84-eb4c3b5430ae"
	streamInstance := NewStream()
	resposne, err := streamInstance.Delete(streamID)

	if err != nil {
		t.Error(err)
	} else {
		t.Log(resposne.Success)
	}
}
