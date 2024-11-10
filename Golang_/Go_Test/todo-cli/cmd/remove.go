package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RemoveCmd = &cobra.Command{
	Use:   "remove [index]",
	Short: "Удалить задачу по индексу",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index := args[0]
		tasks := viper.GetStringSlice("tasks")

		if len(tasks) == 0 {
			fmt.Println("Нет задач для удаления.")
			return
		}

		if idx, err := strconv.Atoi(index); err == nil && idx < len(tasks) {
			tasks = append(tasks[:idx], tasks[idx+1:]...)
			viper.Set("tasks", tasks)

			if err := viper.WriteConfig(); err != nil {
				fmt.Println("Ошибка записи конфигурационного файла:", err)
				os.Exit(1)
			}

			fmt.Printf("Задача с индексом %s удалена!\n", index)
		} else {
			fmt.Println("Неверный индекс задачи.")
		}
	},
}
