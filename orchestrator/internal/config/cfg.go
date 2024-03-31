package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	GRPCAddress      string
	RestApiAddress   string
	PostgresHost     string
	PostgresPort     uint16
	PostgresDB       string
	PostgresUser     string
	PostgresPassowrd string
}

func ErrMissingEnvParam(param string) error {
	return fmt.Errorf("Missing required env param: %v", param)
}

func Get() (*Config, error) {
	grpcAddr := os.Getenv("GRPC_ADDRESS")
	if grpcAddr == "" {
		return nil, ErrMissingEnvParam("GRPC_ADDRESS")
	}

	restAddr := os.Getenv("REST_API_ADDRESS")
	if restAddr == "" {
		return nil, ErrMissingEnvParam("REST_API_ADDRESS")
	}

	pgHost := os.Getenv("POSTGRES_HOST")
	if pgHost == "" {
		return nil, ErrMissingEnvParam("POSTGRES_HOST")
	}

	pgPortStr := os.Getenv("POSTGRES_PORT")
	if pgPortStr == "" {
		return nil, ErrMissingEnvParam("POSTGRES_PORT")
	}
	pgPortInt, err := strconv.Atoi(pgPortStr)
	if err != nil {
		return nil, err
	}

	pgDB := os.Getenv("POSTGRES_DB")
	if pgDB == "" {
		return nil, ErrMissingEnvParam("POSTGRES_DB")
	}

	pgUser := os.Getenv("POSTGRES_USER")
	if pgUser == "" {
		return nil, ErrMissingEnvParam("POSTGRES_USER")
	}

	pgPassword := os.Getenv("POSTGRES_PASSWORD")
	if pgPassword == "" {
		return nil, ErrMissingEnvParam("POSTGRES_PASSWORD")
	}

	return &Config{
		GRPCAddress:      grpcAddr,
		RestApiAddress:   restAddr,
		PostgresHost:     pgHost,
		PostgresPort:     uint16(pgPortInt),
		PostgresDB:       pgDB,
		PostgresUser:     pgUser,
		PostgresPassowrd: pgPassword,
	}, nil

}
