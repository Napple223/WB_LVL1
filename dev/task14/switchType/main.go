package main

import "fmt"

//Разработать программу, которая в рантайме
//способна определить тип переменной:
//int, string, bool, channel из переменной типа interface{}.

//Способ с переключением типов.

func main() {
	i := 1
	b := false
	s := "hi there"
	chInt := make(chan int)
	chBool := make(chan bool)
	chStr := make(chan string)
	f := 38.5
	whatType(i)
	whatType(b)
	whatType(s)
	whatType(chInt)
	whatType(chBool)
	whatType(chStr)
	whatType(f)
}

func whatType(val interface{}) {
	switch val.(type) {
	case int:
		fmt.Println("type = int")
	case bool:
		fmt.Println("type = bool")
	case string:
		fmt.Println("type = string")
	case chan int:
		fmt.Println("type = chan int")
	case chan bool:
		fmt.Println("type = chan bool")
	case chan string:
		fmt.Println("type = chan string")
	default:
		fmt.Println("what type is it?")

	}
}
