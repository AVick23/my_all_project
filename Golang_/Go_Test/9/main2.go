package main

import (
	"fmt"
	"sync"
	"time"
)

func sumPart(start, end int, wg *sync.WaitGroup, result *int) {
	defer wg.Done()
	sum := 0
	for i := start; i <= end; i++ {
		sum += i
	}
	*result += sum
}

func sumConcurrent(n int, numGoroutines int) int {
	var wg sync.WaitGroup
	result := 0
	partSize := n / numGoroutines

	for i := 0; i < numGoroutines; i++ {
		start := i*partSize + 1
		end := (i + 1) * partSize
		if i == numGoroutines-1 {
			end = n // последний горутина обрабатывает остаток
		}
		wg.Add(1)
		go sumPart(start, end, &wg, &result)
	}

	wg.Wait()
	return result
}

func main() {
	n := 100000000     // 100 миллионов
	numGoroutines := 8 // количество горутин
	start := time.Now()
	result := sumConcurrent(n, numGoroutines)
	elapsed := time.Since(start)

	fmt.Printf("Сумма: %d\n", result)
	fmt.Printf("Время выполнения (параллельно): %s\n", elapsed)
}
