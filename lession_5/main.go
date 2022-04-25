package main
import "fmt"

// 1. Напишите приложение, рекурсивно вычисляющее заданное из стандартного ввода число Фибоначчи.
// 2. Оптимизируйте приложение за счёт сохранения предыдущих результатов в мапе.
// 3. Проверьте себя: вам должны быть знакомы следующие ключевые слова Go: func, return, defer, struct, type.
// 4. Посмотрите задачи из предыдущих уроков. Как можно улучшить дизайн задач? Что бы вы разбили на отдельные функции или даже пакеты?

func main() {
	var a int
	fmt.Print("Введите порядковый номер: ")
	fmt.Scanln(&a)
	fmt.Printf("Результат: %v", fib(a))
}

func fib(n int) int {
    if n == 0 {
      return 0
    }
    if n == 1 {
      return 1
    }
    return fib(n - 1) + fib(n - 2)
}