package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var AddCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Добавить задачу",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := args[0]
		tasks := viper.GetStringSlice("tasks")
		tasks = append(tasks, task)
		viper.Set("tasks", tasks)

		if err := viper.WriteConfig(); err != nil {
			fmt.Println("Ошибка записи кофигурационного файла:", err)
			os.Exit(1)
		}

		fmt.Printf("Задача '%s' добавлена!\n ", task)
	},
}
