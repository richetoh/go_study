package main

import (
	"fmt"
	"github.com/getveryrichet/go_study/function/variadic"
)

func main() {

	fmt.Println(variadic.Sum(1,2,3,4,5))

	fmt.Println(variadic.Sum(10, 20))

	fmt.Println(variadic.Sum())

	variadic.Print(1, "2323", 3)
}
