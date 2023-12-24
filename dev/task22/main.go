package main

import (
	"fmt"
	"math/big"
)

//Разработать программу, которая перемножает, делит,
//складывает, вычитает две числовых переменных a,b,
//значение которых > 2^20.

//Способ с помощью пакета big.
//Удобный способ из коробки.

// Функция сложения.
func sum(f, s *big.Int) *big.Int {
	var res big.Int
	res.Add(f, s)
	return &res
}

// Функция вычитания.
func sub(f, s *big.Int) *big.Int {
	var res big.Int
	res.Sub(f, s)
	return &res
}

// Функция умножения.
func multiply(f, s *big.Int) *big.Int {
	var res big.Int
	res.Mul(f, s)
	return &res
}

// Функция деления.
func div(f, s *big.Int) *big.Int {
	var res big.Int
	res.Div(f, s)
	return &res
}

func main() {
	f := new(big.Int)
	s := new(big.Int)

	f.SetString("710360490247306597412", 10)
	s.SetString("60123794287369540287", 10)
	fmt.Println(sum(f, s))
	fmt.Println(sub(f, s))
	fmt.Println(multiply(f, s))
	fmt.Println(div(f, s))
}
