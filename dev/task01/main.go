package main

import (
	"errors"
	"fmt"
	"strings"
)

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от
//родительской структуры Human (аналог наследования).

// Константы, определяющие возможный вариант пола человека.
const (
	male   string = "male"
	female string = "female"
)

// Тип данных Human.
type Human struct {
	age    int
	gender string
	name   string
}

// Тип данных Action.
type Action struct {
	// Нам необходимо встроить тип данных Human в Action
	// для наследования методов типа Human.
	//Т.е. мы встраивааем родительский тип Human в дочерний тип Action.
	Human
}

// Функция-конструктор, возвращающая новый экземпляр структуры Human.
func New(age int, gender, name string) (*Human, error) {
	//Имитация бурной деятельности...
	if age <= 0 {
		return nil, errors.New("возраст человека не может быть меньше нуля")
	}

	genderCheck := false
	if strings.ToLower(gender) == male || strings.ToLower(gender) == female {
		genderCheck = true
	}

	if genderCheck == false {
		genderErr := fmt.Sprintf("выберите один из предложенных полов: %s или %s", male, female)
		return nil, errors.New(genderErr)
	}

	if name == "" {
		return nil, errors.New("имя не может быть пустым")
	}

	h := Human{
		age:    age,
		gender: gender,
		name:   name,
	}
	return &h, nil
}

// Метод прогулка.
func (h *Human) Walk() {
	fmt.Printf("%s гуляет\n", h.name)
}

func main() {
	//Инициализируем новый экземпляр Human.
	human, err := New(18, "male", "Ivan")
	if err != nil {
		fmt.Println(err)
		return
	}
	//Инициализируем новый экземпляр Action и передаем по указателю human.
	action := Action{
		Human: *human,
	}
	//Вызываем метод Walk от структуры нашего экземпляра action.
	//что наглядно демонстрирует встраивание (наследование).
	action.Walk() //Ivan гуляет

	//Также мы унаследованили все поля стуктуры human.
	action.name = "Dima"
	action.Walk() //Dima гуляет
	//Но следует помнить, что это не затрагивает экземпляр
	//структуры human.
	human.Walk() //Ivan
}
