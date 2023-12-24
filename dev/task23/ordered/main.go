package main

import "fmt"

//Удалить i-ый элемент из слайса.

//Способ, если порядок элементов имеет значение.

func delElemFromSlice(input []int, idx int) []int {
	l := len(input) - 1
	if idx < 0 || idx > l {
		fmt.Println("Индекс вне диапазона.")
		return input
	}
	//Сдвигаем элементы слайса от удаляемого до len(input)-1.
	for i := idx; i < l; i++ {
		input[i] = input[i+1]
	}
	//Возвращаем обрезанный слайс.
	return input[:l]
}

func main() {
	input := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(delElemFromSlice(input, 2))
}
