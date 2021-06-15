package variadic

import "fmt"

func Sum(nums ...int) int{
	sum := 0

	fmt.Printf("nums 타입 : %T\n", nums)
	for _, v := range nums {
		sum += v
	}

	return sum
}

func Print(args ...interface{}) {

	for _, arg := range args {
		switch f:= arg.(type) {
		case bool:
			val := arg.(bool)
			fmt.Println(val, "=", f)
		case float64:
			val := arg.(float64)
			fmt.Println(val, "=", f)
		case int:
			val := arg.(int)
			fmt.Println(val, "=", f)
		case string:
			val := arg.(string)
			fmt.Println(val, "=", f)
		}
	}
}