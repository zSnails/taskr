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
	"github.com/teris-io/cli"
	"github.com/zSnails/taskr/internal/store"
)

func All(manager *store.Manager) cli.Command {
    return cli.NewCommand("all", "Shows all tasks").WithAction(
		func(args []string, options map[string]string) int {

              _, verbose := options["verbose"]

			tasks, err := manager.GetTasks()
			if err != nil {
				panic(err)
			}
			PrintTasks(tasks, verbose)
               return 0
		},
	)

}
