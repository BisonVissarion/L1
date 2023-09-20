package main

import (
	"fmt"
	"strings"
)

// ReverseWords переворачивает слова в строке
func ReverseWords(input string) string {
	// Разделяем строку на слова
	words := strings.Fields(input)

	// Переворачиваем порядок слов
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Объединяем слова обратно в строку с пробелами между ними
	return strings.Join(words, " ")
}

func main() {
	input := "snow dog sun — sun dog snow"
	reversed := ReverseWords(input)
	fmt.Println(reversed)
}
