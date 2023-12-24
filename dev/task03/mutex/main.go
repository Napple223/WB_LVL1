package main

import (
	"fmt"
	"sync"
)

//Дана последовательность чисел: 2,4,6,8,10.
//Найти сумму их квадратов(2^2+3^2+4^2….)
//с использованием конкурентных вычислений.

//Вариант с использованием мьютекса.

//Мьютексы позволяют разграничить доступ
//к некоторым общим ресурсам, гарантируя,
//что только одна горутина имеет к ним доступ
//в определенный момент времени. И пока одна
//горутина не освободит общий ресурс, другая
//горутина не может с ним работать.

func main() {
	i := [5]int{2, 4, 6, 8, 10}
	var res int //Переменная для хранения результата.

	//Объявляем мьютекс.
	//В нашем случае можно использовать обычный, не RW,
	//так как мы хотим только писать в разделяемый ресурс.
	var m sync.Mutex
	//WaitGroup для ожидания завершения работы всех горутин.
	var wg sync.WaitGroup

	for _, v := range i {
		wg.Add(1)
		go func(v int) {
			//Блокируем разделяемый ресурс.
			//В нашем случае он также блокирует
			//выполнение остальных горутин
			//до тех пор, пока не будет освобожден
			//мьютекс.
			m.Lock()
			res += v * v
			//После прибавления результата вычисления
			//разблокируем разделяемый ресурс.
			m.Unlock()
			wg.Done()
		}(v)
	}
	//Дожидаемся завершения работы всех горутин.
	wg.Wait()
	fmt.Println(res) //220
}

//Тут стоит добавить, что существует еще 4 вариант - небезопасный.
//Но не вижу смысла его рассматривать, так как состояние гонки
//данных недопустимо.
