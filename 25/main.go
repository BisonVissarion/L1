package main

import (
	"fmt"
	"time"
)

// MySleep приостанавливает выполнение программы на указанное количество миллисекунд
func MySleep(milliseconds time.Duration) {
	<-time.After(milliseconds * time.Millisecond)
}

func main() {
	fmt.Println("Начало работы программы")
	MySleep(2000) // Приостановить выполнение на 2 секунды
	fmt.Println("Прошло 2 секунды")
}
