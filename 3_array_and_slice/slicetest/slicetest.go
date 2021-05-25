package slicetest

import "fmt"

func CreateArraySlice() []int16 {
	// var intArray [5]int16 = [...]int16{1, 2, 3, 4, 5}
	// var intArray [5]int16 = [5]int16{1, 2, 3, 4, 5}
	var intArray = []int16{1, 2, 3, 4, 5}

	return intArray

}
func CreateArraySlice2() []int16 {
	// var intArray [5]int16 = [...]int16{1, 2, 3, 4, 5}
	// var intArray [5]int16 = [5]int16{1, 2, 3, 4, 5}
	var intSlice = make([]int16, 3)
	intSlice[1] = 0
	return intSlice
}

func printArraySlice(a []int16) {
	fmt.Println("PrintArraySlice")
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}

func InsertValueToSlice(a []int16) {
	fmt.Println("InsertValueToSlice")
	for i := 0; i < len(a); i++ {
		a[i] = 999
	}
}
func TraverseSlice(a []int16) {
	fmt.Println("ProcessArraySlice ver1")
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	fmt.Println("ProcessArraySlice ver2")
	for i, v := range a {
		fmt.Println(i, v)
	}
}

func AppendValueToSlice(a []int16) {
	//append returns new slice
	a = append(a, 4)
}

func CopySlice() {
	a := CreateArraySlice()
	b := []int16{1, 2, 3, 4, 5}

	b = a
	b[4] = 9999
	fmt.Println("a ", a)
	fmt.Println("b ", b)

}

func TestSlice() {
	a := CreateArraySlice()
	printArraySlice(a)
	InsertValueToSlice(a)
	printArraySlice(a)
}

// Multi Slice Section

func CreateMultiArraySlice() [][]int16 {
	// var intArray [5]int16 = [...]int16{1, 2, 3, 4, 5}
	// var intArray [5]int16 = [5]int16{1, 2, 3, 4, 5}
	var intArray = [][]int16{{1, 2}, {1, 2}, {1, 3}}

	return intArray

}
func printMultiArraySlice(a [][]int16) {
	fmt.Println("Print Multi Array Slice")
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			fmt.Print(a[i][j], " ")
		}
		fmt.Println()
	}
}

func TestMultiArraySlice() {
	a := CreateMultiArraySlice()
	printMultiArraySlice(a)
}
