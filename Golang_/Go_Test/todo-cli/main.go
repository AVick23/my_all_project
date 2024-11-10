package main

import (
	"fmt"
	"os"

	"todo-cli/cmd" // Импортируйте пакет cmd

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	cobra.OnInitialize(initConfig)

	var rootCmd = &cobra.Command{
		Use:   "todo",
		Short: "Приложение для управления задачами",
		Long:  `Приложение для управления задачами, позволяющее добавлять, удалять и отображать задачи.`,
	}

	// Добавьте команды
	rootCmd.AddCommand(cmd.AddCmd)
	rootCmd.AddCommand(cmd.RemoveCmd)
	rootCmd.AddCommand(cmd.ListCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	viper.SetConfigName("config") // имя конфигурационного файла без расширения
	viper.SetConfigType("json")   // тип конфигурационного файла
	viper.AddConfigPath(".")      // путь к конфигурационному файлу

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Ошибка чтения конфигурационного файла:", err)
	}
}
