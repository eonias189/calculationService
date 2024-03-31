package service

import (
	"context"
	"errors"

	"github.com/jackc/pgx"
)

type Agent struct {
	Id             int64 `json:"id"`
	Ping           int64 `json:"ping"`
	MaxThreads     int   `json:"maxThreads"`
	RunningThreads int   `json:"runningThreads"`
}

type AgentService struct {
	conn *pgx.Conn
}

func (as *AgentService) init() error {
	query := `CREATE TABLE IF NOT EXISTS agents (
		id SERIAL PRIMARY KEY,
		ping INTEGER,
		maxThreads INTEGER,
		runningThreads INTEGER
	)`

	_, err := as.conn.Exec(query)
	return err
}

func (as *AgentService) Add(ctx context.Context, agent Agent) (int64, error) {
	var res int64
	query := `INSERT INTO agents (ping, maxThreads, runningThreads) values ($1, $2, $3) RETURNING id`
	err := as.conn.QueryRowEx(ctx, query, nil, agent.Ping, agent.MaxThreads, agent.RunningThreads).Scan(&res)
	return res, err
}

func (as *AgentService) GetAll(ctx context.Context, limit, offset int) ([]Agent, error) {
	res := []Agent{}
	query := `SELECT * FROM agents LIMIT $1 OFFSET $2`
	rows, err := as.conn.QueryEx(ctx, query, nil, limit, offset)

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

func (as *AgentService) GetById(ctx context.Context, id int64) (Agent, error) {
	var res Agent
	query := `SELECT * FROM agents WHERE id=$1 LIMIT 1`
	row := as.conn.QueryRowEx(ctx, query, nil, id)

	if row == nil {
		return res, ErrNotFound
	}

	err := row.Scan(&res.Id, &res.Ping, &res.MaxThreads, &res.RunningThreads)
	if errors.Is(pgx.ErrNoRows, err) {
		return res, ErrNotFound
	}

	return res, err
}

func (as *AgentService) Update(ctx context.Context, agent Agent) error {
	query := `UPDATE agents SET ping=$1, maxThreads=$2, runningThreads=$3 WHERE id=$5`
	_, err := as.conn.ExecEx(ctx, query, nil, agent.Ping, agent.MaxThreads, agent.RunningThreads, agent.Id)
	return err
}

func (as *AgentService) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM agents WHERE id=$1`
	_, err := as.conn.ExecEx(ctx, query, nil, id)
	return err
}

func NewAgentService(conn *pgx.Conn) (*AgentService, error) {
	as := &AgentService{conn: conn}
	err := as.init()
	return as, err
}
