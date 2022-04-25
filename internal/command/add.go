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
	"os"
	"time"
)

func Add(manager *store.Manager) cli.Command {
	return cli.NewCommand("new", "Create a new task").WithShortcut("n").WithArg(
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

		err = manager.AddTask(t, desc)
		if err != nil {
			os.Stderr.WriteString(err.Error() + "\n")
			return 1
		}
		return 0
	})
}
