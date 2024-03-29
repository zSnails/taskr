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
package store

import (
	"time"
)

type Reminder struct {
	ID          int
	Hour        time.Time
	Description string
}

type Task struct {
	ID          int
	Date        time.Time
	Description string
	Done        bool
	Expired     bool
}

func (m *Manager) AllReminders() (reminders []Reminder, any bool, err error) {
	rows, err := m.db.Query("SELECT id, hour, description FROM reminders")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		reminder := Reminder{}
		rows.Scan(&reminder.ID, &reminder.Hour, &reminder.Description)
		reminders = append(reminders, reminder)
	}

	any = len(reminders) != 0

	return
}

func (m *Manager) AllTasks() (tasks []Task, err error) {
	rows, err := m.db.Query("SELECT id, taskdate, description, done, (taskdate < datetime('now', 'localtime')) FROM tasks")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		task := Task{}
		rows.Scan(&task.ID, &task.Date, &task.Description, &task.Done, &task.Expired)
		tasks = append(tasks, task)
	}

	return
}

func (m *Manager) ValidTasks() (tasks []Task, err error) {
	rows, err := m.db.Query("SELECT id, taskdate, description, done FROM tasks WHERE date('now', 'localtime') < taskdate AND taskdate < date('now', '+7 days', 'localtime') AND done IS NOT TRUE")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		task := Task{}
		rows.Scan(&task.ID, &task.Date, &task.Description, &task.Done)
		tasks = append(tasks, task)
	}

	return
}

func (m *Manager) NotDoneTasks() (tasks []Task, any bool, err error) {
	rows, err := m.db.Query(`SELECT id, taskdate, description, done FROM tasks WHERE taskdate < date('now','-1 days', 'localtime') AND done IS NOT TRUE`)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		task := Task{}
		rows.Scan(&task.ID, &task.Date, &task.Description, &task.Done)
		tasks = append(tasks, task)
	}

	any = len(tasks) != 0

	return
}
