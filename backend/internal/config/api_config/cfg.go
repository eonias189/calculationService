package api_config

import (
	"os"

	"github.com/eonias189/calculationService/backend/internal/errors"
)

type Config struct {
	Address             string
	OrchestratorAddress string
	PostgresConn        string
}

func Get() (*Config, error) {
	addr := os.Getenv("ADDRESS")
	if addr == "" {
		return nil, errors.ErrMissingEnvParam("ADDRESS")
	}

	orchAddr := os.Getenv("ORCHESTRATOR_ADDRESS")
	if orchAddr == "" {
		return nil, errors.ErrMissingEnvParam("ORCHESTRATOR_ADDRESS")
	}

	pgConn := os.Getenv("POSTGRES_CONN")
	if pgConn == "" {
		return nil, errors.ErrMissingEnvParam("POSTGRES_CONN")
	}

	return &Config{
		Address:             addr,
		OrchestratorAddress: orchAddr,
		PostgresConn:        pgConn,
	}, nil
}
