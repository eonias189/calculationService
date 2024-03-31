package service

import (
	"context"
	"errors"

	"github.com/jackc/pgx"
)

var (
	TaskStatusPending   TaskStatus = "pending"
	TaskStatusExecuting TaskStatus = "executing"
	TaskStatusSuccess   TaskStatus = "success"
	TaskStatusError     TaskStatus = "error"
)

type TaskStatus string

type Task struct {
	Id         int64      `json:"id"`
	Executor   int64      `json:"-"`
	Expression string     `json:"expression"`
	Result     float64    `json:"result"`
	Status     TaskStatus `json:"status"`
}

type TaskService struct {
	conn *pgx.Conn
}

func (ts *TaskService) init() error {
	query := `CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		executor INTEGER,
		expression TEXT,
		result double precision,
		status TEXT
	)`
	_, err := ts.conn.Exec(query)
	return err
}

func (ts *TaskService) Add(ctx context.Context, task Task) (int64, error) {
	var res int64
	query := `INSERT INTO tasks (executor, expression, result, status) values ($1, $2, $3, $4) RETURNING id`
	err := ts.conn.QueryRowEx(ctx, query, nil, task.Executor, task.Expression, task.Result, task.Status).Scan(&res)
	return res, err
}

func (ts *TaskService) GetAll(ctx context.Context, limit, offset int) ([]Task, error) {
	res := []Task{}
	query := `SELECT * FROM tasks LIMIT $1 OFFSET $2`
	rows, err := ts.conn.QueryEx(ctx, query, nil, limit, offset)

	if err != nil {
		return res, err
	}

	defer rows.Close()
	for rows.Next() {
		task := Task{}
		err = rows.Scan(&task.Id, &task.Executor, &task.Expression, &task.Result, &task.Status)
		if err != nil {
			continue
		}
		res = append(res, task)
	}

	return res, nil
}

func (ts *TaskService) GetById(ctx context.Context, id int64) (Task, error) {
	var res Task
	query := `SELECT * FROM tasks WHERE id=$1 LIMIT 1`
	row := ts.conn.QueryRowEx(ctx, query, nil, id)

	if row == nil {
		return res, ErrNotFound
	}

	err := row.Scan(&res.Id, &res.Executor, &res.Expression, &res.Result, &res.Status)
	if errors.Is(pgx.ErrNoRows, err) {
		return res, ErrNotFound
	}

	return res, err
}

func (ts *TaskService) Update(ctx context.Context, task Task) error {
	query := `UPDATE tasks SET executor=$1, expression=$2::text, result=$3, status=$4::text WHERE id=$5`
	_, err := ts.conn.ExecEx(ctx, query, nil, task.Executor, task.Expression, task.Result, task.Status, task.Id)
	return err
}

func (ts *TaskService) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM tasks WHERE id=$1`
	_, err := ts.conn.ExecEx(ctx, query, nil, id)
	return err
}

func NewTaskService(conn *pgx.Conn) (*TaskService, error) {
	ts := &TaskService{conn: conn}
	err := ts.init()
	return ts, err
}
