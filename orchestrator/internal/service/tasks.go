package service

import (
	"context"
	"errors"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	TaskStatusPending   TaskStatus = "pending"
	TaskStatusExecuting TaskStatus = "executing"
	TaskStatusSuccess   TaskStatus = "success"
	TaskStatusError     TaskStatus = "error"
)

type TaskStatus string

type Task struct {
	Id         int64
	UserId     int64
	Executor   int64
	Expression string
	Result     float64
	Status     TaskStatus
}

type TaskService struct {
	pool *pgxpool.Pool
}

func (ts *TaskService) init() error {
	query := `CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		user_id INTEGER,
		executor INTEGER,
		expression TEXT,
		result double precision,
		status TEXT
	)`
	return ts.pool.AcquireFunc(context.TODO(), func(c *pgxpool.Conn) error {
		_, err := c.Exec(context.TODO(), query)
		return err
	})
}

func (ts *TaskService) Add(task Task) (int64, error) {
	var id int64
	query := `INSERT INTO tasks (user_id, executor, expression, result, status) values ($1, $2, $3, $4, $5) RETURNING id`
	err := ts.pool.AcquireFunc(context.TODO(), func(c *pgxpool.Conn) error {
		return c.QueryRow(context.TODO(), query, task.UserId, task.Executor, task.Expression, task.Result, task.Status).Scan(&id)
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (ts *TaskService) GetAllForUser(userId int64, limit, offset int) ([]Task, error) {
	res := []Task{}
	query := `SELECT * FROM tasks WHERE user_id=$1 LIMIT $1 OFFSET $2`
	conn, err := ts.pool.Acquire(context.TODO())

	if err != nil {
		return []Task{}, err
	}

	defer conn.Release()
	rows, err := conn.Query(context.TODO(), query, userId, limit, offset)

	if err != nil {
		return res, err
	}

	defer rows.Close()
	for rows.Next() {
		task := Task{}
		err = rows.Scan(&task.Id, &task.UserId, &task.Executor, &task.Expression, &task.Result, &task.Status)
		if err != nil {
			continue
		}
		res = append(res, task)
	}

	return res, nil
}

func (ts *TaskService) GetById(id int64) (Task, error) {
	var res Task
	query := `SELECT * FROM tasks WHERE id=$1 LIMIT 1`
	conn, err := ts.pool.Acquire(context.TODO())

	if err != nil {
		return Task{}, err
	}

	defer conn.Release()
	row := conn.QueryRow(context.TODO(), query, id)

	err = row.Scan(&res.Id, &res.Executor, &res.Expression, &res.Result, &res.Status)
	if errors.Is(pgx.ErrNoRows, err) {
		return res, ErrNotFound
	}

	if err != nil {
		return Task{}, err
	}

	return res, nil
}

func (ts *TaskService) Update(task Task) error {
	query := `UPDATE tasks SET executor=$1, expression=$2::text, result=$3, status=$4::text WHERE id=$5`
	return ts.pool.AcquireFunc(context.TODO(), func(c *pgxpool.Conn) error {
		_, err := c.Exec(context.TODO(), query, task.Executor, task.Expression, task.Result, task.Status, task.Id)
		return err
	})
}

func (ts *TaskService) Delete(id int64) error {
	query := `DELETE FROM tasks WHERE id=$1`
	return ts.pool.AcquireFunc(context.TODO(), func(c *pgxpool.Conn) error {
		_, err := c.Exec(context.TODO(), query, id)
		return err
	})
}

func NewTaskService(pool *pgxpool.Pool) (*TaskService, error) {
	ts := &TaskService{pool: pool}
	err := ts.init()

	if err != nil {
		return nil, err
	}

	return ts, nil
}
