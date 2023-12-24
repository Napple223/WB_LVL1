package main

import (
	"errors"
	"fmt"
	"sync"
)

//Реализовать конкурентную запись данных в map.

//Вариант с использованием mutex + map.
//Решение аналогичное sync.Map, но в этом случае,
//мы можем точнее настраивать случаи блокировки чтения
//и записи для нашего случая.

// Тип данных collection.
type collection struct {
	m     sync.Mutex
	store map[int]int
}

// Функция - конструктор, возвращающая новый экземпляр collection.
func new() *collection {
	s := make(map[int]int)
	c := collection{
		store: s,
	}
	return &c
}

// Функция для конкурентного добавления значений в мап.
func (c *collection) add(k, v int) error {
	c.m.Lock()
	_, ok := c.store[k]
	if ok {
		err := fmt.Sprintf("Ключ %d уже существует", k)
		return errors.New(err)
	}
	c.store[k] = v
	c.m.Unlock()
	return nil
}

// Функция для получения всех пар ключ-значения из мап.
func (c *collection) getAll() {
	c.m.Lock()
	fmt.Println(c.store)
	c.m.Unlock()
}

func main() {
	//WaitGroup чтобы дождаться завершения работ всех горутин.
	var wg sync.WaitGroup
	//Создаем новый экземпляр типа collection.
	nums := new()
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		//Запускаем конкурентное добавление значений в мап.
		go func(i int) {
			nums.add(i, i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	//Печатаем все пары ключ - значение.
	nums.getAll()
}
