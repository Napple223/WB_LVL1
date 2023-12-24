package main

import (
	"fmt"
	"sync"
)

//Реализовать все возможные способы остановки выполнения горутины.

//Использование done канала. Принцип действия схож с
//механизмом работы context.WithCancel().
//Создаем done канал, передаем его всем горутинам.
//При необходимости завершения работы горутин,
//закрываем done канал.
//Удобно при использовании большого количества пишуших и
//читающих горутин. Закрывать done канал должен только
//тот, кто его создал.

func main() {
	//Создаем done канал.
	done := make(chan struct{})
	//Создаем канал для передачи и получения данных.
	data := make(chan int)
	//WaitGroup для наглядности.
	var wg sync.WaitGroup
	wg.Add(1)
	//Горутина чтения данных с канала. Передаем ей наш done канал.
	go func(done <-chan struct{}, data <-chan int) {
		for {
			select {
			//В случае закрытия done канала, завершаем работу горутины.
			case <-done:
				fmt.Println("Done канал закрыт. Горутина завершает работу.")
				wg.Done()
				return
			case d := <-data:
				fmt.Printf("Поучены данные: %d\n", d)
			}
		}
	}(done, data)

	for i := 1; i <= 5; i++ {
		data <- i
	}
	//Инициализируем завершение работы горутины.
	close(done)
	wg.Wait()
	//Закрываем канал.
	close(data)
}
