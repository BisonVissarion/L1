package main

import (
	"fmt"
)

func removeElement(slice []int, index int) []int {
	// Проверяем, находится ли индекс в допустимом диапазоне
	if index < 0 || index >= len(slice) {
		return slice // Возвращаем исходный слайс без изменений, если индекс недопустим
	}

	// Удаляем элемент из слайса
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	// Пример слайса
	slice := []int{1, 2, 3, 4, 5}

	// Индекс элемента, который нужно удалить (например, индекс 2 соответствует тройке)
	indexToRemove := 2

	// Удаляем элемент с заданным индексом
	slice = removeElement(slice, indexToRemove)

	// Выводим результат
	fmt.Println(slice)
}
