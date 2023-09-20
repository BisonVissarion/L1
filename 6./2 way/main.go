//Использование контекста для отмены выполнения

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина завершена")
				return
			default:
				fmt.Println("Горутина работает")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// Главный поток продолжает работу
	// ...

	// Отменяем выполнение горутины
	cancel()
	time.Sleep(2 * time.Second) // Подождать для наблюдения завершения
}
