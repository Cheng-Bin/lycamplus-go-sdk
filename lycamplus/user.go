package lycamplus

import (
	"encoding/json"
	"fmt"

	"github.com/lycam-dev/lycamplus-go-sdk/lycamplus/lib"
)

// User struct define.
type User struct {
	client *lib.HTTPClient
}

// NewUser .
func NewUser() *User {
	return &User{client: lib.NewHTTPClient()}
}

// Create method.
func (u *User) Create(userRequestModel *UserRequestModel) (*UserResponseModel, error) {

	path := fmt.Sprintf("%s/%s/%s", lib.DefaultAPIURL, lib.DefaultAPIVersion, "users")

	params, err := lib.Struct2Map(userRequestModel)
	if err != nil {
		return nil, err
	}

	data, err := u.client.Post(path, params)
	if err != nil {
		return nil, err
	}

	response := new(UserResponseModel)

	err = json.Unmarshal(data, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Assume method.
func (u *User) Assume(uuid string) (*TokenResponseModel, error) {
	path := fmt.Sprintf("%s/%s/%s/%s/%s", lib.DefaultAPIURL, lib.DefaultAPIVersion, "users", uuid, "assume")

	data, err := u.client.Post(path, nil)

	if err != nil {
		return nil, err
	}

	response := new(TokenResponseModel)
	err = json.Unmarshal(data, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
