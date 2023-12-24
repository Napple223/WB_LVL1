package main

import (
	"fmt"
	"sync"
)

//Реализовать конкурентную запись данных в map.

//Решение с использованием потокобезопасной версии map - sync.Map
//sync.Map оптимизована для нескольких кейсов:
//1. Когда запись происходит кратно реже, чем чтение.
//2. Когда много горутин часто читает, записывает и перезаписывает
//данные для разных ключей.
//В таких случаев использование sync.Map сильно уменьшает
//вероятность блокировки, а значит мы выигрываем в скорости работы.
//Но в оставшихся может замедлить работу программы из-за
//постоянных блокировок.

func main() {
	//Инициализируем sync.Map
	var collection sync.Map
	//Используем WaitGroup чтобы дождаться завершения работы всех горутин.
	var wg sync.WaitGroup
	//Создаем пул горутин, выполняющих запись в map.
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func(i int) {
			//Вызываем метод для записи.
			collection.Store(i, i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	//Итерируемся по мап и выводим все существующие пары ключ - значение.
	collection.Range(func(k, v interface{}) bool {
		fmt.Printf("Ключ: %d, значение:%d\n", k, v)
		return true
	})
}