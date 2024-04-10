package service

import (
	"context"
	"errors"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Agent struct {
	Id             int64
	Ping           int64
	MaxThreads     int
	RunningThreads int
}

type AgentService struct {
	pool *pgxpool.Pool
}

func (as *AgentService) init() error {
	query := `CREATE TABLE IF NOT EXISTS agents (
		id SERIAL PRIMARY KEY,
		ping INTEGER,
		max_threads INTEGER,
		running_threads INTEGER
	)`

	return as.pool.AcquireFunc(context.TODO(), func(c *pgxpool.Conn) error {
		_, err := c.Exec(context.TODO(), query)
		return err
	})
}

func (as *AgentService) Add(agent Agent) (int64, error) {
	var id int64
	query := `INSERT INTO agents (ping, max_threads, running_threads) values ($1, $2, $3) RETURNING id`
	err := as.pool.AcquireFunc(context.TODO(), func(c *pgxpool.Conn) error {
		return c.QueryRow(context.TODO(), query, agent.Ping, agent.MaxThreads, agent.RunningThreads).Scan(&id)
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (as *AgentService) GetAll(limit, offset int) ([]Agent, error) {
	res := []Agent{}
	query := `SELECT * FROM agents LIMIT $1 OFFSET $2`
	conn, err := as.pool.Acquire(context.TODO())

	if err != nil {
		return []Agent{}, err
	}

	defer conn.Release()
	rows, err := conn.Query(context.TODO(), query, limit, offset)

	if err != nil {
		return res, err
	}

	defer rows.Close()
	for rows.Next() {
		agent := Agent{}
		err = rows.Scan(&agent.Id, &agent.Ping, &agent.MaxThreads, &agent.RunningThreads)
		if err != nil {
			continue
		}
		res = append(res, agent)
	}

	return res, nil
}

func (as *AgentService) GetById(id int64) (Agent, error) {
	var res Agent
	query := `SELECT * FROM agents WHERE id=$1 LIMIT 1`
	conn, err := as.pool.Acquire(context.TODO())

	if err != nil {
		return Agent{}, err
	}

	defer conn.Release()
	row := conn.QueryRow(context.TODO(), query, id)

	err = row.Scan(&res.Id, &res.Ping, &res.MaxThreads, &res.RunningThreads)
	if errors.Is(pgx.ErrNoRows, err) {
		return res, ErrNotFound
	}

	if err != nil {
		return Agent{}, err
	}

	return res, nil
}

func (as *AgentService) Update(agent Agent) error {
	query := `UPDATE agents SET ping=$1, max_threads=$2, running_threads=$3 WHERE id=$5`
	return as.pool.AcquireFunc(context.TODO(), func(c *pgxpool.Conn) error {
		_, err := c.Exec(context.TODO(), query, agent.Ping, agent.MaxThreads, agent.RunningThreads, agent.Id)
		return err
	})
}

func (as *AgentService) Delete(id int64) error {
	query := `DELETE FROM agents WHERE id=$1`
	return as.pool.AcquireFunc(context.TODO(), func(c *pgxpool.Conn) error {
		_, err := c.Exec(context.TODO(), query, id)
		return err
	})
}

func NewAgentService(pool *pgxpool.Pool) (*AgentService, error) {
	as := &AgentService{pool: pool}
	err := as.init()

	if err != nil {
		return nil, err
	}

	return as, err
}
