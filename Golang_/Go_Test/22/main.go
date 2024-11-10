/*
You are required to develop a program utilizing structures that would enable the selection of a user name, as well as the determination of the amount by which their salary should be increased or decreased.*/

package main

import (
	"fmt"
)

type Employee struct {
	name   string
	salary int
}

func (e *Employee) IncreaseSalary(amount int) {
	e.salary += amount
}

func (e *Employee) DecreaseSalary(amount int) {
	e.salary -= amount
}

func main() {
	var employee Employee
	var decrease int
	var increase int

	fmt.Println("Введите имя пользователя: ")
	fmt.Scan(&employee.name)
	fmt.Println("Введите зарплату пользователя: ")
	fmt.Scan(&employee.salary)
	fmt.Println("Введенное имя пользователя: ", employee.name)
	fmt.Println("Введенная зарплата пользователя: ", employee.salary)

	fmt.Println("Введите сумму, на которую хотите увеличить зарплату: ")
	fmt.Scan(&increase)
	employee.IncreaseSalary(increase)
	fmt.Println("Новая зарплата пользователя: ", employee.salary)

	fmt.Println("Введите сумму, на которую хотите уменьшить зарплату: ")
	fmt.Scan(&decrease)
	employee.DecreaseSalary(decrease)
	fmt.Println("Новая зарплата пользователя: ", employee.salary)
}
