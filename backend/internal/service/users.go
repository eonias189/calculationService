package service

import (
	"context"
	"errors"

	errs "github.com/eonias189/calculationService/backend/internal/errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	Id             int64
	Login          string
	HashedPassword string
}

type UserService struct {
	pool *pgxpool.Pool
}

func (us *UserService) init() error {
	query := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		login VARCHAR(255) UNIQUE,
		hashed_password TEXT
	)`

	return us.pool.AcquireFunc(context.TODO(), func(c *pgxpool.Conn) error {
		_, err := c.Exec(context.TODO(), query)
		return err
	})
}

func (us *UserService) Add(user User) (int64, error) {
	query := `INSERT INTO users (login, hashed_password) VALUES ($1, $2) RETURNING id`
	var id int64
	err := us.pool.AcquireFunc(context.TODO(), func(c *pgxpool.Conn) error {
		return c.QueryRow(context.TODO(), query, user.Login, user.HashedPassword).Scan(&id)
	})
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (us *UserService) GetById(id int64) (User, error) {
	query := `SELECT * FROM users WHERE id=$1`
	var user User
	err := us.pool.AcquireFunc(context.TODO(), func(c *pgxpool.Conn) error {
		return c.QueryRow(context.TODO(), query, id).Scan(&user.Id, &user.Login, &user.HashedPassword)
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return User{}, errs.ErrNotFound
	}

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (us *UserService) GetByLogin(login string) (User, error) {
	query := `SELECT * FROM users WHERE login=$1`
	var user User
	err := us.pool.AcquireFunc(context.TODO(), func(c *pgxpool.Conn) error {
		return c.QueryRow(context.TODO(), query, login).Scan(&user.Id, &user.Login, &user.HashedPassword)
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return User{}, errs.ErrNotFound
	}

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func NewUserService(pool *pgxpool.Pool) (*UserService, error) {
	us := &UserService{pool: pool}
	err := us.init()
	if err != nil {
		return nil, err
	}

	return us, nil
}
