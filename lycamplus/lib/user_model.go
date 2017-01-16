package lib

// UserRequestModel struct.
type UserRequestModel struct {
	UserName    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
	DisplayName string `json:"displayName"`
}

// UserResponseModel struct.
type UserResponseModel struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	UUID     string `json:"uuid"`
	Success  bool   `json:"success"`
}

//
// TokenResponseModel
//

// tokenModel struct.
type tokenModel struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int32  `json:"expires_in"`
}

// TokenResponseModel struct.
type TokenResponseModel struct {
	Success bool       `json:"success"`
	Scope   string     `json:"scope"`
	Token   tokenModel `json:"token"`
}
