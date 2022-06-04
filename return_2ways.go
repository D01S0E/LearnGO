package main

// 나누기 프로그램 // return 방법 두가지 학습
import "fmt"

func main() {
	fmt.Println(division1(10, 4))
	fmt.Println(division2(10, 4))
}

func division1(divided, divisor int) (int, int) {
	var quotient int = divided / divisor
	var remainder int = divided % divisor

	return quotient, remainder
}

func division2(divided, divisor int) (quotient, remainder int) {
	quotient = divided / divisor
	remainder = divided % divisor

	return
}
