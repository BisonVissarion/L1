package main

import (
	"fmt"
)

func main() {
	// Создаем каналы для передачи данных
	inputChannel := make(chan int)
	outputChannel := make(chan int)

	// Горутина для умножения чисел на 2 и отправки в outputChannel
	go func() {
		for {
			num := <-inputChannel
			result := num * 2
			outputChannel <- result
		}
	}()

	// Горутина для вывода результатов в stdout
	go func() {
		for {
			result := <-outputChannel
			fmt.Println(result)
		}
	}()

	// Пример массива чисел для конвейера
	numbers := []int{1, 2, 3, 4, 5}

	// Отправляем числа из массива в inputChannel
	for _, num := range numbers {
		inputChannel <- num
	}

	// Закрываем inputChannel, чтобы сигнализировать о завершении второй горутине
	close(inputChannel)

	// Ждем завершения программы (в данном случае, это бесконечное ожидание)
	select {}
}
