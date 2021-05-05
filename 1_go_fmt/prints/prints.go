package prints

import "fmt"

///prints는 Print, Println, Printf의 차이점을 보기위해 만든 함수
func Prints() {
	var a int = 10
	var b int = 20
	var f float64 = 32234234.454

	fmt.Print("a:", a, "b:", b)

	fmt.Println("a:", a, "b:", b, "f:", f)
	fmt.Printf("a: %d b: %d f: %f \n", a, b, f)
}

//최소 출력 너비를 정하는 Printf
func PrintfTest() {

	var a int = 10
	var b int = 20
	var c int = 12312

	fmt.Printf("%5d, %5d\n", a, b)
	fmt.Printf("%05d, %05d\n", a, b)
	fmt.Printf("%-5d %-05d\n", a, b)

	fmt.Printf("%5d, %5d\n", c, c)
	fmt.Printf("%5d, %05d\n", c, c)
	fmt.Printf("%-5d %-05d\n", c, c)

}

//소수점 확인 printf
func DecimalTest() {

	var a = 324.13455
	var c = 3.14

	//최소 너비 8, 소숫점 이하 2자리, 나머지는 0을 채움
	fmt.Printf("%08.2f\n", a)

	//최소 너비 8, 총숫자 2자리, 0을 채움
	fmt.Printf("%08.2g\n", a)

	//최소너비 8, 총 숫자 5자리, 나머지는 빈공간
	fmt.Printf("%8.5g\n", a)

	//소수점 이하 6자리까지 출력
	fmt.Printf("%f\n", c)

}
