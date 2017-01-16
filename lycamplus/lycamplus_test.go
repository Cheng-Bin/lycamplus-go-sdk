package lycamplus

import (
	"testing"
)

func TestLycamPlus(t *testing.T) {

	appKey := "488ITUGN1G"
	appSecret := "z1oyx55jNQEXeRUu1iltfINZegWuGx"
	masterSecret := "9O1MZJ5UJwnuZky3tUBiZFPAlDJNs2"

	lycamPlus := NewLycamPlus(appKey, appSecret, masterSecret)

	t.Log(lycamPlus.StreamInstance)

	t.Log(lycamPlus.UserInstance)
}
