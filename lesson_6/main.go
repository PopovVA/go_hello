package main

// 1. Проанализируйте задания предыдущих уроков.
// a. В каких случаях необходима была явная передача указателя в качестве входных параметров и возвращаемых результатов или в качестве приёмника в методах?
// b. В каких случаях мы фактически имеем дело с указателями при передаче параметров, хотя явно их не указываем?

// В четвертом уроке в этой функции можно было использовать указатель для переменной ar, чтобы избежать копирования,
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

// больше я в своём коде как-то не нашел применения указателям. Мне кажется это больше применимо при работе с объектами данных, а не с примитивами

// 2. Для арифметического умножения и разыменования указателей в Go используется один и тот же символ — оператор (*).
// Как вы думаете, как компилятор Go понимает, в каких случаях в выражении имеется в виду умножение, а в каких — разыменование указателя?

// Отличный ответ из гугла:
// Умножение является инфиксным оператором, который запрашивает два значения,
// в то время как оператор разыменования используется в роли префикса для единственной переменной.