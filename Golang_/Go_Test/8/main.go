package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Добро пожаловать в интерфейс командной строки калькулятора!")
	fmt.Println("Пожалуйста, выберите операцию:")
	fmt.Println("1. Сложение")
	fmt.Println("2. Вычитание")
	fmt.Println("3. Умножение")
	fmt.Println("4. Деление")

	var choice int
	fmt.Print("Введите номер операции: ")
	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		os.Exit(1)
	}

	var num1, num2 float64
	fmt.Print("Введите первое число: ")
	_, err = fmt.Scanln(&num1)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		os.Exit(1)
	}

	fmt.Print("Введите второе число: ")
	_, err = fmt.Scanln(&num2)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		os.Exit(1)
	}

	var result float64
	switch choice {
	case 1:
		result = num1 + num2
		fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, result)
	case 2:
		result = num1 - num2
		fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, result)
	case 3:
		result = num1 * num2
		fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, result)
	case 4:
		if num2 == 0 {
			fmt.Println("Ошибка: Деление на ноль")
			os.Exit(1)
		}
		result = num1 / num2
		fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, result)
	default:
		fmt.Println("Неверный выбор операции")
		os.Exit(1)
	}
	fmt.Println("Окно закроется через 5 секунд...")
	time.Sleep(5 * time.Second)
}
