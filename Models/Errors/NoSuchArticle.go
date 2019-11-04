package Errors

import (
	"fmt"
)

type NoSuchArticle struct {
	ID     string
	Name   string
	RawErr error
}

func (err NoSuchArticle) Error() string {
	return fmt.Sprintf("no such article %s", err.ID)
}
