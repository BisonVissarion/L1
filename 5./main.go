package main

import (
	"fmt"
	"time"
)

func main() {
	// Создаем канал для передачи данных
	dataChannel := make(chan int)

	// Запускаем горутину для отправки данных в канал
	go func() {
		for i := 1; ; i++ {
			dataChannel <- i
			time.Sleep(time.Second) // Ожидание 1 секунду между отправками
		}
	}()

	// Задаем N секунд для работы программы
	N := 5

	// Создаем таймер на N секунд
	timer := time.NewTimer(time.Duration(N) * time.Second)

	// Ожидание данных из канала или срабатывания таймера
	select {
	case data := <-dataChannel:
		fmt.Printf("Прочитано из канала: %d\n", data)
	case <-timer.C:
		fmt.Printf("Программа завершена после %d секунд.\n", N)
		return
	}
}
