package main

import "fmt"

//Реализовать бинарный поиск встроенными методами языка.
//Сложность О(logn).
//Алгоритм работает с отсортированными (в возрастающем
//или убывающем порядке) массивами.

func binarySearch(arr []int, elem int) (int, bool) {
	maxIdx := len(arr) - 1
	//Если массив отсортирован в возрастающем порядке,
	//возвращаем функцию ascendingOrder.
	if elem > arr[0] && elem < arr[maxIdx] {
		return ascendingOrder(arr, elem, maxIdx)
	}
	//Если массив отсортирован в убывающем порядке,
	//возвращаем функцию descendingOrder.
	if elem < arr[0] && elem > arr[maxIdx] {
		return descendingOrder(arr, elem, maxIdx)
	}
	//Проверяем граничные значения из конца и начала
	//массива.
	if elem == arr[0] {
		return 0, true
	}
	if elem == arr[maxIdx] {
		return maxIdx, true
	}
	//Иначе возвращаем.
	return -1, false
}

//Функция для бинарного поиска в массиве, отсортированном
//в возрастающем порядке.
func ascendingOrder(arr []int, elem int, maxIdx int) (int, bool) {
	//Левая граница
	i := 0
	//Находим середину массива.
	mid := (i + maxIdx) / 2
	for i < maxIdx && arr[mid] != elem {
		//Если середина массива меньше,
		//чем искомый элемент, то двигаем левую границу
		if arr[mid] < elem {
			i++
		} else {
			//В ином случае двигаем правую границу до
			//середины рассматриваемого диапазона массива.
			maxIdx = mid
		}
		//Пересчитываем центр.
		mid = (i + maxIdx) / 2
	}
	//Если нашли элемент, возвращаем его.
	if arr[mid] == elem {
		return mid, true
	}
	return -1, false
}

//Функция для бинарного поиска в массиве, отсортированно в
//убывающем порядке.
func descendingOrder(arr []int, elem int, maxIdx int) (int, bool) {
	//Левая граница.
	i := 0
	//Середина массива.
	mid := (i + maxIdx) / 2
	for i < maxIdx && arr[mid] != elem {
		//Если значение из середины массива больше,
		//чем искомый элемент, то двигаем левую границу.
		if arr[mid] > elem {
			i++
		} else {
			//Иначе двигаем правую границу до середины
			//рассматриваемого диапазона массива.
			maxIdx = mid
		}
		//Пересчитываем центр.
		mid = (i + maxIdx) / 2
	}
	//Если нашли элемент, возвращаем его.
	if arr[mid] == elem {
		return mid, true
	}
	return -1, false
}

func main() {
	arrAsc := []int{-5, -2, 0, 8, 10, 22}
	idx, ok := binarySearch(arrAsc, -6)
	if ok {
		fmt.Printf("index = %d\n", idx)
	} else {
		fmt.Println("not found")
	}
}
