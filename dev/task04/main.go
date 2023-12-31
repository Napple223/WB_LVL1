package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

//Реализовать постоянную запись данных в канал
//(главный поток). Реализовать набор из N воркеров,
//которые читают произвольные данные из канала и
//выводят в stdout. Необходима возможность выбора
//количества воркеров при старте.
//Программа должна завершаться по нажатию Ctrl+C.
//Выбрать и обосновать способ завершения работы всех воркеров.

//В данной задаче количество воркеров заранее не определено,
//поэтому предлагаю использовать 1 общий done канал, который
//позволит закончить выполнение операций горутинами и завершит
//их работу. Такая конструкция позволит оповестить все
//горутины единовременно, без использования доп.ресурсов.

func main() {
	//Используем WaitGroup для ожидания завершения работы горутин.
	var wg sync.WaitGroup
	//Создаем done канал.
	done := make(chan struct{})
	//Создаем канал для передачи данных.
	data := make(chan int)
	//Переменная, определяющая количество воркеров.
	var w int
	fmt.Print("Введите количество воркеров: ")
	//Считываем количество воркеров из stdin.
	_, err := fmt.Scanln(&w)
	if err != nil {
		fmt.Println("Введенное значение должно быть целым числом")
		return
	}

	//Создаем пул воркеров.
	for i := 1; i <= w; i++ {
		wg.Add(1)
		go worker(i, done, data, &wg)
	}
	//Создаем сигнальный канал.
	notify := make(chan os.Signal, 1)
	//Объявляем какие сигналы необходимо слушать и,
	//в случае наступления одного из событий, пишем
	//в сигнальный канал.
	signal.Notify(notify, syscall.SIGTERM, syscall.SIGINT)
	//Рандомизируем рандом.
	rand.NewSource(time.Now().UnixNano())
	loop := true
	//Создаем тикер, который раз в 0,5 секунды будет подавать
	//сигнал в канал.
	t := time.NewTicker(500 * time.Millisecond)
	for loop {
		//Слушаем каналы.
		//В случае наступления одного из
		//сигнальных случаев, выходим из цикла.
		//Если сигнальный случай не наступил,
		//пишем данные в канал для горутин раз в 0,5 секунды.
		select {
		case <-notify:
			loop = false
		case <-t.C:
			data <- rand.Intn(1000)
		}
	}
	//Не забываем остановить тикер.
	t.Stop()
	//Подаем сигнал для остановки работы всех горутин.
	close(done)
	//Ожидаем завершение работы всех горутин.
	wg.Wait()
	fmt.Println("Все горутины остановлены. Программа завершила работу.")
}

// Функция для создания воркера.
func worker(id int, done <-chan struct{}, data <-chan int, wg *sync.WaitGroup) {
	for {
		//Select ожидает наступления одного из описанных кейсов.
		//В нашем случае будет происходить блокировка горутины,
		//так как нет кейса default.
		select {
		//В случае закрытия канала done все горутины получат
		//сигнал и плавно завершат работу.
		case <-done:
			fmt.Printf("Горутина №%d завершила работу.\n", id)
			wg.Done()
			return
		case d := <-data:
			fmt.Printf("Получены данные горутиной №%d: %d\n", id, d)
		}
	}
}
