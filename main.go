package main

import (
	// "context"
	"database/sql"
	"fmt"
	"os"
	"strconv"
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
			var tasks []manager.Task
			tasks, err = mngr.GetTasks()
			if err != nil {
				panic(err)
			}

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
				fmt.Printf("%v - %v: %v\n", task.ID, col.Sprint(carbonDate.DiffForHumans()), task.Description)
			}
			return 0
		},
	)

	app := cli.New("Tool for creating tasks").WithAction(
		func(args []string, options map[string]string) int {
			var tasks []manager.Task
			tasks, err = mngr.ValidByDate()
			if err != nil {
				panic(err)
			}

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
				fmt.Printf("[%v] %v: %v\n", task.ID, col.Sprint(carbonDate.DiffForHumans()), task.Description)
			}
			return 0
		}).WithCommand(addCommand).WithCommand(deleteCommand).WithCommand(allCommand)
	os.Exit(app.Run(os.Args, os.Stdout))
}
