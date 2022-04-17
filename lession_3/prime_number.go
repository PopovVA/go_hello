package main

import (
	"fmt"
	"math/big"
)

// 2. Задание для продвинутых (необязательное).
// Написать приложение, которое ищет все простые числа от 0 до N включительно.
// Число N должно быть задано из стандартного потока ввода.

func main() {
	var n, res int64
	fmt.Print("Введите число N: ")
	fmt.Scanln(&n)

	var i int64 = 0
	for ; i <= n; i++ {
		if big.NewInt(i).ProbablyPrime(0) {
			fmt.Println(i, " - простое")
		}
	}

	fmt.Printf("Число простых чисел от 0 до %v: %v \n", n, res)
}
