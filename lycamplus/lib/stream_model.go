package lib

//
// request
//
// StreamRequest struct.
type StreamRequest struct {
	UUID         string                 `json:"uuid"`
	Title        string                 `json:"title"`
	Description  string                 `json:"description"`
	ThumbnailURL string                 `json:"thumbnailUrl"`
	StartLat     float32                `json:"startLat"`
	StartLon     float32                `json:"startLon"`
	EndLat       float32                `json:"endLat"`
	EndLon       float32                `json:"endLon"`
	Country      string                 `json:"country"`
	State        string                 `json:"state"`
	City         string                 `json:"city"`
	Privacy      bool                   `json:"privacy"`
	ExtraInfo    map[string]interface{} `json:"extraInfo"`
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

//StreamResponse struct.
type StreamResponse struct {
	StreamRequest
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
