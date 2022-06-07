package main

import (
	"fmt"

	"github.com/PopovVA/go_level_1/lesson_10/numbers"
)

func main() {
	var n int64
	fmt.Print("Введите число N: ")
	fmt.Scanln(&n)

	res := numbers.CountPrimes(n)

	fmt.Print(res)

}
