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
	"database/sql"
	"os"

	"github.com/teris-io/cli"
	"github.com/zSnails/taskr/internal/command"
	"github.com/zSnails/taskr/internal/store"

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

	mngr := store.NewManager(db)
	defer mngr.Close()

	app := cli.New("Tool for creating tasks")
	app.WithAction(command.Action(mngr))
	app.WithCommand(command.Add(mngr)).WithCommand(command.Delete(mngr)).WithCommand(command.Toggle(mngr))

	app.WithOption(cli.NewOption("verbose", "Show verbose output").WithType(cli.TypeBool).WithChar('v'))
	app.WithOption(cli.NewOption("no-color", "Disable colored output").WithType(cli.TypeBool).WithChar('c'))
	app.WithOption(cli.NewOption("version", "Shows program version info").WithType(cli.TypeBool).WithChar('V'))
	app.WithOption(cli.NewOption("all", "Shows all tasks").WithType(cli.TypeBool).WithChar('a'))

	os.Exit(app.Run(os.Args, os.Stdout))
}
