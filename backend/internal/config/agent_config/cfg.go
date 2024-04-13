package agent_config

import (
	"os"
	"strconv"

	"github.com/eonias189/calculationService/backend/internal/errors"
)

type Config struct {
	OrchestratorAddr string
	MaxThreads       int
}

func getInt(key string) (int, error) {
	keyStr := os.Getenv(key)
	if keyStr == "" {
		return 0, errors.ErrMissingEnvParam(key)
	}

	keyInt, err := strconv.Atoi(keyStr)
	if err != nil {
		return 0, err
	}

	return keyInt, nil
}

func Get() (*Config, error) {
	orchArrd := os.Getenv("ORCHESTRATOR_ADDRESS")
	if orchArrd == "" {
		return nil, errors.ErrMissingEnvParam("ORCHESTRATOR_ADDRESS")
	}

	maxThreads, err := getInt("MAX_THREADS")
	if err != nil {
		return nil, err
	}

	return &Config{
		OrchestratorAddr: orchArrd,
		MaxThreads:       maxThreads,
	}, err
}
