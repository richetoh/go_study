package stringexample

import (
	"fmt"
	"reflect"
	"unsafe"
)

//StringExample shows many ways to write string
func StringExample() {

	first_string := "haha"

	second_string := `"dfdfdfd"`

	fmt.Println("first_string", first_string, "second_string", second_string)

}

//StringLenInByte shows how does len() function works in golang
func StringLenInByte() {

	//1word = 3byte
	korean := "안녕"
	//1word = 1byte
	english := "abc"

	fmt.Println("korean len ", len(korean), "english len ", len(english))
}

//StringLen shows how to count Number of words in string
func StringLen() {

	korean := "안녕"

	english := "abcdefg"

	korean_len := []rune(korean)

	english_len := []rune(english)

	fmt.Println("Korean", korean, " len in byte", len(korean), "len in rune", len(korean_len))
	fmt.Println("English", english, " len in byte", len(english), "len in rune", len(english_len))

}

//ForIterateWrong shows just iterating string with len(string) yields wrong result when it iterates Korean words
func ForIterateWrong() {
	str := "Hello 월드"

	for i := 0; i < len(str); i++ {
		fmt.Printf("타입 : %T 값 : %d 문자값 : %c \n", str[i], str[i], str[i])
	}
}

// For IterateCorrect shows that in order to iterate string word by word, we should use []rune
func ForIterateCorrect() {
	str := "Hello 월드"
	arr := []rune(str)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("타입 : %T 값 : %d 문자값 : %c \n", arr[i], arr[i], arr[i])
	}
}

func UseRangeToIterate() {
	str := "Hello 월드!"

	for _, v := range str {
		fmt.Printf("타입 : %T 값 : %d 문자값 : %c \n", v, v, v)
	}
}

func CompareString() {

	//When comparing string, it compares utf-8 value of each word sequentially. And when there is a difference, it compares and yields that result
	str1 := "BBB"
	str2 := "aaaaAAA"
	str3 := "BBAD"
	str4 := "ZZZ"

	// a = 97, B = 66
	fmt.Printf("%s > %s : %v\n", str1, str2, str1 > str2)

	//BBA -> A = 65, BBB->B = 66
	fmt.Printf("%s > %s : %v\n", str1, str3, str1 > str3)

	// B = 66, Z = 90
	fmt.Printf("%s > %s : %v\n", str1, str4, str1 > str4)
}

func CompareAddressOfString() {
	str := "Hello World"
	slice := []byte(str)

	//unsafe.Pointer(&str) = conver str to pointer type
	//*reflect.StringHeader allows to access Data, Len part of string
	stringheader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	sliceheader := (*reflect.StringHeader)(unsafe.Pointer(&slice))

	fmt.Printf("str : %s, address : %x \n", str, stringheader.Data)
	fmt.Printf("slice : %s, address : %x \n", slice, sliceheader.Data)
}
