package manager

import (
	"database/sql"
	"time"
	// "github.com/golang-module/carbon"
)

type Manager struct {
	db    *sql.DB
	today time.Time
}

func NewManager(db *sql.DB) (m *Manager) {
	m = &Manager{}
	m.db = db
	m.today = time.Now()
	m.CheckTable()
	return
}

func (m *Manager) CheckTable() {
	_, err := m.db.Query("SELECT * FROM tasks")

	if err != nil {
		m.initDB()
		return
	}
}

func (m *Manager) initDB() {
	creation, err := m.db.Prepare("CREATE TABLE tasks (id INTEGER PRIMARY KEY, taskdate DATE, description TEXT)")
	if err != nil {
		panic(err)
	}
	defer creation.Close()
	creation.Exec()
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

/*
Task struct type definition

A struct to be used internally by the program, there's no reason anyone should be reading this lol
*/
type Task struct {
	ID          int
	Date        time.Time
	Description string
}

func (m *Manager) RemoveTask(id int) (err error) {
	statement, err := m.db.Prepare("DELETE FROM tasks where id = ?")
	if err != nil {
		panic(err)
	}

	_, err = statement.Exec(id)
	if err != nil {
		return
	}

	return
}

/*
GetTasks method retrieves all tasks from the database.

If there are no tasks it returns an empty list.
*/
func (m *Manager) GetTasks() (tasks []Task, err error) {
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

func (m *Manager) ValidByDate() (tasks []Task, err error) {
	rows, err := m.db.Query("SELECT id, taskdate, description FROM tasks WHERE date('now') < taskdate")
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

func (m *Manager) Close() error {
	return m.db.Close()
}
