/*Проверка на чётность
Напишите функцию, которая принимает целое число и возвращает true, если это число чётное, и false, если нечётное.*/

package main

import (
	"fmt"
)

func main() {

	var num int

	fmt.Print("Введите целое число: ")
	nums, err := fmt.Scanln(&num)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		fmt.Println("Пожалуйста, введите именно целое число.")
		return
	}

	if nums%2 == 0 {
		fmt.Println("Число чётное")
	} else {
		fmt.Println("Число нечётное")
	}
}
