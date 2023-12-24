package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//Реализовать структуру-счетчик, которая
//будет инкрементироваться в конкурентной
//среде. По завершению программа должна
//выводить итоговое значение счетчика.

//Способ с использованием атомиков.

// Структура счетчика.
type counter struct {
	val int64
}

// Метод инкремент.
func (c *counter) increment() {
	atomic.AddInt64(&c.val, 1)
}

// Метод для вывода значения счетчика.
func (c *counter) printVal() {
	fmt.Println(c.val)
}

func main() {
	//WaitGroup для наглядности.
	var wg sync.WaitGroup
	//Новый экземпляр счетчика.
	counter := counter{}
	//Запускаем 5 горутин, которые по 3 раза
	//накрутят счетчик.
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 3; j++ {
				counter.increment()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	counter.printVal()
}
