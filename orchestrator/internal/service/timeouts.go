package service

import (
	"context"
	"errors"

	"github.com/jackc/pgx"
)

type Timeouts struct {
	Add int64 `json:"add"`
	Sub int64 `json:"sub"`
	Mul int64 `json:"mul"`
	Div int64 `json:"div"`
}

var (
	DefaultTimeouts = Timeouts{}
)

type TimeoutSerice struct {
	conn *pgx.Conn
}

func (ts *TimeoutSerice) init() error {
	query := `CREATE TABLE IF NOT EXISTS timeouts (
		id SERIAL PRIMARY KEY,
		userId INTEGER,
		add INTEGER,
		sub INTEGER,
		mul INTEGER,
		div INTEGER
	)`

	_, err := ts.conn.Exec(query)
	return err
}

func (ts *TimeoutSerice) Load(ctx context.Context) (Timeouts, error) {
	query := `SELECT FROM timeouts WHERE id=$1`

	var (
		id     int64
		userId int64
		res    Timeouts
	)

	row := ts.conn.QueryRowEx(ctx, query, nil, 1)
	err := row.Scan(&id, &userId, &res.Add, &res.Sub, &res.Mul, &res.Div)

	if errors.Is(err, pgx.ErrNoRows) {
		return DefaultTimeouts, nil
	}

	return res, err
}

func (ts *TimeoutSerice) Save(ctx context.Context, timeouts Timeouts) error {
	return nil
}

func NewTimeoutsService(conn *pgx.Conn) (*TimeoutSerice, error) {
	ts := &TimeoutSerice{conn: conn}
	err := ts.init()

	if err != nil {
		return nil, err
	}

	return ts, nil
}
