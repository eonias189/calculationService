package orchestrator_config

import (
	"os"

	"github.com/eonias189/calculationService/backend/internal/errors"
)

type Config struct {
	GRPCAddress  string
	HttpAddress  string
	PostgresConn string
}

func Get() (*Config, error) {
	grpcAddr := os.Getenv("GRPC_ADDRESS")
	if grpcAddr == "" {
		return nil, errors.ErrMissingEnvParam("GRPC_ADDRESS")
	}

	httpAddr := os.Getenv("HTTP_ADDRESS")
	if httpAddr == "" {
		return nil, errors.ErrMissingEnvParam("HTTP_ADDRESS")
	}

	pgConn := os.Getenv("POSTGRES_CONN")
	if pgConn == "" {
		return nil, errors.ErrMissingEnvParam("POSTGRES_CONN")
	}

	return &Config{
		GRPCAddress:  grpcAddr,
		HttpAddress:  httpAddr,
		PostgresConn: pgConn,
	}, nil

}
