package utils

import (
	"strconv"
	"time"
)

// Timestamp struct
type Timestamp struct {
	time.Time
}

func (t Timestamp) String() string {
	return t.Time.String()
}

// UnmarshalJSON Unmarshal time
func (t *Timestamp) UnmarshalJSON(data []byte) error {
	str := string(data)
	i, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		t.Time = time.Unix(i, 0)
		return nil
	}

	return err
}

// Equal reports whether t and u are equal based on time.Equal.
func (t Timestamp) Equal(u Timestamp) bool {
	return t.Time.Equal(u.Time)
}
