package pointertest

import "fmt"

type Data struct {
	Value string
}

func AssignValueToPointer() {

	var a int = 500
	var p *int

	p = &a
	fmt.Printf("p의 주소값 %p\n", p)
	fmt.Printf("p의 주솟값이 가리키는 값 %d\n", *p)

}

func ChangeValueOfPinter() {

	var a int = 500
	var q *int

	q = &a
	fmt.Println("a---------")
	fmt.Printf("a의 값 %d \n", a)
	fmt.Printf("a의 주솟값 %p\n", a)
	fmt.Printf("&a의 주솟값 %p\n", &a)
	fmt.Printf("&a의 주솟값 in digit %d\n", &a)

	fmt.Println("q--------")
	fmt.Printf("q가 가리키는 주소값 %p\n", q)
	fmt.Printf("q가 가리키는 값 %d\n", q)
	fmt.Printf("*q가 가리키는 값 %d\n", *q)

	fmt.Println("------q의 값 변경!-----")
	fmt.Println("------q의 값 변경!-----")
	*q = 300
	fmt.Println("a---------")
	fmt.Printf("a의 값 %d \n", a)
	fmt.Printf("a의 주솟값 %p\n", a)
	fmt.Printf("&a의 주솟값 %p\n", &a)
	fmt.Printf("&a의 주솟값 in digit %d\n", &a)

	fmt.Println("q---------")
	fmt.Printf("q가 가리키는 주소값 %p\n", q)
	fmt.Printf("q가 가리키는 값 %d\n", q)
	fmt.Printf("*q가 가리키는 값 %d\n", *q)
}

func CompareAddress() {

	var a int = 10
	var b int = 20

	var p1 *int = &a
	var p2 *int = &a
	var p3 *int = &b

	fmt.Printf("p1 == p2, %v\n", p1 == p2)
	fmt.Printf("p2 == p3 %v\n ", p2 == p3)

}

func changeValue(arg *[5]int) {

	arg[2] = 999999

}

func printArray(arg *[5]int) {
	for i := 0; i < len(arg); i++ {
		fmt.Print(arg[i], " ")
	}
	fmt.Println()
}

func TestChangeValue() {

	var a [5]int

	printArray(&a)
	changeValue(&a)
	fmt.Println("after change")
	printArray(&a)

}

func TestChangeValueUsePointer() {
	var q *[5]int
	var a [5]int

	q = &a

	printArray(q)
	changeValue(q)
	fmt.Println("after change")
	printArray(q)

	fmt.Println("---watch a-----")
	printArray(&a)
	changeValue(&a)
	fmt.Println("after change")
	printArray(&a)

	fmt.Println(q)
	fmt.Println(a)
}

func PrintAddressOfString(a *string, b string) {

	fmt.Printf("%s's address = %x \n", *a, *a)
	fmt.Printf("%s's address = %x \n", b, b)
}

func PrintAddressOfStruct(a *Data, b Data) {

	//pointer's address
	fmt.Println("a's address = ", &a)

	//struct's pointer's address
	fmt.Printf("b's address = %p\n", &b)
}

func CompareAddressOfPointers() {
	// a := 100
	// b := 100

	// fmt.Println("a's address = ", &a, "b's address", &b)

	c := "haha"
	// d := "haha"
	// fmt.Println("c's address = ", &c, "b's address", &d)

	var p *string
	p = &c

	fmt.Printf("value in address where p points %v\n", *p)
	fmt.Printf("value in address where p points %v\n", p)
	fmt.Printf("value in address where p points %v\n", *&c)
	fmt.Printf("address where p points %p\n", p)
	fmt.Printf("c's addres %p\n", &c)
	fmt.Println("pointer p's address", &p, "c's address", &c)
}
