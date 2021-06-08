package convert

//인터페이스 변수를 타입 변환을 통해서 구체화된 다른 타입이나 다른 인터페이스로 변환할 수 있다.

//1) 구체화된 다른 타입으로 변환하기
//2) 다른 인터페이스로 타입 변환하기

//구체화된 다른 타입으로 변환하기
//인터페이스 변수를 다른 구체화된 타입으로 타입변환할 수 있다.
//이 방법은 인터페이스를 본래의 구체화된 타입으로 복원할 때 주로 사용한다.
//사용 방법은 인터페이스 변수 뒤에 점 .을 찍고 소괄호() 안에 변경하려는 타입을 쓰면된다.

import "fmt"

type Stringer interface {
	MyName()
}

type Student struct {
	Name string
}

type Teacher struct {
}

func (t Teacher) MyName() {
	fmt.Println("I'm Teacher Dude")
}
func (s Student) MyName() {
	fmt.Println("My name is ", s.Name)
}

func PrintName(stringer Stringer) {
	s := stringer.(Student)
	fmt.Printf("Name %s", s.Name)
}

func TestConversion() {
	s := Student{"Richet"}
	PrintName(s)

	//below code will incur runtime error because PrintName will try to convert stringer interface to Student but t is pointing Teacher interface
	t := Teacher{}
	PrintName(t)
}

//다른 인터페이스로 타입 변환하기
//인터페이스 변환을 통해 구체화된 타입 뿐 아니라 다른 인터페이스로 타입 변환할 수 있다.
//이 때는 구체화된 타입으로 변환할 때와는 달리 변경되는 인터페이스가 변경 전 인터페이스를 포함하지 않아도 된다.
//하지만 인터페이스가 가리키고 있는 실제 인스턴스가 변환하고자 하는 다른 인터페이스를 포함해야 한다.

//타입 변환 성공 여부를 반환하여, 확ㅇ니하기
func TestConversion2() {
	var a Stringer

	s, ok := a.(Student)
	if !ok {
		fmt.Println("can't convert")
	} else {
		fmt.Printf("s is converted to %T\n", s)
	}

	if s, ok := a.(Student); ok {
		fmt.Printf("s is converted to %T\n", s)
	}
}
