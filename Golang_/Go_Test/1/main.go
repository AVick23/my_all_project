package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fullName, err := readFullName(reader)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if confirmFullName(fullName) {
			if err := saveToFile(fullName); err != nil {
				fmt.Println("Ошибка при сохранении данных:", err)
			} else {
				fmt.Println("Хорошо, данные сохранены.")
			}
			break
		} else {
			fmt.Println("Тогда введите ФИО сначала.")
		}
	}
}

func readFullName(reader *bufio.Reader) (string, error) {
	fmt.Print("Введите ваше ФИО через пробел: ")
	fullName, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("ошибка при чтении ввода: %v", err)
	}
	fullName = strings.TrimSpace(fullName)

	parts := strings.Fields(fullName)
	if len(parts) != 3 {
		return "", fmt.Errorf("ошибка: Пожалуйста, введите ФИО в формате 'Фамилия Имя Отчество'")
	}

	re := regexp.MustCompile(`^[А-Яа-яЁё\s]+$`)
	if !re.MatchString(fullName) {
		return "", fmt.Errorf("ошибка: Пожалуйста, используйте только кириллицу без цифр и специальных символов")
	}

	return fullName, nil
}

func confirmFullName(fullName string) bool {
	surname, name, middleName := parseFullName(fullName)
	fmt.Printf("Ваше ФИО: %s %s %s? (Да/Нет): ", surname, name, middleName)

	var response string
	fmt.Scan(&response)
	response = strings.ToLower(response)
	return response == "да"
}

func parseFullName(fullName string) (string, string, string) {
	parts := strings.Fields(fullName)
	return parts[0], parts[1], parts[2]
}

func saveToFile(fullName string) error {
	file, err := os.Create("full_name.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(fullName)
	if err != nil {
		return err
	}
	return writer.Flush()
}
