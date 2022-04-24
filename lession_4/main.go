package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 1. Познакомьтесь с алгоритмом сортировки вставками.
// Напишите приложение, которое принимает на вход набор целых чисел и отдаёт на выходе его же в отсортированном виде.
// Сохраните код, он понадобится нам в дальнейших уроках.
// 2. Проверьте себя:
// a. вам должны быть знакомы следующие ключевые слова Go: map, range;
// b. вам должны быть знакомы функции: make, len, cap, append, copy, delete

func main() {
	fmt.Println("Введите числа через запятую (1,2,3, ...)")
	fmt.Println("После окончания ввода нажмите ENTER")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	text := strings.Split(scanner.Text(), ",")
	numbers := []int64{}
	for _, el := range text {
		num, err := strconv.ParseInt(el, 10, 64)
		if err != nil {
			panic("Failed to read the number")
		}
		numbers = append(numbers, num)
	}

	insertionSort(numbers)

	fmt.Printf("Отсортированный массив: %v \n", numbers)
}

func insertionSort(ar []int64) {
	for i := 1; i < len(ar); i++ {
		x := ar[i]
		j := i
		for ; j >= 1 && ar[j-1] > x; j-- {
			ar[j] = ar[j-1]
		}
		ar[j] = x
	}
}
