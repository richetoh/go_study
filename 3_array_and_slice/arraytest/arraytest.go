package arraytest

import "fmt"

func CreateArray() [5]int16 {
	// var intArray [5]int16 = [...]int16{1, 2, 3, 4, 5}
	// var intArray [5]int16 = [5]int16{1, 2, 3, 4, 5}
	var intArray = [5]int16{1, 2, 3, 4, 5}

	return intArray

}

func printArray(a [5]int16) {

	fmt.Println("PrintArraySlice")
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

}

func processArray(a [5]int16) {
	fmt.Println("ProcessArraySlice")
	for i := 0; i < len(a); i++ {
		a[i] = 999
	}

	printArray(a)
}

func CopyArray() {
	a := CreateArray()
	b := [5]int16{1, 2, 3, 4, 5}

	b = a
	b[4] = 9999
	fmt.Println("a ", a)
	fmt.Println("b ", b)
}

func TestArray() {
	a := CreateArray()
	printArray(a)
	processArray(a)
	printArray(a)
}

//MultiArray

func CreateMultiArray() [3][2]int16 {

	var intArray = [3][2]int16{{1, 2}, {1, 2}, {1, 3}}

	return intArray

}

func printMultiArray(a [3][2]int16) {
	fmt.Println("Print Multi Array")
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print(a[i][j], " ")
		}
		fmt.Println()
	}
}

func TestMultiArray() {
	a := CreateMultiArray()
	printMultiArray(a)
}
