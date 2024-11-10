/*Поиск максимального элемента
Напишите функцию, которая принимает массив целых чисел и возвращает максимальный элемент из этого массива.*/

package main

import (
	"fmt"
)

func main() {
	var nums = []int{2, 7, 11, 15}
	fmt.Println(max(nums))
}

func max(nums []int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max
}
