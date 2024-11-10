/*Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.



Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]


Constraints:

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
Only one valid answer exists.


Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?*/

package main

import (
	"fmt"
)

func main() {
	var nums = []int{2, 7, 11, 15}
	var target = 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	// создаем хэш-таблицу для хранения значений и их индексов
	numMap := make(map[int]int)
	for i, num := range nums {
		// вычисляем разницу между текущим числом и целевым
		diff := target - num
		// ищем разницу в хэш-таблице
		if j, ok := numMap[diff]; ok {
			return []int{j, i}
		}
		// добавляем текущее значение в хэш-таблицу
		numMap[num] = i
	}
	return nil
}
