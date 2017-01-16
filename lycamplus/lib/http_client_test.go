package lib

import "testing"

func init() {
	appKey = "488ITUGN1G"
	appSecret = "z1oyx55jNQEXeRUu1iltfINZegWuGx"
	password = "9O1MZJ5UJwnuZky3tUBiZFPAlDJNs2"
}

func TestGet(t *testing.T) {

	client := NewHTTPClient()
	body, err := client.Get(DefaultAPIURL)
	if err != nil {
		t.Errorf("occured error: %s", err.Error())
	} else {
		t.Logf("the result is : %s", body)
	}
}

func TestPost(t *testing.T) {
	client := NewHTTPClient()

	params := map[string]string{"title": "lycamplus test"}

	body, err := client.Post(DefaultAPIURL+"/v1/streams", params)
	if err != nil {
		t.Errorf("occured error: %s", err.Error())
	} else {
		t.Logf("the result is : %s", body)
	}
}

func TestPut(t *testing.T) {
	client := NewHTTPClient()
	streamID := "fef2f3e0-d968-11e6-981c-1df634f0f8cf"
	params := map[string]string{
		"description": "666",
	}

	body, err := client.Put(DefaultAPIURL+"/v1/streams/"+streamID, params)
	if err != nil {
		t.Errorf("occured error: %s", err.Error())
	} else {
		t.Logf("the result is : %s", body)
	}
}

func TestDelete(t *testing.T) {
	client := NewHTTPClient()
	streamID := "fef2f3e0-d968-11e6-981c-1df634f0f8cf"
	body, err := client.Delete(DefaultAPIURL + "/v1/streams/" + streamID)
	if err != nil {
		t.Errorf("occured error: %s", err.Error())
	} else {
		t.Logf("the result is : %s", body)
	}
}
