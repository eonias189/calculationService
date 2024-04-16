package service

import (
	"context"
	"errors"

	errs "github.com/eonias189/calculationService/backend/internal/errors"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Timeouts struct {
	UserId int64
	Add    uint
	Sub    uint
	Mul    uint
	Div    uint
}

var (
	DefaultTimeouts = Timeouts{}
)

type TimeoutsSerice struct {
	pool *pgxpool.Pool
}

func (ts *TimeoutsSerice) init() error {
	query := `CREATE TABLE IF NOT EXISTS timeouts (
		user_id INTEGER references users(id),
		add INTEGER,
		sub INTEGER,
		mul INTEGER,
		div INTEGER
	)`

	return ts.pool.AcquireFunc(context.TODO(), func(c *pgxpool.Conn) error {
		_, err := c.Exec(context.TODO(), query)
		return err
	})
}

func (ts *TimeoutsSerice) GetForUser(userId int64) (Timeouts, error) {
	query := `SELECT * FROM timeouts WHERE user_id=$1`
	var timeouts Timeouts
	err := ts.pool.AcquireFunc(context.TODO(), func(c *pgxpool.Conn) error {
		return c.QueryRow(context.TODO(), query, userId).Scan(&timeouts.UserId, &timeouts.Add, &timeouts.Sub, &timeouts.Mul, &timeouts.Div)
	})

	if err != nil && err.Error() == pgx.ErrNoRows.Error() {
		return Timeouts{}, errs.ErrNotFound
	}

	if err != nil {
		return Timeouts{}, err
	}

	return timeouts, nil
}

func (ts *TimeoutsSerice) Add(timeouts Timeouts) error {
	query := `INSERT INTO timeouts (user_Id, add, sub, mul, div) VALUES ($1, $2, $3, $4, $5)`
	return ts.pool.AcquireFunc(context.TODO(), func(c *pgxpool.Conn) error {
		_, err := c.Exec(context.TODO(), query, timeouts.UserId, timeouts.Add, timeouts.Sub, timeouts.Mul, timeouts.Div)
		return err
	})
}

func (ts *TimeoutsSerice) Update(timeouts Timeouts) error {
	query := `UPDATE timeouts SET add=$1, sub=$2, mul=$3, div=$4 WHERE user_id=$5`
	return ts.pool.AcquireFunc(context.TODO(), func(c *pgxpool.Conn) error {
		_, err := c.Exec(context.TODO(), query, timeouts.Add, timeouts.Sub, timeouts.Mul, timeouts.Div, timeouts.UserId)
		return err
	})
}

// Add if not exists, update if exists
func (ts *TimeoutsSerice) Put(timeouts Timeouts) error {
	_, wasntErr := ts.GetForUser(timeouts.UserId)
	if wasntErr != nil && errors.Is(wasntErr, errs.ErrNotFound) {
		return ts.Add(timeouts)
	}

	return ts.Update(timeouts)
}

func NewTimeoutsService(pool *pgxpool.Pool) (*TimeoutsSerice, error) {
	ts := &TimeoutsSerice{pool: pool}
	err := ts.init()

	if err != nil {
		return nil, err
	}

	return ts, nil
}
