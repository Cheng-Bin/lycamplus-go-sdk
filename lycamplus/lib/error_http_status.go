package lib

import (
	"fmt"
)

// StatusError define.
type StatusError struct {
	fun    string
	reason string
}

// Error override
func (that *StatusError) Error() string {
	return fmt.Sprintf("function : %s, because %s", that.fun, that.reason)
}

// NewStatusError creates an Status code
func NewStatusError(fun string, reason string) *StatusError {
	return &StatusError{
		fun:    fun,
		reason: reason,
	}
}
