package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "myCLI",
		Short: "My CLI is a simple example",
		Long:  `A longer description of my CLI application.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print("Введите 'Привет': ")

			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			if input == "Привет" {
				fmt.Println("Привет и добро пожаловать в моё CLI приложение")
			} else {
				fmt.Println("Извините, я не понимаю. Попробуйте ввести 'Привет'.")
			}

			fmt.Println("Окно закроется через 5 секунд...")
			time.Sleep(5 * time.Second)
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
