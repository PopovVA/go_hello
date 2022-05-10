// 1. С какими интерфейсами мы уже сталкивались в предыдущих уроках?
// Обратите внимание на уроки, в которых мы читали из стандартного ввода и писали в стандартный вывод.

// Для считавания ввода в функции Scanln используется type any = interface{}, здесь не много странно,
// выглядит как будто этому интерфейсу соотвествует любой тип и строгой типизации тут нет
// Тоже самое для функции из fmt пакета, есть некий any интерфейс
// Поправьте меня если я не прав

// 2. Посмотрите примеры кода в своём портфолио. Везде ли ошибки обрабатываются грамотно? Хотите ли вы переписать какие-либо функции?
// В lesson 2 можно было вместо паники возвращать ошибки
// В lesson 3 (калькулятор) уже использовал обработку ошибок, но стоило бы поставить последним аргументом, как рекомендуется по код стайлу
// В lesson 4 панику можно было заменить на continue с выводом сообщения, что не удалось считать число, думаю что можно было бы обойтись без обработки ошибок

// 3. Проверьте себя:
// a. Вам должны быть знакомы следующие ключевые слова Go: interface.
// b. Вам должны быть знакомы следующие функции: panic, recover.

// *4. (не обязательное) Задание для желающих попробовать поработать с интерфейсами. На оценку не влияет:
// Необходимо объявить свой тип, обернув в него тип []byte - (слайс байтов).

// Затем, необходимо реализовать на нем такие методы,
// чтобы он удовлетворял интферйесам io.Reader (из него можно читать байты) и io.Writer (а так же писать их туда).
// Затем, используя функции пакета io:

// С помощью io.WriteString записать в переменную вашего типа произвольную строку
// С помощью io.ReadAll( ) считать вашу строку обратно (вообще говоря, он возвращает слайс байт, но его легко привести к виду строки)
// В случае затруднений можно писать мне, обсудим (tg @VladimirLozhkin),
// или обратиться к коду реализации типа Buffer{} из пакета bytes. (Но там спойлеры, сначала попробуйте сами :D )

package main

import (
	"fmt"
	"io"
)

const st = "hello world"

type MySlice struct {
	data  []byte
	index int64
}

func (slice *MySlice) Read(p []byte) (n int, err error) {
	if slice.index >= int64(len(slice.data)) {
		err = io.EOF
		return n, err
	}
	n = copy(p, slice.data[slice.index:])
	slice.index += int64(n)
	return n, err
}

func (slice *MySlice) Write(p []byte) (n int, err error) {
	slice.data = p
	return n, nil
}

func main() {
	fmt.Printf("Initial string - %v\n", st)
	var slice MySlice = MySlice{
		data: []byte{},
	}
	_, err := io.WriteString(&slice, st)
	if err != nil {
		fmt.Errorf("Unable to write string", err.Error())
		return
	}
	fmt.Printf("Converted to bytes %v\n", slice.data)
	res, err := io.ReadAll(&slice)
	if err != nil {
		fmt.Errorf("Unable to read bytes", err.Error())
		return
	}
	fmt.Printf("Converted to string %v\n", string(res))
}
