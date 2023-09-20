package main

import (
	"fmt"
	"sort"
)

func main() {
	// Массива для сортировки
	arr := []int{6, 3, 9, 8, 1, 2, 7, 5, 4}

	// Используем функцию sort.Ints для сортировки массива в порядке возрастания
	sort.Ints(arr)

	// Выводим отсортированный массив
	fmt.Println(arr)
}
