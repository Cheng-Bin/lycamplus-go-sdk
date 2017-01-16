package lib

const (
	// DefaultOAuth2URL is default OAuth2 server address.
	DefaultOAuth2URL = "https://oauth.lycam.tv"

	// DefaultAPIURL is default api server address.
	DefaultAPIURL = "https://api.lycam.tv"

	// DefaultTokenPath is default api path.
	DefaultTokenPath = "/oauth2/token"

	// DefaultAPIVersion is api version.
	DefaultAPIVersion = "v1"

	// DefaultUsername master name.
	DefaultUsername = "master"
)

// variables
var username = DefaultUsername
var (
	appKey    string
	appSecret string
	password  string
)

// InitKey .
func InitKey(_appKey, _appSecret, _masterSecret string) {
	appKey = _appKey
	appSecret = _appSecret
	password = _masterSecret
}

// SetAppKey .
func SetAppKey(_appKey string) {
	appKey = _appKey
}

// SetAppSecret .
func SetAppSecret(_appSecret string) {
	appSecret = _appSecret
}

// SetMasterSecret .
func SetMasterSecret(_masterSecret string) {
	password = _masterSecret
}
