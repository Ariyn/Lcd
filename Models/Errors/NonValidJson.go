package Errors

import (
	"fmt"
	"reflect"
)

type InvalidJson struct {
	Type   reflect.Type
	Name   string
	RawErr error
}

func (err InvalidJson) Error() string {
	return fmt.Sprintf("string is not valid %s json", err.Type.String())
}
