// Copyright (C) 2023  Aaron González

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

	"github.com/teris-io/cli"
	"github.com/zSnails/taskr/internal/store"
)

func Forget(manager *store.Manager) cli.Command {
	return cli.NewCommand("forget", "Forget a reminder").WithShortcut("ff").WithArg(
		cli.NewArg("id", "Reminder id to delete").WithType(cli.TypeInt),
	).WithAction(func(args []string, options map[string]string) int {
		err := manager.RemoveReminder(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			return 1
		}

		return 0
	})
}
