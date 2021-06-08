package basic

import "fmt"

//인터페이스 변수 선언이 가능하고, 변수의 값으로 사용할 수 있음.
//1) 메서드는 반드시 메서드 명이 있어야 한다.
//2) 매개변수와 반환이 다르더라도 이름이 같은 메서드는 있을 수 없다.
//3) 인터페이스에서는 메서드 구현을 포함하지 않는다.
type DuckeInterface interface {
	Fly()
	Walk(distance int) int
}

// type WrongInterface interface {
// 	String() string
// 	String(int) string //wrong : String 메서드명이 겹침
// 	_(x int)           //wrong : 이름이 있어야함
// }

//////////////////////

//매개변수 없이 string을 반환하는 Introduce함수를 구현한 모든 타입은 Stringer 인터페이스로 사용될 수 있음.
type Stringer interface {
	Introduce() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) Introduce() string {
	return fmt.Sprintf("나는 %d살 이름은 %s\n", s.Age, s.Name)
}

func TestInterface() {
	student := Student{"철수", 12}
	var stringer Stringer
	stringer = student // stringer 인터페이스 변수에 Student 타입 대입

	fmt.Println(stringer.Introduce())
}
