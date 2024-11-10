package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "myCLI",
		Short: "My CLI is a simple example",
		Long:  `A longer description of my CLI application.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Привет, это мое первое просто CLI-приложение!")

			fmt.Println("Окно закроется через 5 секунд...")
			time.Sleep(5 * time.Second)
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
