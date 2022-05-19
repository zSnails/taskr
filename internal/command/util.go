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

func printTasks(tasks []store.Task, verbose bool) int {
	str := strings.Builder{}
    lang := carbon.NewLanguage()
    lang.SetResources(resources.Resources)
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

        diff := carbonDate.SetLanguage(lang).DiffForHumans()

		diff = col.Sprint(diff)
		if verbose {
            fmt.Fprintf(&str, "[%v] ", task.ID)
		}

        fmt.Fprintf(&str, "%v: %v\n", diff, task.Description)
	}

	fmt.Print(str.String())
	return 0
}
