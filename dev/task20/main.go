package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

// Функция, разворачивающая порядок слов в строке.
func reversWords(input string) string {
	//Разбиваем стоку на слайс, разделителем является пробел.
	s := strings.Split(input, " ")
	//Разворачиваем порядок слов в слайсе.
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	//Объединяем строку, разделителем является пробел.
	return strings.Join(s, " ")
}

func main() {
	//Работает с символами юникод.
	fmt.Println(reversWords("the quick brown 狐 jumped over the lazy 犬"))
}
