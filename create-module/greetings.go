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
	message := fmt.Sprintf(randomHail(), name)
	return message, nil
}

func Hails(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hail(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func randomHail() string {
	formats := []string{
		"Hail to the king, %v!",
		"What we do in life echoes in eternity, %v!",
		"Imagine where you will be, %v, and it will be so!",
		"If you find yourself alone, %v, riding in the green fields with the sun on your face, do not be troubled. For you are in Elysium, and you're already dead!",
	}

	return formats[rand.Intn(len(formats))]
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
