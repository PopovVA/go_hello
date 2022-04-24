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
  	var i int64
	for i = 0; i <= n; i++ {
		if big.NewInt(i).ProbablyPrime(0) {
			fmt.Println(i, " - простое")
      			res++
		}
	}

	fmt.Printf("Число простых чисел от 0 до %v: %v \n", n, res)
}
