package forexample

import "fmt"

func Gugudan(n int) {

	dan := 2
	base := 1

	for {

		for base <= n {
			fmt.Printf("%d * %d = %d\n", dan, base, dan*base)
			base++
		}

		dan++
		base = 1

		if dan > n {
			break
		}

	}

}
