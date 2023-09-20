package main

import (
	"fmt"
	"strings"
)

// IsUnique проверяет, что все символы в строке уникальные (регистронезависимо)
func IsUnique(input string) bool {
	// Приводим строку к нижнему регистру для регистронезависимой проверки
	input = strings.ToLower(input)

	// Создаем мапу для хранения символов
	charMap := make(map[rune]bool)

	// Перебираем символы строки
	for _, char := range input {
		// Если символ уже есть в мапе, возвращаем false
		if charMap[char] {
			return false
		}
		// Добавляем символ в мапу.
		charMap[char] = true
	}

	// Если все символы уникальны, возвращаем true
	return true
}

func main() {
	// Примеры строк для проверки
	str1 := "abcd"      // true
	str2 := "abCdefAaf" // false
	str3 := "aabcd"     // false

	// Проверяем строки и выводим результаты
	fmt.Printf("%s — %v\n", str1, IsUnique(str1))
	fmt.Printf("%s — %v\n", str2, IsUnique(str2))
	fmt.Printf("%s — %v\n", str3, IsUnique(str3))
}
