package structtest

import "fmt"

type Student struct {
	Name  string
	Class int
	Num   int
}

type HonorStudent struct {
	Student Student
	Score   int
}

type EmbedHonorStudent struct {
	Student
	Score int
}

type EmbedHonorStudentVer2 struct {
	Student
	Name  string
	Score int
}

func CopyStruct() {

	d := EmbedHonorStudentVer2{Student{Name: "haha", Class: 1, Num: 32}, "HonoredHaha", 10000}

	c := d
	c.Name = "changed haha"
	fmt.Println(c)
	fmt.Println(d)

}

type SizeWasteStruct struct {
	a int8 // 1byte
	b int  // 8 byte
	c int8
	d int
}
type SizeCompactStruct struct {
	a int8 // 1byte
	b int8 // 8 byte
	c int
	d int
}
