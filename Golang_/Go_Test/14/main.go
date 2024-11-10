/*Обратный массив
Напишите функцию, которая принимает массив целых чисел и возвращает новый массив, содержащий те же элементы, но в обратном порядке.*/

package main

import (
	"fmt"
)

func main() {

	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(reverse(nums))
}

func reverse(nums []int) []int {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}
