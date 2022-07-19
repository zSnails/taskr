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
	"time"

	"github.com/teris-io/cli"
	"github.com/zSnails/taskr/internal/store"
)

func Remind(manager *store.Manager) cli.Command {
    return cli.NewCommand("remind", "Create a new reminder").WithShortcut("re").WithArg(
        cli.NewArg("time", "Reminder time").WithType(cli.TypeString),
    ).WithArg(
        cli.NewArg("description", "Reminder description").WithType(cli.TypeString),
    ).WithAction(func(args []string, options map[string]string) int {
        tm := args[0]
        t, err := time.Parse("15:4:5", tm)
        if err != nil {
            fmt.Fprintf(os.Stderr, err.Error())
            return 1
        }

        desc := args[1]
        err = manager.AddReminder(t, desc)
        if err != nil {
            fmt.Fprintf(os.Stderr, err.Error())
            return 1
        }

        return 0
    })
}
