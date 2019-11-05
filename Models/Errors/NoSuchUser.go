package Errors

import (
	"fmt"
)

type NoSuchUser struct {
	ID     string
	Name   string
	RawErr error
}

func (err NoSuchUser) Error() string {
	return fmt.Sprintf("no such user %s", err.ID)
}
