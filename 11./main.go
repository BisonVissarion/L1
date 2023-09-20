package main

import (
	"fmt"
)

func intersectSets(set1, set2 map[int]bool) map[int]bool {
	result := make(map[int]bool)

	// Перебираем элементы из первого множества
	for item := range set1 {
		// Если элемент также присутствует во втором множестве, добавляем его в результат
		if set2[item] {
			result[item] = true
		}
	}

	return result
}

func main() {
	// Пример двух неупорядоченных множеств
	set1 := map[int]bool{1: true, 2: true, 3: true, 4: true}
	set2 := map[int]bool{3: true, 4: true, 5: true, 6: true}

	// Вызываем функцию для выполнения пересечения множеств
	intersection := intersectSets(set1, set2)

	// Выводим результат
	fmt.Println("Результат пересечения множеств:")
	for item := range intersection {
		fmt.Println(item)
	}
}
