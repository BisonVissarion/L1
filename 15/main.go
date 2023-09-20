package main //Если строка v короче 100 символов, это вызовет панику

import (
	"fmt"
	"strings"
)

// Функция createHugeString создает и возвращает большую строку
func createHugeString(size int) string {
	// Создаем строку, повторяя "A" нужное количество раз
	return strings.Repeat("A", size)
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	if len(v) >= 100 {
		justString = v[:100]
	} else {
		justString = v
	}
}

func main() {
	someFunc()
	fmt.Println(justString)
}
