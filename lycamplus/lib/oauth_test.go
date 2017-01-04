package lib

import "testing"

func TestOAuth2(t *testing.T) {
	appKey := "488ITUGN1G"
	appSecret := "z1oyx55jNQEXeRUu1iltfINZegWuGx"
	username := DefaultUsername
	password := "9O1MZJ5UJwnuZky3tUBiZFPAlDJNs2"

	lycamPlusOAuth2 := NewLycamPlusOAuth2(appKey, appSecret)
	token, err := lycamPlusOAuth2.OAuth(username, password)

	if err != nil {
		t.Fatalf("get token failed. the error is : %s", err.Error())
	}

	t.Log(token)
}
