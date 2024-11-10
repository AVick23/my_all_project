package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Путь к программам
	paths := []string{
		`C:\Program Files\Mozilla Firefox\firefox.exe`,
		`C:\Users\ilege\OneDrive\Рабочий стол\v2rayN-With-Core\v2rayN-With-Core\v2rayN.exe`,
		`C:\Users\ilege\AppData\Local\Programs\Microsoft VS Code\Code.exe`,
	}

	// Запуск программ
	for _, path := range paths {
		err := exec.Command(path).Start()
		if err != nil {
			fmt.Printf("Ошибка при запуске %s: %s\n", path, err)
		} else {
			fmt.Printf("Запущена программа: %s\n", path)
		}
	}
}
