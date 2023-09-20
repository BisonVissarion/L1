package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	// Количество воркеров
	numWorkers := 3

	// Создаем канал для передачи данных от главного потока к воркерам
	dataChannel := make(chan int)

	// Создаем WaitGroup для ожидания завершения всех воркеров
	var wg sync.WaitGroup

	// Создаем канал для обработки сигнала завершения (Ctrl+C)
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, syscall.SIGINT, syscall.SIGTERM)

	// Запускаем воркеры.
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for {
				// Чтение данных из канала и вывод на stdout
				select {
				case data, ok := <-dataChannel:
					if !ok {
						fmt.Printf("Воркер %d завершил работу.\n", workerID)
						return
					}
					fmt.Printf("Воркер %d получил данные: %d\n", workerID, data)
				case <-sigChannel:
					fmt.Printf("Воркер %d получил сигнал завершения.\n", workerID)
					return
				}
			}
		}(i)
	}

	// Главный поток записывает данные в канал
	go func() {
		for i := 1; ; i++ {
			dataChannel <- i
		}
	}()

	// Ожидание сигнала завершения (Ctrl+C)
	<-sigChannel

	// Закрытие канала данных и ожидание завершения всех воркеров
	close(dataChannel)
	wg.Wait()

	fmt.Println("Главный поток завершил работу.")
}
