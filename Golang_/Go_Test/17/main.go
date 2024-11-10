package main

import (
	"fmt"
)

func main() {
	var age, name = add(13, 14, "Ивано", "Иванов")
	fmt.Print(age, "\n", name)
}

func add(a, b int, firsname, lastname string) (int, string) {
	var z = a + b
	var fullName = firsname + " " + lastname
	return z, fullName
}
