package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	// Создаем канал для передачи результатов вычислений
	resultChannel := make(chan int)

	// Создаем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Запускаем горутины для вычисления квадратов и отправки их в канал
	for _, num := range numbers {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			square := x * x
			resultChannel <- square
		}(num)
	}

	// Запускаем горутину для закрытия канала после завершения всех вычислений
	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	// Суммируем квадраты чисел из канала
	sum := 0
	for square := range resultChannel {
		sum += square
	}

	fmt.Printf("Сумма квадратов: %d\n", sum)
}
