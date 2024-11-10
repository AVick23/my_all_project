package main

import (
	"fmt"
)

func main() {
	var a = []int{1, 2, 3}
	var b = []int{1, 2, 3}
	fmt.Println(Equal(a, b))
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
