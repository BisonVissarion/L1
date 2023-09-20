//Использование runtime.Goexit()

package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		defer fmt.Println("Эта строка не выполнится")
		fmt.Println("Горутина работает")
		runtime.Goexit()
	}()

	// Главный поток продолжает работу
	// ...
}
