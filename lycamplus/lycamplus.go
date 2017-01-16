package lycamplus

import "github.com/lycam-dev/lycamplus-go-sdk/lycamplus/lib"

// LycamPlus struct.
type LycamPlus struct {
	UserInstance   *lib.User
	StreamInstance *lib.Stream
}

// NewLycamPlus  instance.
func NewLycamPlus(appKey, appSecret, masterSecret string) *LycamPlus {
	lib.InitKey(appKey, appSecret, masterSecret)
	return &LycamPlus{
		UserInstance:   lib.NewUser(),
		StreamInstance: lib.NewStream(),
	}
}
