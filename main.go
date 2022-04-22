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
package main

import (
	// "context"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-module/carbon"
	"github.com/gookit/color"
	"github.com/teris-io/cli"
	"github.com/zSnails/taskr/internal/manager"

	_ "github.com/mattn/go-sqlite3"
)

var dataFile string

func init() {
	// Check if there is a taskr folder in %APPDATA%
	dataDir, err := os.UserCacheDir()
	if err != nil {
		panic(err)
	}
	dataFile = dataDir + "/taskr/data.db"

	_, err = os.Stat(dataDir + "/taskr")
	if os.IsNotExist(err) {
		err = os.Mkdir(dataDir+"/taskr", os.ModeDir)
		if err != nil {
			panic(err)
		}
	}
	// check if the db file exists
	_, err = os.Stat(dataDir + "/taskr/data.db")
	if os.IsNotExist(err) {
		_, err = os.Create(dataDir + "/taskr/data.db")
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	db, err := sql.Open("sqlite3", dataFile+"?cache=shared&mode=memory")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	mngr := manager.NewManager(db)
	defer mngr.Close()

	addCommand := cli.NewCommand("new", "Create a new task").WithShortcut("n").WithArg(
		cli.NewArg("date", "Task date").WithType(cli.TypeString),
	).WithArg(
		cli.NewArg("description", "Task description").WithType(cli.TypeString),
	).WithAction(func(args []string, options map[string]string) int {
		date := args[0]
		t, err := time.Parse("2006-01-02 15:4:5", date)
		if err != nil {
			os.Stderr.WriteString(err.Error() + "\n")
			return 1
		}

		desc := args[1]

		err = mngr.AddTask(t, desc)
		if err != nil {
			os.Stderr.WriteString(err.Error() + "\n")
			return 1
		}
		return 0
	})

	deleteCommand := cli.NewCommand("delete", "Delete a task").WithShortcut("d").WithArg(
		cli.NewArg("id", "Task id to delete").WithType(cli.TypeInt),
	).WithAction(func(args []string, options map[string]string) int {

		id, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}

		err = mngr.RemoveTask(id)
		if err != nil {
			panic(err)
		}

		return 0
	})

	allCommand := cli.NewCommand("all", "Shows all tasks").WithAction(
		func(args []string, options map[string]string) int {
			tasks, err := mngr.GetTasks()
			if err != nil {
				panic(err)
			}
			printTasks(tasks)
			return 0
		},
	)

	app := cli.New("Tool for creating tasks").WithAction(
		func(args []string, options map[string]string) int {
			tasks, err := mngr.ValidByDate()
			if err != nil {
				return 1
			}
			printTasks(tasks)
			return 0
		}).WithCommand(addCommand).WithCommand(deleteCommand).WithCommand(allCommand)
	os.Exit(app.Run(os.Args, os.Stdout))
}

func printTasks(tasks []manager.Task) {
	str := strings.Builder{}
	last := len(tasks) - 1
	for i, task := range tasks {
		carbonDate := carbon.CreateFromDateTime(
			task.Date.Year(),
			int(task.Date.Month()),
			task.Date.Day(),
			task.Date.Hour(),
			task.Date.Minute(),
			task.Date.Second(),
		)
		col := color.New(color.FgLightGreen)
		end := "\n"
		if last == i {
			end = ""
		}
		str.WriteString(fmt.Sprintf("[%v] %v: %v%v", task.ID, col.Sprint(carbonDate.DiffForHumans()), task.Description, end))
	}
	println(str.String())
}