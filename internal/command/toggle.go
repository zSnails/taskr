// copyright (c) 2023  aaron gonzález

// this program is free software: you can redistribute it and/or modify
// it under the terms of the gnu general public license as published by
// the free software foundation, either version 3 of the license, or
// (at your option) any later version.

// this program is distributed in the hope that it will be useful,
// but without any warranty; without even the implied warranty of
// merchantability or fitness for a particular purpose.  see the
// gnu general public license for more details.

// you should have received a copy of the gnu general public license
// along with this program.  if not, see <https://www.gnu.org/licenses/>.
package command

import (
	"fmt"
	"os"

	"github.com/teris-io/cli"
	"github.com/zSnails/taskr/internal/store"
)

func Toggle(mngr *store.Manager) cli.Command {
	return cli.NewCommand("toggle", "Toggles the status of a task").WithShortcut("tg").WithArg(
		cli.NewArg("id", "The id of the task to modify"),
	).WithAction(func(args []string, options map[string]string) int {
		err := mngr.ToggleStatus(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			return 1
		}

		return 0
	})
}
