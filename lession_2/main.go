// 1. Напишите программу для вычисления площади прямоугольника. Длины сторон прямоугольника должны вводиться пользователем с клавиатуры.
// 2. Напишите программу, вычисляющую диаметр и длину окружности по заданной площади круга. Площадь круга должна вводиться пользователем с клавиатуры.
// 3. С клавиатуры вводится трехзначное число. Выведите цифры, соответствующие количество сотен, десятков и единиц в этом числе.
// 4. Проверьте себя:
// a. Вам должны быть знакомы следующие ключевые слова Go: package, import, func, var, const.
// b. Вам должны быть знакомы следующие константы: true, false, iota, nil.
// c. Вам должны быть знакомы следующие типы: int, int8, int16 , int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, complex128, complex64, bool, byte, rune, string.
// d. Вам должны быть знакомы функции complex, real, imag, len.

package main

import (
	"fmt"
	"math"
)

func main() {

	rectangleArea()

	diameterCircumference()

	parseNumber()

}

func rectangleArea() {
	a, b := 0, 0

	fmt.Println("Задание №1. Площадь прямоугольника")

	fmt.Println("Введите длину стороны A: ")
	fmt.Scanln(&a)

	fmt.Println("Введите длину стороны B: ")
	fmt.Scanln(&b)

	fmt.Printf("Площадь прямоугольника: %v \n", a*b)
}

func diameterCircumference() {
	fmt.Println("Задание №2. Диаметр и длина окружности по заданной площади круга")

	var s float64

	fmt.Println("Введите площадь круга: ")
	fmt.Scanln(&s)

	r := math.Sqrt(s / math.Pi)

	fmt.Printf("Радиус круга %.2f \n", r)

	c := float64(2 * math.Pi * r)

	fmt.Printf("Длина круга %.2f \n", c)

}

func parseNumber() {
	fmt.Println("Задание №3. Количество сотен, десятков и единиц в числе")

	var units int

	fmt.Println("Введите число: ")
	fmt.Scanln(&units)

	if units > 999 || units < 100 {
		panic("Неверное число")
	}

	hundreds := math.Abs(float64(units / 100))
	units = units - int(hundreds)*100

	tens := math.Abs(float64(units / 10))
	units = units - int(tens)*10

	fmt.Printf("Количество сотен: %v\n", hundreds)
	fmt.Printf("Количество десятков: %v\n", tens)
	fmt.Printf("Количество единиц: %v\n", units)

}
