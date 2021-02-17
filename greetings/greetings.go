package greetings

import "fmt"

func Hello(name string) string {
	// Return a greeting that embedes the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
