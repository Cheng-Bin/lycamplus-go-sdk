package lib

import (
	"testing"
)

func init() {
	appKey = "488ITUGN1G"
	appSecret = "z1oyx55jNQEXeRUu1iltfINZegWuGx"
	password = "9O1MZJ5UJwnuZky3tUBiZFPAlDJNs2"
}

func TestCreateUser(t *testing.T) {
	userInstance := NewUser()
	userRequestModel := UserRequestModel{UserName: "zhangsan"}
	response, err := userInstance.Create(userRequestModel)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(response)
	}
}

func TestUserAssume(t *testing.T) {
	userInstance := NewUser()
	uuid := "50391b70-dbcc-11e6-be2c-7f822d4bffa1"
	response, err := userInstance.Assume(uuid)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(response.Token.AccessToken)
	}
}
