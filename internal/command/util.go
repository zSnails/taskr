package command

import (
	"strings"
	"fmt"

	"github.com/golang-module/carbon"
	"github.com/gookit/color"
	"github.com/zSnails/taskr/internal/store"
)

func printTasks(tasks []store.Task) {
	str := strings.Builder{}
	for _, task := range tasks {
		carbonDate := carbon.CreateFromDateTime(
			task.Date.Year(),
			int(task.Date.Month()),
			task.Date.Day(),
			task.Date.Hour(),
			task.Date.Minute(),
			task.Date.Second(),
		)

		col := color.New(color.FgLightGreen)
		str.WriteString(fmt.Sprintf("[%v] %v: %v\n", task.ID, col.Sprint(carbonDate.DiffForHumans()), task.Description))
	}
	print(str.String())
}
