package ducktyping

import "fmt"

//인터페이스를 포함하는 인터페이스
//구조체에서 다른 구조체를 포함된 필드로 가질 수 있듯이, 인터페이스도 다른 인터페이스를 포함할 수 있다.
type Reader interface {
	Read() (n int, err error)
	Close() error
}

type Writer interface {
	Write() (n int, err error)
	Close() error
}

type ReadWriter interface {
	//Reader의 메서드 집합을 포함한다
	Reader
	//Writer의 메서드 집합을 포함한다.
	Writer
}

//인터페이스를 선언할 때 이름이 겹치면 안디지만, Close() error은 같은 메서드 형식이므로 합쳐짐
//1) Read(), Write(), Close()메서드를 모두 포함한 타입 = ReadWriter 인터페이스 사용 가능
//2) Read(), Close() 메서드를 포함한 타입 = Reader 인터페이스 사용가능
//3) Write(), Close() 메서드를 포함한 타입 = Writer 인터페이스 사용가능
//4) Read(), Write() 메서드를 포함한 타입 = Close()가 없으므로 모든 인터페이스 사용 불가능

//빈 인터페이스 interface{}를 인수로 받기
//빈 interface는 어떤 값이든 받을 수 있는 함수, 메서드, 변숫값을 만들 때 사용한다.

func Classify(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Printf("v is int %d \n", int(t))
	case string:
		fmt.Printf("v is string %s\n", string(t))
	default:
		fmt.Printf("v is Not supported type Type = %T, Value = %v\n", t, t)
	}
}

type Student struct {
	Name string
}

func ClassifyTypes() {
	Classify(10)
	Classify("haha")
	Classify(Student{"hehe"})
}

// 인터페이스 기본값 Nil
// 인터페이스의 기본값은 유효하지 않은 메모리 주소를 나타내는 nil이다.

type Attacker interface {
	Attack()
}

func TryNilInterface() {
	var att Attacker
	//att의 초기값이 없기 때문에 att는 현재 기본값은 nil
	//따라서 아래는 nil.Attack()이기 때문에 런 타임 에러가 발생함.
	att.Attack()

	//이것 외에도 nil 값을 기본으로 갖는 타입으로는 포인터, 인터페이스, 함수 타입, 슬라이스, 맵, 채널 등이 있다.
}
