package main

import "fmt"

//Разработать программу, которая в рантайме
//способна определить тип переменной:
//int, string, bool, channel из переменной типа interface{}.

//Способ с использованием флага %T.

func main() {
	i := 1
	b := false
	s := "hi there"
	chInt := make(chan int)
	chBool := make(chan bool)
	chStr := make(chan string)
	whatType(i)
	whatType(b)
	whatType(s)
	whatType(chInt)
	whatType(chBool)
	whatType(chStr)
}

func whatType(val interface{}) {
	fmt.Printf("type = %T\n", val)
}
