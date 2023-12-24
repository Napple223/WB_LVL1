package main

import (
	"fmt"
	"strconv"
	"strings"
)

//Дана последовательность температурных
//колебаний: -25.4, -27.0 13.0, 19.0,
//15.5, 24.5, -21.0, 32.5.
//Объединить данные значения в группы с
//шагом в 10 градусов.
//Последовательность в подмножноствах не важна.
//Пример: -20:{-25.0, -27.0, -21.0},
//10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

//Способ с применением мап.
// - способа:
//невозможность контролировать порядок вывода
//множеств температур.

func main() {
	//Входные данные.
	input := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	//Инициализируем мап, где будем хранить
	//подмножества.
	groups := make(map[int][]float64)

	for idx, val := range input {
		//Пример формирования ключа:
		//"Отрезаем" у числа дробную часть -25,4 -> -25,
		//"отрезаем" единицы -25 -> -2,
		//восстанавливаем единицы -2 -> -20.
		intKey := int(val) / 10 * 10
		//Проверяем существование ключа в мап.
		_, ok := groups[intKey]
		//Если не существует, то инициализируем слайс.
		if !ok {
			groups[intKey] = []float64{}
		}
		//Если существует, добавляем температуру в слайс.
		groups[intKey] = append(groups[intKey], input[idx])
	}

	//Формируем строку.
	buildStr := strings.Builder{}
	for key, val := range groups {
		buildStr.WriteString(strconv.Itoa(key))
		buildStr.WriteString(":{")
		for i, v := range val {
			buildStr.WriteString(fmt.Sprintf("%.1f", v))
			if i+1 != len(val) {
				buildStr.WriteString(", ")
			}
		}
		buildStr.WriteString("}, ")
	}
	//Срезаем последние ", ".
	output := buildStr.String()[0 : buildStr.Len()-2]
	fmt.Println(output)

}
