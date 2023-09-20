package main

import (
	"fmt"
	"sort"
)

func main() {
	// Пример отсортированного среза для бинарного поиска
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}

	// Искомый элемент.
	target := 11

	// Используем функцию sort.Search для бинарного поиска
	index := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})

	// Проверяем, найден ли элемент
	if index < len(arr) && arr[index] == target {
		fmt.Printf("Элемент %d найден в позиции %d\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден\n", target)
	}
}
