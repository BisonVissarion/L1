package main

import (
	"fmt"
	"sync"
)

// Counter - структура счетчика
type Counter struct {
	mu    sync.Mutex
	value int
}

// Increment - метод для инкрементирования счетчика
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// GetValue - метод для получения текущего значения счетчика
func (c *Counter) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	// Создаем счетчик
	counter := &Counter{}

	// Запускаем несколько горутин для инкрементирования счетчика
	numWorkers := 5
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Выводим итоговое значение счетчика
	fmt.Printf("Итоговое значение счетчика: %d\n", counter.GetValue())
}
