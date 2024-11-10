package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func main() {
	var input string

	var rootCmd = &cobra.Command{
		Use:   "checknumber",
		Short: "Проверка четности числа",
		Run: func(cmd *cobra.Command, args []string) {
			if input == "" {
				reader := bufio.NewReader(os.Stdin)
				fmt.Print("Введите какое-то целое число: ")
				input, _ = reader.ReadString('\n')
				input = input[:len(input)-1] // Убираем символ новой строки
			}

			// Пробуем преобразовать ввод в целое число
			a, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Ошибка ввода. Пожалуйста, введите именно целое число.")
				return
			}

			// Проверяем четность или нечетность
			if a%2 == 0 {
				fmt.Println("Введённое вами число является чётным.")
			} else {
				fmt.Println("Введённое вами число является нечётным.")
			}
		},
	}

	rootCmd.Flags().StringVarP(&input, "input", "i", "", "Введите целое число")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
