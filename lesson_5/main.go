package main

import "fmt"

// 1. Напишите приложение, рекурсивно вычисляющее заданное из стандартного ввода число Фибоначчи.
// 2. Оптимизируйте приложение за счёт сохранения предыдущих результатов в мапе.
// 3. Проверьте себя: вам должны быть знакомы следующие ключевые слова Go: func, return, defer, struct, type.
// 4. Посмотрите задачи из предыдущих уроков. Как можно улучшить дизайн задач? Что бы вы разбили на отдельные функции или даже пакеты?

var cache map[int]int = make(map[int]int)

func main() {
	cache[0], cache[1] = 0, 1
	var a int
	fmt.Println("Введите порядковый номер: ")
	fmt.Scanln(&a)
	fmt.Printf("Результат: %v", fib(a))
}

func fib(n int) int {
	key, hasKey := cache[n]
	if hasKey {
		return key
	} else {
		a := fib(n - 1)
		b := fib(n - 2)

		cache[n-1] = a
		cache[n-2] = b
		return a + b
	}

}
