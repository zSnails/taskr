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
package store

import (
	"time"
)


func (m *Manager) All() (tasks []Task, err error) {
	rows, err := m.db.Query("SELECT id, taskdate, description FROM tasks")
	if err != nil {
		return
	}
	for rows.Next() {
		task := Task{}
		var _date string
		rows.Scan(&task.ID, &_date, &task.Description)
		task.Date, err = time.Parse("2006-01-02T15:4:5Z", _date)
		if err != nil {
			return
		}
		tasks = append(tasks, task)
	}
	return
}

func (m *Manager) Valid() (tasks []Task, err error) {
	rows, err := m.db.Query("SELECT id, taskdate, description FROM tasks WHERE date('now') < taskdate AND taskdate < date('now', '+7 days')")
	if err != nil {
		return
	}

	for rows.Next() {
		task := Task{}
		rows.Scan(&task.ID, &task.Date, &task.Description)
		tasks = append(tasks, task)
	}
	return
}
