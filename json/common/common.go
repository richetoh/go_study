package common

type Book struct {
	Page   int
	Fruits []string
}

type Book2 struct {
	//Json으로 바인딩할 때, key 이름 바꿔줌.
	Page   int      `json:"add"`
	Fruits []string `json:"fruit"`
}
