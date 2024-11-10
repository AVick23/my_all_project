package main

import "fmt"

func main() {
	type Student struct {
		Name string
		Age  int
	}

	var pa *Student
	pa = new(Student)
	pa.Name = "Иван"

	fmt.Println(pa.Name)
}
