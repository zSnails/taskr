package main

import (
	// "context"
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-module/carbon"
	"github.com/gookit/color"
	"github.com/zSnails/taskr/internal/manager"

	_ "github.com/mattn/go-sqlite3"
)

var (
	modeAdd      bool
	dataFile     string
	showAll      bool
	noColor      bool
	deletionMode bool
)

func init() {
	// flag.BoolVar(&modeAdd, "add", false, "TODO: give this a more detailed description")
	// flag.BoolVar(&showAll, "all", false, "TODO: give this a more detailed description")
	// flag.BoolVar(&noColor, "no-color", false, "TODO: give this a more detailed description")
	// flag.BoolVar(&deletionMode, "delete", false, "TODO: give this a more detailed description")
	// flag.Parse()

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

// unpack que va dijo rafita
func unpack(s []string, vars ...*string) {
	for i, str := range s {
		*vars[i] = str
	}
}

func main() {
	if noColor {
		color.Disable()
	}

	db, err := sql.Open("sqlite3", dataFile+"?cache=shared&mode=memory")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	mngr := manager.NewManager(db)
	defer mngr.Close()

	if modeAdd {
		var taskDate string
		var taskDescription string

		unpack(flag.Args(), &taskDate, &taskDescription)
		t, err := time.Parse("2006-01-02 15:4:5", taskDate)
		if err != nil {
			panic(err)
		}

		err = mngr.AddTask(t, taskDescription)
		if err != nil {
			panic(err)
		}
	} else if deletionMode {

		id, err := strconv.Atoi(flag.Args()[0])
		if err != nil {
			panic(err)
		}
		err = mngr.RemoveTask(id)
		if err != nil {
			panic(err)
		}
	} else {
		var tasks []manager.Task
		var err error
		if showAll {
			tasks, err = mngr.GetTasks()
		} else {
			tasks, err = mngr.ValidByDate()
		}
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
	}
}
