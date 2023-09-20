package main

import (
	"fmt"
)

func main() {
	// Исходная последовательность температурных значений
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Создаем map для группировки температурных значений с шагом в 10 градусов
	groupedTemperatures := make(map[int][]float64)

	// Группируем температуры по шагу в 10 градусов
	for _, temp := range temperatures {
		key := int(temp/10) * 10 // Вычисляем ключ для группировки
		groupedTemperatures[key] = append(groupedTemperatures[key], temp)
	}

	// Выводим группы температурных значений
	for key, group := range groupedTemperatures {
		fmt.Printf("Группа %d градусов:\n", key)
		fmt.Println(group)
	}
}
