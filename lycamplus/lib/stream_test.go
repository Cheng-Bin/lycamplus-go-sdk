package lib

import (
	"fmt"
	"testing"
)

func init() {
	appKey = "488ITUGN1G"
	appSecret = "z1oyx55jNQEXeRUu1iltfINZegWuGx"
	password = "9O1MZJ5UJwnuZky3tUBiZFPAlDJNs2"
}

func TestStreamCreate(t *testing.T) {
	streamInstance := NewStream()
	extInfo := map[string]interface{}{
		"height": "180",
		"width":  "70",
	}
	requestModel := StreamRequest{
		Title:     "perfect",
		ExtraInfo: extInfo,
	}

	data, err := streamInstance.Create(requestModel)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
	}
}
