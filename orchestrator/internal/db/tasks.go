package db

import (
	c "backend/contract"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func (db *DB) AddTask(t c.Task) error {
	query := fmt.Sprintf(`INSERT INTO tasks VALUES ("%v", "%v", %v, "%v")`, t.Id, t.Expression, t.Result, t.Status)
	_, err := db.db.Exec(query)
	return err
}

func (db *DB) GetTasks() ([]c.Task, error) {
	tasks := []c.Task{}
	query := `SELECT * FROM tasks`
	rows, err := db.db.Query(query)
	if err != nil {
		return tasks, err
	}
	defer rows.Close()
	for rows.Next() {
		task := c.Task{}
		var status string
		err = rows.Scan(&task.Id, &task.Expression, &task.Result, &status)
		if err != nil {
			continue
		}
		task.Status = c.TaskStatus(c.TaskStatus_value[status])
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (db *DB) GetTask(id string) (c.Task, error) {
	var task c.Task
	query := fmt.Sprintf(`SELECT * FROM tasks WHERE id="%v"`, id)
	row, err := db.db.Query(query)
	if err != nil {
		return task, err
	}
	defer row.Close()
	if !row.Next() {
		return task, fmt.Errorf("Task Not Found")
	}
	var status string
	err = row.Scan(&task.Id, &task.Expression, &task.Result, &status)
	if err != nil {
		return task, err
	}
	task.Status = c.TaskStatus(c.TaskStatus_value[status])
	return task, nil
}

func (db *DB) DeleteTask(id string) error {
	query := fmt.Sprintf(`DELETE FROM tasks WHERE id="%v"`, id)
	_, err := db.db.Exec(query)
	return err
}

func (db *DB) UpdateTask(id string, result int, status c.TaskStatus) error {
	query := fmt.Sprintf(`UPDATE tasks
	SET result=%v, status="%v"
	WHERE id="%v"`, result, status, id)
	_, err := db.db.Exec(query)
	return err
}
