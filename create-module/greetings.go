package greetings

import "fmt"

//   FuncName(pars)       returnType {
func Hail(name string) string {
	// Return a greeting that embeds the name in a message.

	// var message string (":=" below does the same. It declares and initialize a variable in one line. The value in right determines the "type")
	message := fmt.Sprintf("Hail to the king, %v!", name)
	fmt.Println(message)
	return message
}
