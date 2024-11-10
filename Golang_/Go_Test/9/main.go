package main

import (
	"fmt"
	"time"
)

func sumSequential(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func main() {
	n := 100000000 // 100 миллионов
	start := time.Now()
	result := sumSequential(n)
	elapsed := time.Since(start)

	fmt.Printf("Сумма: %d\n", result)
	fmt.Printf("Время выполнения (последовательно): %s\n", elapsed)
}
