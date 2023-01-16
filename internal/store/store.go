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
	"database/sql"
	"time"
)

type Manager struct {
	db    *sql.DB

    // Deprecated: I was going to do something 
    // with this but I forgot about it
	// today time.Time
}

func NewManager(db *sql.DB) (m *Manager, err error) {
	m = &Manager{}
	m.db = db
	// m.today = time.Now()
	err = m.initDB()
	return
}

func (m *Manager) initDB() (err error) {
	creation, err := m.db.Prepare("CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY, taskdate DATE, description TEXT, done BOOL DEFAULT FALSE NOT NULL)")
	if err != nil {
		return
	}
	defer creation.Close()
	creation.Exec()

	reminders, err := m.db.Prepare("CREATE TABLE IF NOT EXISTS reminders (id INTEGER PRIMARY KEY, hour DATE, description TEXT)")
	if err != nil {
		return
	}
	defer reminders.Close()
	reminders.Exec()

	return
}

func (m *Manager) AddTask(taskdate time.Time, description string) (err error) {
	insertion, err := m.db.Prepare("INSERT INTO tasks (taskdate, description) VALUES (?, ?)")
	if err != nil {
		return
	}
	defer insertion.Close()

	_, err = insertion.Exec(taskdate, description)
	if err != nil {
		return
	}

	return
}

func (m *Manager) RemoveTask(id string) (err error) {
	statement, err := m.db.Prepare("DELETE FROM tasks where id = ?")
	if err != nil {
        return
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		return
	}

	return
}

func (m *Manager) AddReminder(t time.Time, description string) (err error) {
	insertion, err := m.db.Prepare("INSERT INTO reminders (hour, description) VALUES (?, ?)")
	if err != nil {
		return
	}
	defer insertion.Close()

    _, err = insertion.Exec(t, description)
	if err != nil {
		return
	}

	return
}

func (m *Manager) RemoveReminder(id string) (err error) {
	statement, err := m.db.Prepare("DELETE FROM reminders WHERE id = ?")
	if err != nil {
		return
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		return
	}

	return
}

func (m *Manager) Close() error {
	return m.db.Close()
}
