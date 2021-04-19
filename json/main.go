package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Page   int
	Fruits []string
}

type Book2 struct {
	//Json으로 바인딩할 때, key 이름 바꿔줌.
	Page   int      `json:"add"`
	Fruits []string `json:"fruit"`
}

func ConvertBasicTypes() {
	//encoding basic data types to JSON strings
	fmt.Println("bool-----------")
	bool_, _ := json.Marshal(true)
	fmt.Println(bool_)
	fmt.Println(string(bool_))

	fmt.Println("int-----------")
	int_, _ := json.Marshal(1)
	fmt.Println(int_)
	fmt.Println(string(int_))

	fmt.Println("string-----------")
	strB, _ := json.Marshal("testcase")
	fmt.Println(strB)
	fmt.Println(string(strB))

}

func ConvertSlicesMaps() {

	fmt.Println("str slice------------")
	slice_test := []string{"apple", "peach", "pear"}
	fmt.Println(slice_test)

	slice_json, _ := json.Marshal(slice_test)
	fmt.Println(slice_json)
	fmt.Println(string(slice_json))

	fmt.Println("map ---------")
	map_test := map[string]int{"apple": 5, "lettuce": 7}
	fmt.Println(map_test)

	map_json, _ := json.Marshal(map_test)
	fmt.Println(map_json)
	fmt.Println(string(map_json))

}

func StructTest() {
	struct_test := &Book{
		Page:   1,
		Fruits: []string{"apple", "peach"},
	}
	fmt.Println(struct_test)

	struct_json, _ := json.Marshal(struct_test)
	fmt.Println(struct_json)
	fmt.Println(string(struct_json))

}
func StructBindingTest() {
	struct_test := &Book2{
		Page:   1,
		Fruits: []string{"apple", "peach"},
	}
	fmt.Println(struct_test)

	struct_json, _ := json.Marshal(struct_test)
	fmt.Println(struct_json)
	fmt.Println(string(struct_json))

}
func main() {
	// ConvertBasicTypes()
	// ConvertSlicesMaps()
	StructTest()
	StructBindingTest()

}
