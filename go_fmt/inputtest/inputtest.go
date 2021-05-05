package inputtest

import "fmt"

//Scan() 입력에서 값을 받음
//Scanf() 입력서식 형태로 값을 입력
//Scanln() 입력에서 한 줄을 읽어서 값을 입력 받음
func ScanningTest() {

	var a int
	var b int

	n, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}

}

//read inputs with specific format
func ScanfTest() {
	var a int
	var b int

	n, err := fmt.Scanf("%d %d\n", &a, &b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}

}

//read one line and split elements by whitespace
func ScanlnTest() {
	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}

}
