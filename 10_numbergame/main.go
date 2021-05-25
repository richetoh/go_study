package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int

	// fmt.Printf("맞추고자하는 숫자를 입력하세요~ : ")
	_, err := fmt.Scanln(&n)

	if err != nil {
		stdin.ReadString('\n')
	}
	return n, nil

}
func Numbergame() {

	rand.Seed(time.Now().UnixNano())
	answer := rand.Intn(10)

	var user_input int
	var err error
	count := 1
	for {
		user_input, err = InputIntValue()

		if err != nil {
			fmt.Println("숫자만 입력하세요~")
		}

		if user_input > answer {
			fmt.Println("입력하신 값이 더 큽니다.")
		} else if user_input < answer {
			fmt.Println("입력하신 값이 더 작습니다.")
		} else {
			fmt.Println("정답입니다~")
			fmt.Printf("입력한 횟수 : %d\n", count)
			break
		}
		count++
	}
}

func Slot() {

	money := 1000

	var user_input int
	var err error
	var answer int
	min := 1
	max := 5

	for {
		rand.Seed(time.Now().UnixNano())
		answer = rand.Intn((max - min) + 1 + min)

		fmt.Printf("슬롯 값을 입력하세요 : ")

		user_input, err = InputIntValue()
		if err != nil {
			fmt.Println("숫자만 입력하세요~")
		}

		if user_input == answer {
			money += 500
			if money >= 5000 {
				fmt.Println("돈을 넘 많이 따셔서 그만하세요~ ")
				break
			}
			fmt.Printf("축하합니다 500원 추가 ~ : %d\n", money)
		} else {
			money -= 100
			if money <= 0 {
				fmt.Println("돈없으면 게임 그만하세요~ ")
				break
			}
			fmt.Printf("틀렸습니다 100원 개이득 ~ : %d\n", money)
		}

	}

}
func main() {
	Slot()
}
