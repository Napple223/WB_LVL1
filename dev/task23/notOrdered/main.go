package main

import "fmt"

//Удалить i-ый элемент из слайса.

//Способ без сохранения порядка.

func delElemFromSlice(input []int, idx int) []int {
	l := len(input) - 1
	if idx < 0 || idx > l {
		fmt.Println("Индекс вне диапазона.")
		return input
	}
	//Меняем значение по удаляемому индексу на последний
	//эдемент слайса.
	input[idx] = input[l]
	//Обрезаем слайс.
	return input[:l]
}

func main() {
	input := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(delElemFromSlice(input, 2))
}
