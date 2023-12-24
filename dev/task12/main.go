package main

import "fmt"

//Имеется последовательность строк
//(cat, cat, dog, cat, tree)
//создать для нее собственное множество.

//Способ решения с помощью мап.

func main() {
	//Входные данные.
	str := []string{"cat", "cat", "dog", "cat", "tree"}

	//Создаем мап, где ключом является множетсво,
	//а значением - пустая структура.
	set := make(map[string]struct{})

	//Добавляем все множетсва в мап,
	//если множетсво уже существует,
	//игнорируем значение.
	for _, v := range str {
		_, ok := set[v]
		if ok {
			continue
		}
		set[v] = struct{}{}
	}
	//Слайс для вывода результата.
	res := make([]string, len(set), len(set))
	idx := 0

	//Добавлем множества в слайс.
	for k := range set {
		res[idx] = k
		idx++
	}

	fmt.Println(res)
}
