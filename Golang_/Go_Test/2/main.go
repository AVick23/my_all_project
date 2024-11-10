package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите слова через пробел: ")

	var words []string
	for {
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "" {
			break
		}
		words = append(words, strings.Fields(text)...)
	}
	for i := len(words) - 1; i >= 0; i-- {
		fmt.Println(words[i])
	}
}
