package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Показать все задачи",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := viper.GetStringSlice("tasks")

		if len(tasks) == 0 {
			fmt.Println("Нет задач.")
			return
		}

		fmt.Println("Список задач:")
		for i, task := range tasks {
			fmt.Printf("%d: %s\n", i, task)
		}
	},
}
