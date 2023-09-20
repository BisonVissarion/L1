package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем мьютекс для синхронизации доступа к map
	var mu sync.Mutex

	// Создаем map, в которую будем записывать данные
	data := make(map[string]int)

	// Количество горутин для конкурентной записи
	numGoroutines := 10

	// Создаем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Запускаем горутины для конкурентной записи данных в map
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Заблокируем мьютекс для доступа к map
			mu.Lock()
			defer mu.Unlock()

			key := fmt.Sprintf("key%d", id)
			value := id * 10

			// Записываем данные в map
			data[key] = value

			fmt.Printf("Горутина %d записала в map: %s -> %d\n", id, key, value)
		}(i)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Выводим содержимое map после всех записей
	fmt.Println("Содержимое map:")
	for key, value := range data {
		fmt.Printf("%s -> %d\n", key, value)
	}
}
