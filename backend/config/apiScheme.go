package config

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

type OrchestratorScheme struct {
	AddTask               Endpoint
	GetTasksStatus        Endpoint
	GetResult             Endpoint
	GetOperationsTimeouts Endpoint
	GetTask               Endpoint
	SetResult             Endpoint
	Register              Endpoint
}

type Endpoint struct {
	Url    string `json:"url"`
	Method string `json:"method"`
}

type AgentScheme struct {
	GetTaskStatus Endpoint
	IsWorking     Endpoint
}

type ApiScheme struct {
	Orchestrator OrchestratorScheme `json:"orchestrator"`
	Agent        AgentScheme        `json:"agent"`
}

func getRootPath() string {
	backendDir, _ := filepath.Abs("")
	rootDir := filepath.Dir(backendDir)
	return rootDir
}

type ApiSchemeProvider struct {
	schemePath string
	scheme     *ApiScheme
}

func (asp *ApiSchemeProvider) Update() error {
	file, err := os.Open(asp.schemePath)
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	json.Unmarshal(data, asp.scheme)
	return nil
}

func (asp *ApiSchemeProvider) GetOrchestratorScheme() OrchestratorScheme {
	return asp.scheme.Orchestrator
}

func (asp *ApiSchemeProvider) GetAgentScheme() AgentScheme {
	return asp.scheme.Agent
}

func NewApiSchemeProvider() *ApiSchemeProvider {
	schemePath := filepath.Join(getRootPath(), "apiScheme.json")
	asp := &ApiSchemeProvider{schemePath: schemePath, scheme: &ApiScheme{}}
	asp.Update()
	return asp
}
