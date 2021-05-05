package main

import (
	"fmt"
	"unsafe"

	"github.com/richet.oh/go_study/struct/structtest"
)

func main() {
	a := structtest.Student{Name: "haha", Class: 1, Num: 32}
	fmt.Println(a)

	b := structtest.HonorStudent{a, 100}
	//fmt.Println(b.Name)= Error
	fmt.Println(b.Student.Name)
	fmt.Println(b)

	// c := structtest.EmbedHonorStudent{Name: "haha", Class: 1, Num: 32, Score: 10000} // creates an error

	c := structtest.EmbedHonorStudent{structtest.Student{Name: "haha", Class: 1, Num: 32}, 10000} // equivalent to below
	// c := structtest.EmbedHonorStudent{Student: structtest.Student{Name: "haha", Class: 1, Num: 32}, Score: 10000}
	fmt.Println(c)
	fmt.Println(c.Name)

	d := structtest.EmbedHonorStudentVer2{structtest.Student{Name: "haha", Class: 1, Num: 32}, "HonoredHaha", 10000}
	fmt.Println(d.Name)
	fmt.Println(d.Student.Name)

	structtest.CopyStruct()

	//print address
	fmt.Printf("%p\n", &d)

	e := structtest.SizeWasteStruct{}
	f := structtest.SizeCompactStruct{}

	fmt.Println(unsafe.Sizeof(e))
	fmt.Println(unsafe.Sizeof(f))

}
