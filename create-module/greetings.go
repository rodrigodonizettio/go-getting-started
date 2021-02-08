package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//   FuncName(pars)       returnTypes {
func Hail(name string) (string, error) {
	// Return a greeting that embeds the name in a message.
	if name == "" {
		return "", errors.New("Empty Name")
	}

	// var message string (":=" below does the same. It declares and initialize a variable in one line. The value in right determines the "type")
	return message, nil
}
