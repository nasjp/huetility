package main

import (
	"errors"
	"fmt"
)

var (
	ErrNoCmd          = errors.New("no cmd")
	ErrIDNotSpecified = errors.New("set cmd require id")
)

type ErrIDNotFound string

func (e ErrIDNotFound) Error() string {
	return fmt.Sprintf("id: %s not found", string(e))
}

type ErrPutHueState string

func (e ErrPutHueState) Error() string {
	return fmt.Sprintf("error occured on set state: %s", string(e))
}
