package main

// 1. Доработать калькулятор из методички: больше операций и валидации данных (проверка деления на 0, возведение в степень, факториал)
// + форматирование строки при выводе дробного числа ( 2 знака после точки)
// 3. Проверьте себя. Вам должны быть знакомы следующие ключевые слова Go:
// if, else, switch, case, default, for, break, continue.

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func main() {
	var a, b float32
	var op string
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)
	fmt.Print("Введите арифметическую операцию (+, -, *, /, exp, fac): ")
	fmt.Scanln(&op)

	err := validate(a, b, op)
	if err != nil {
		fmt.Printf("%v \n", err.Error())
		os.Exit(1)
	}
	err, res := calculate(a, b, op)
	if err != nil {
		fmt.Printf("%v \n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("Результат выполнения операции: %.2f\n", res)
}

func validate(a float32, b float32, op string) error {
	switch op {
	case "/":
		if b == 0 {
			return errors.New("Деление на ноль")
		}
	case "+":
		return nil
	case "-":
		return nil
	case "*":
		return nil
	case "exp":
		return nil
	case "fac":
		return nil
	default:
		return errors.New("Операция выбрана неверно")
	}
	return nil
}

func calculate(a float32, b float32, op string) (err error, res float32) {
	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	case "exp":
		res = float32(math.Pow(float64(a), float64(b)))
	case "fac":
		aFac := factorial(a)
		bFac := factorial(b)
		res = aFac / bFac
	}
	return err, res
}

func factorial(n float32) (res float32) {
	if n > 0 {
		res = n * factorial(n-1)
		return res
	}
	return 1
}
