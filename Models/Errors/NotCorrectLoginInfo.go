package Errors

import (
	"fmt"
)

type NotCorrectLoginInfo struct {
	ID     string
	Name   string
	RawErr error
}

func (err NotCorrectLoginInfo) Error() string {
	return fmt.Sprintf("not correct login info %s", err.ID)
}
