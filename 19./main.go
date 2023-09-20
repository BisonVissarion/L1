package main

import (
	"fmt"
)

// ReverseString переворачивает строку с учетом символов Unicode
func ReverseString(input string) string {
	runes := []rune(input)
	length := len(runes)

	// Переворачиваем строку
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	input := "главрыба"
	reversed := ReverseString(input)
	fmt.Println(reversed)
}
