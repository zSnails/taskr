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

	"github.com/gookit/color"
	"github.com/teris-io/cli"
	"github.com/zSnails/taskr/internal/store"
)

var (
	BuildUser  = ""
	Version    = ""
	CommitHash = ""
	license    = "Copyright (C) 2022 zSnails\nThis program comes with ABSOLUTELY NO warranty\nThis is free software, and you are welcome to redistribute it\nunder certain conditions."
)

func Action(mngr *store.Manager) cli.Action {
	return func(args []string, options map[string]string) int {
		if _, showVersion := options["version"]; showVersion {
			fmt.Printf("taskr (built by %v) %v %v\n",
				BuildUser,
				Version,
				CommitHash,
			)
			println(license)
			return 0
		}

		if _, noColor := options["no-color"]; noColor {
			color.Disable()
		}
		_, verbose := options["verbose"]
		var (
			tasks []store.Task
			err   error
		)

		if _, showAll := options["all"]; showAll {
			tasks, err = mngr.All()
		} else {
			tasks, err = mngr.Valid()
		}
		if err != nil {
			println(err.Error())
			return 1
		}
		PrintTasks(tasks, verbose)
		return 0
	}
}
