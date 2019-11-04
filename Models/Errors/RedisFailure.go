package Errors

import (
	"fmt"
)

type RedisFailure struct {
	ID     string
	Name   string
	RawErr error
}

func (err RedisFailure) Error() string {
	return fmt.Sprintf("redis failed %s", err.RawErr)
}
