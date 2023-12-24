package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая проверяет,
//что все символы в строке уникальные
//(true — если уникальные, false etc).
//Функция проверки должна быть регистронезависимой.
//Например:
//abcd — true
//abCdefAaf — false
//aabcd — false

//Решение с использованием мап.

func symbolsIsUnique(str string) bool {
	//Делаем нашу строку регистронезависимой.
	lowerStr := strings.ToLower(str)
	//Создаем мап с длиной, равной количеству символов в строке,
	//чтобы избежать доп. аллокаций памяти.
	m := make(map[rune]struct{}, len(lowerStr))
	//Проходим циклом по мап и ищем повторяющиеся символы,
	//если находим, возвращаем false.
	for _, v := range lowerStr {
		_, ok := m[v]
		if ok {
			return false
		}
		m[v] = struct{}{}
	}
	return true
}

func main() {
	str := "abcd"
	fmt.Println(symbolsIsUnique(str))
	str = "abCdefAaf"
	fmt.Println(symbolsIsUnique(str))
	str = "aabcd"
	fmt.Println(symbolsIsUnique(str))
}
