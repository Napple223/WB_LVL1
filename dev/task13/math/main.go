package main

import (
	"fmt"
)

//Поменять местами два числа без создания временной переменной.

//С помощью XOR или математический метод.

func main() {
	a := 28
	b := 5
	fmt.Println(a, b)
	a ^= b
	b ^= a
	a ^= b
	fmt.Println(a, b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Println(a, b)
}
