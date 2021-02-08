package greetings

import (
	"errors"
	"fmt"
)

//   FuncName(pars)       returnTypes {
func Hail(name string) (string, error) {
	// Return a greeting that embeds the name in a message.
	if name == "" {
		return "", errors.New("Empty Name")
	}

	// var message string (":=" below does the same. It declares and initialize a variable in one line. The value in right determines the "type")
	message := fmt.Sprintf("Hail to the king, %v!", name)
	fmt.Println(message)
	return message, nil
}
