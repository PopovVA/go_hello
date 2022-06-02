package calculator

import (
	"errors"
	"math"
)

func Validate(b float32, op string) error {
	switch op {
	case "/":
		if b == 0 {
			return errors.New("деление на ноль")
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
		return errors.New("операция выбрана неверно")
	}
	return nil
}

func Calculate(a float32, b float32, op string) (res float32, err error) {
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
	return res, err
}

func factorial(n float32) (res float32) {
	if n > 0 {
		res = n * factorial(n-1)
		return res
	}
	return 1
}
