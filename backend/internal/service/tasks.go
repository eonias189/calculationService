package service

import (
	"context"

	errs "github.com/eonias189/calculationService/backend/internal/errors"
	"github.com/eonias189/calculationService/backend/internal/logger"
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

type TaskWithTimeouts struct {
	Task
	Timeouts
}
type TaskService struct {
	pool *pgxpool.Pool
}

func (ts *TaskService) init() error {
	query := `CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		user_id INTEGER references users(id),
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
	query := `SELECT * FROM tasks WHERE user_id=$1 LIMIT $2 OFFSET $3`
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
			logger.Error(err)
			continue
		}
		res = append(res, task)
	}

	return res, nil
}

func (ts *TaskService) GetExecutingForAgent(id int64) ([]TaskWithTimeouts, error) {
	query := `SELECT * FROM tasks JOIN timeouts ON tasks.user_id=timeouts.user_id WHERE executor=$1 and status='executing'`
	conn, err := ts.pool.Acquire(context.TODO())
	if err != nil {
		return []TaskWithTimeouts{}, err
	}

	defer conn.Release()
	rows, err := conn.Query(context.TODO(), query, id)
	if err != nil {
		return []TaskWithTimeouts{}, err
	}

	res := []TaskWithTimeouts{}
	defer rows.Close()
	for rows.Next() {
		var task TaskWithTimeouts
		err := rows.Scan(&task.Task.Id, &task.Task.UserId, &task.Task.Executor, &task.Task.Expression,
			&task.Task.Result, &task.Task.Status, &task.Timeouts.UserId, &task.Timeouts.Add,
			&task.Timeouts.Sub, &task.Timeouts.Mul, &task.Timeouts.Div)
		if err != nil {
			logger.Error(err)
			continue
		}
		res = append(res, task)
	}

	return res, nil
}

func (ts *TaskService) SetUnexecutingForAgent(id int64) error {
	query := `UPDATE tasks SET executor=0, status='pending' WHERE executor=$1 and status='executing'`
	return ts.pool.AcquireFunc(context.TODO(), func(c *pgxpool.Conn) error {
		_, err := c.Exec(context.TODO(), query, id)
		return err
	})
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

	err = row.Scan(&res.Id, &res.UserId, &res.Executor, &res.Expression, &res.Result, &res.Status)
	if err != nil && err.Error() == pgx.ErrNoRows.Error() {
		return Task{}, errs.ErrNotFound
	}

	if err != nil {
		return Task{}, err
	}

	return res, nil
}

func (ts *TaskService) SetPendingForDisactiveAgents() error {
	query := `
	UPDATE tasks SET executor=0, status='pending' WHERE status='executing' AND executor IN (SELECT id FROM agents WHERE active=false);
	UPDATE agents SET running_threads=0 WHERE active=false
	`
	return ts.pool.AcquireFunc(context.TODO(), func(c *pgxpool.Conn) error {
		_, err := c.Exec(context.TODO(), query)
		return err
	})
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

func (ts *TaskService) GetAllPending() ([]TaskWithTimeouts, error) {
	query := `SELECT * FROM tasks JOIN timeouts ON tasks.user_id=timeouts.user_id WHERE status='pending'`
	conn, err := ts.pool.Acquire(context.TODO())
	if err != nil {
		return []TaskWithTimeouts{}, err
	}

	defer conn.Release()
	rows, err := conn.Query(context.TODO(), query)
	if err != nil {
		return []TaskWithTimeouts{}, err
	}

	defer rows.Close()
	res := []TaskWithTimeouts{}
	for rows.Next() {
		var task TaskWithTimeouts
		err := rows.Scan(&task.Task.Id, &task.Task.UserId, &task.Task.Executor, &task.Task.Expression,
			&task.Task.Result, &task.Task.Status, &task.Timeouts.UserId, &task.Timeouts.Add, &task.Timeouts.Sub,
			&task.Timeouts.Mul, &task.Timeouts.Div)
		if err != nil {
			logger.Error(err)
			continue
		}
		res = append(res, task)
	}

	return res, nil
}

func NewTaskService(pool *pgxpool.Pool) (*TaskService, error) {
	ts := &TaskService{pool: pool}
	err := ts.init()

	if err != nil {
		return nil, err
	}

	return ts, nil
}
