package lib

import (
	"testing"
)

func TestStruct2Map(t *testing.T) {
	streamRequest := StreamRequest{}
	streamRequest.Title = "zhangsan"
	streamRequest.ExtraInfo = map[string]interface{}{
		"height": "170",
		"weight": "60",
	}

	Struct2Map(streamRequest)
}
