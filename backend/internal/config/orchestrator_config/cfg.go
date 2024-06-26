package orchestrator_config

import (
	"os"

	"github.com/eonias189/calculationService/backend/internal/errors"
)

type Config struct {
	Address      string
	PostgresConn string
}

func Get() (*Config, error) {
	addr := os.Getenv("ADDRESS")
	if addr == "" {
		return nil, errors.ErrMissingEnvParam("ADDRESS")
	}

	pgConn := os.Getenv("POSTGRES_CONN")
	if pgConn == "" {
		return nil, errors.ErrMissingEnvParam("POSTGRES_CONN")
	}

	return &Config{
		Address:      addr,
		PostgresConn: pgConn,
	}, nil

}
