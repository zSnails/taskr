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

func printTasks(tasks []store.Task, verbose bool) {
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
			fmt.Fprintf(&str, "[%d] ", task.ID)
		}

		fmt.Fprintf(&str, "%s: %s\n", diff, task.Description)
	}

	fmt.Print(str.String())
}

func printReports(tasks []store.Task, verbose bool) {
	str := strings.Builder{}
	fmt.Fprintln(&str, "The following tasks have expired but have not\nbeen marked as done:")
	for _, report := range tasks {
		carbonDate := carbon.CreateFromDateTime(
			report.Date.Year(),
			int(report.Date.Month()),
			report.Date.Day(),
			report.Date.Hour(),
			report.Date.Minute(),
			report.Date.Second(),
		)

		col := color.New(color.FgRed)
		diff := col.Render(carbonDate.SetLanguage(lang).DiffForHumans())
		fmt.Fprint(&str, "\t- ")
		if verbose {
			fmt.Fprintf(&str, "[%d] ", report.ID)
		}

		fmt.Fprintf(&str, "%s: %s\n", diff, report.Description)
	}

	fmt.Print(str.String())
}

func printReminders(reminders []store.Reminder, verbose bool) {
	str := strings.Builder{}
	fmt.Fprintln(&str, "Remember to do these today:")
	for _, reminder := range reminders {
		carbonDate := carbon.CreateFromTime(
			reminder.Hour.Hour(),
			reminder.Hour.Minute(),
			reminder.Hour.Second(),
		)
		col := color.New(color.FgLightBlue)
		diff := col.Render(carbonDate.SetLanguage(lang).DiffForHumans())
		fmt.Fprintf(&str, "\t- ")
		if verbose {
			fmt.Fprintf(&str, "[%d] ", reminder.ID)
		}

		fmt.Fprintf(&str, "%s: %s\n", diff, reminder.Description)
	}

	fmt.Print(str.String())
}
