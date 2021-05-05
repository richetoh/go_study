package greetings

import "fmt"

//func Hello returns a greeting for the named person
//func {Function name} ({variable} {variable type}) {return type}
func Hello(name string) string {
	//:= operator is a shortcut for declaring and intializing variable in one line
	//var messages string
	//message = fmt.Sprintf("Hi, %v Welcome!", name)
	//you can skip "var messages string" part using :=

	message := fmt.Sprintf("Hi, %v Welcome!", name)
	return message
}
