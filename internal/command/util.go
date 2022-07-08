// Copyright (C) 2022  Aaron González

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
package command

import (
	"fmt"
	"strings"

	"github.com/golang-module/carbon"
	"github.com/gookit/color"
	"github.com/zSnails/taskr/internal/resources"
	"github.com/zSnails/taskr/internal/store"
)

var (
    lang = carbon.NewLanguage()
)

func init() {
    lang.SetResources(resources.Resources)
}

func printTasks(tasks []store.Task, verbose bool) int {
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
		var col color.Style
		if task.Done {
			col = color.New(color.FgLightGreen, color.OpStrikethrough)
		} else if task.Expired {
			col = color.New(color.FgRed)
		} else {
			col = color.New(color.FgYellow)
		}
        diff := col.Render(carbonDate.SetLanguage(lang).DiffForHumans())
		if verbose {
            fmt.Fprintf(&str, "[%v] ", task.ID)
		}

        fmt.Fprintf(&str, "%v: %v\n", diff, task.Description)
	}

	fmt.Print(str.String())
	return 0
}

func printReminders(tasks []store.Task, verbose bool) {
    str := strings.Builder{}
    fmt.Fprintln(&str, "The following tasks have expired but have not\nbeen marked as done:")
    for _, reminder := range tasks {
        carbonDate := carbon.CreateFromDateTime(
            reminder.Date.Year(),
            int(reminder.Date.Month()),
            reminder.Date.Day(),
            reminder.Date.Hour(),
            reminder.Date.Minute(),
            reminder.Date.Second(),
        )

        col := color.New(color.FgRed)
        diff := col.Render(carbonDate.SetLanguage(lang).DiffForHumans())
        fmt.Fprint(&str, "\t- ")
        if verbose {
            fmt.Fprintf(&str, "[%v] ", reminder.ID)
        }

        fmt.Fprintf(&str, "%v: %v\n", diff, reminder.Description)
    }

    fmt.Print(str.String())
}
