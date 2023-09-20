package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, num := range numbers {
		// Для каждого числа создаем новый горутин и увеличиваем счетчик WaitGroup
		wg.Add(1)
		go func(x int) {
			defer wg.Done() // Уменьшаем счетчик WaitGroup при завершении горутины
			square := x * x
			fmt.Printf("%d^2 = %d\n", x, square)
		}(num)
	}

	// Ждем, пока все горутины завершат свою работу
	wg.Wait()
}
