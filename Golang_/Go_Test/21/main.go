package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func count(from, to int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := from; i <= to; i++ {
		// Здесь ничего не делаем, просто считаем
	}
}

func main() {
	const limit int64 = 1000000000

	// Получаем количество доступных процессоров
	numGoroutines := int64(runtime.NumCPU())

	// Выводим количество создаваемых горутин
	fmt.Printf("Количество создаваемых горутин: %d\n", numGoroutines)

	// Запоминаем время начала выполнения
	start := time.Now()

	var wg sync.WaitGroup
	step := limit / numGoroutines

	for i := int64(0); i < numGoroutines; i++ {
		from := i*step + 1
		to := (i + 1) * step
		if i == numGoroutines-1 {
			to = limit // Последняя горутина должна считать до лимита
		}
		wg.Add(1)
		go count(from, to, &wg)
	}

	wg.Wait() // Ждем завершения всех горутин

	// Запоминаем время окончания выполнения
	elapsed := time.Since(start)

	// Выводим время выполнения в секундах
	fmt.Printf("Счет до %d завершен.\n", limit)
	fmt.Printf("Время выполнения: %.2f секунд\n", elapsed.Seconds())
}
