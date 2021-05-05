package main

import (
	"fmt"

	"github.com/richet/gotest/greetings"
)

func main() {
	message := greetings.Hello("Richet")
	fmt.Println(message)

}
