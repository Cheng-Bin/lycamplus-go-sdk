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
