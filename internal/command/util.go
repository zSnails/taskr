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
	"os"
	"strings"

	"github.com/gookit/color"
	"github.com/uniplaces/carbon"
	"github.com/zSnails/taskr/internal/store"
)

func PrintTasks(tasks []store.Task, verbose bool) int {
	str := strings.Builder{}
	for _, task := range tasks {
		carbonDate := carbon.NewCarbon(task.Date)
		var col color.Style
		if task.Done {
			col = color.New(color.FgLightGreen, color.OpStrikethrough)
		} else if task.Expired {
			col = color.New(color.FgRed)
		} else {
			col = color.New(color.FgYellow)
		}
		diff, err := carbonDate.DiffForHumans(nil, true, false, false)
		if err != nil {
			os.Stderr.WriteString(err.Error())
			return 1
		}

		diff = col.Sprint(diff)
		if verbose {
			str.WriteString(fmt.Sprintf("[%v] ", task.ID))
		}

		str.WriteString(fmt.Sprintf("%v: %v\n", diff, task.Description))
	}
	print(str.String())
	return 0
}
