package config

import (
	"fmt"
	"os"
)

type Config struct {
	GRPCAddress  string
	HttpAddress  string
	PostgresConn string
}

func ErrMissingEnvParam(param string) error {
	return fmt.Errorf("Missing required env param: %v", param)
}

func Get() (*Config, error) {
	grpcAddr := os.Getenv("GRPC_ADDRESS")
	if grpcAddr == "" {
		return nil, ErrMissingEnvParam("GRPC_ADDRESS")
	}

	httpAddr := os.Getenv("HTTP_ADDRESS")
	if httpAddr == "" {
		return nil, ErrMissingEnvParam("HTTP_ADDRESS")
	}

	pgConn := os.Getenv("POSTGRES_CONN")
	if pgConn == "" {
		return nil, ErrMissingEnvParam("POSTGRES_CONN")
	}

	return &Config{
		GRPCAddress:  grpcAddr,
		HttpAddress:  httpAddr,
		PostgresConn: pgConn,
	}, nil

}
