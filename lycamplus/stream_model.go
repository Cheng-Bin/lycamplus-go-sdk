package lycamplus

//
// request
//
// StreamRequest struct.
type StreamRequestModel struct {
	UUID         string  `json:"user"`
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	ThumbnailURL string  `json:"thumbnailUrl"`
	StartLat     float32 `json:"startLat"`
	StartLon     float32 `json:"startLon"`
	EndLat       float32 `json:"endLat"`
	EndLon       float32 `json:"endLon"`
	Country      string  `json:"country"`
	State        string  `json:"state"`
	City         string  `json:"city"`
	Privacy      bool    `json:"privacy"`
}

//
// Response
//

// StreamURL struct.
type StreamURL struct {
	Type    string `json:"type"`
	URL     string `json:"url"`
	Bitrate int    `json:"bitrate"`
}

//StreamResponseModel struct.
type StreamResponseModel struct {
	StreamRequestModel
	StreamID         string      `json:"streamId"`
	Status           string      `json:"status"`
	StreamUrls       []StreamURL `json:"streamUrls"`
	ResourceURL      string      `json:"resourceUrl"`
	StreamType       string      `json:"streamType"`
	ChatURL          string      `json:"chatUrl"`
	ChatChannel      string      `json:"chatChannel"`
	ChatToken        string      `json:"chatToken"`
	UploadURL        string      `json:"uploadUrl"`
	VideoWidth       int         `json:"videoWidth"`
	VideoHeight      int         `json:"videoHeight"`
	VideoOrientation int         `json:"videoOrientation"`
	TimeStarted      string      `json:"timeStarted"`
	TimeFinished     string      `json:"timeFinished"`
}

// StreamResponseModelList struct.
type StreamResponseModelList struct {
	TotalItems        int                   `json:"totalItems"`
	ResultsPerPage    int                   `json:"resultsPerPage"`
	PageNumber        int                   `json:"pageNumber"`
	NextPageAvailable bool                  `json:"nextPageAvailable"`
	Items             []StreamResponseModel `json:"items"`
}

//
// KeywordModel
//

// PageModel struct.
type PageModel struct {
	ResultsPerPage int    `json:"resultsPerPage"`
	Page           int    `json:"page"`
	Sort           string `json:"sort"`
	Order          string `json:"order"`
}

//KeywordModel struct.
type KeywordModel struct {
	PageModel
	Keyword string `json:"keyword"`
}

//
// LocationModel
//

// LocationModel struct.
type LocationModel struct {
	PageModel
	Lon    float32 `json:"lon"`
	Lat    float32 `json:"lat"`
	Radius int     `json:"radius"`
}

//
// SuccessResponseModel
//

// SuccessResponseModel struct.
type SuccessResponseModel struct {
	Success bool `json:"success"`
}
