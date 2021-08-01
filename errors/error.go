package errors

import (
	"errors"
	"fmt"
	"strings"
)

func New(args ...interface{}) error {
	return errors.New(strings.TrimSuffix(fmt.Sprintln(args...), "\n"))
}
func Stack(err error, args ...interface{}) error {
	return New(strings.TrimSuffix(fmt.Sprintln(args...), "\n"), " reason:", err.Error())
}
