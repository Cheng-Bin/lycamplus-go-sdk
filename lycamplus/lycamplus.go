package lycamplus

import "github.com/lycam-dev/lycamplus-go-sdk/lycamplus/lib"

// LycamPlus struct.
type LycamPlus struct {
	UserInstance   *User
	StreamInstance *Stream
}

// NewLycamPlus instance.
func NewLycamPlus(appKey, appSecret, masterSecret string) *LycamPlus {
	lib.InitKey(appKey, appSecret, masterSecret)
	return &LycamPlus{
		UserInstance:   NewUser(),
		StreamInstance: NewStream(),
	}
}
