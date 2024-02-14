package db

import (
	c "backend/contract"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type TasksDB struct {
	db *sql.DB
}

func (tdb *TasksDB) AddTask(t c.Task) error {
	query := fmt.Sprintf(`INSERT INTO tasks VALUES ("%v", "%v", %v, "%v")`, t.Id, t.Expression, t.Result, t.Status)
	_, err := tdb.db.Exec(query)
	return err
}

func (tdb *TasksDB) GetTasks() ([]c.Task, error) {
	tasks := []c.Task{}
	query := `SELECT * FROM tasks`
	rows, err := tdb.db.Query(query)
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

func (tdb *TasksDB) GetTask(id string) (c.Task, error) {
	var task c.Task
	query := fmt.Sprintf(`SELECT * FROM tasks WHERE id="%v"`, id)
	row, err := tdb.db.Query(query)
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

func (tdb *TasksDB) DeleteTask(id string) error {
	query := fmt.Sprintf(`DELETE FROM tasks WHERE id="%v"`, id)
	_, err := tdb.db.Exec(query)
	return err
}

func (tdb *TasksDB) UpdateTask(id string, result int, status c.TaskStatus) error {
	query := fmt.Sprintf(`UPDATE tasks
	SET result=%v, status="%v"
	WHERE id="%v"`, result, status, id)
	_, err := tdb.db.Exec(query)
	return err
}

func NewTasksDB(path string) (*TasksDB, error) {
	scheme := `CREATE TABLE IF NOT EXISTS tasks (
		id text NOT NULL PRIMARY KEY,
		expression text,
		result int,
		status text
	);
	`
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(scheme)
	if err != nil {
		return nil, err
	}
	return &TasksDB{db: db}, nil
}
