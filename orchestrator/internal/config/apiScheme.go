package config

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

type OrchestratorScheme struct {
	AddTask               Endpoint `json:"addTask"`
	GetTasksStatus        Endpoint `json:"getTasksStatus"`
	GetResult             Endpoint `json:"getResult"`
	GetOperationsTimeouts Endpoint `json:"getOperationsTimeouts"`
	SetOperationsTimeouts Endpoint `json:"setOperationsTimeouts"`
	GetTask               Endpoint `json:"getTask"`
	SetResult             Endpoint `json:"setResult"`
	Register              Endpoint `json:"register"`
}

type AgentScheme struct {
	Ping Endpoint `json:"ping"`
}

type Endpoint struct {
	Url        string   `json:"url"`
	Method     string   `json:"method"`
	RestParams []string `json:"restParams"`
}

type ApiScheme struct {
	Orchestrator OrchestratorScheme `json:"orchestrator"`
	Agent        AgentScheme        `json:"agent"`
}

func getSchemePath() (string, error) {
	now, err := filepath.Abs("")
	if err != nil {
		return "", err
	}
	roots := strings.Split(now, string(filepath.Separator))

	if slices.Contains(roots, "orchestrator") {
		goodPathes := []string{}
		for _, path := range roots {
			if path == "orchestrator" {
				break
			}
			goodPathes = append(goodPathes, path)
		}
		goodPathes = append(goodPathes, "orchestrator", "apiScheme.json")
		return strings.Join(goodPathes, string(filepath.Separator)), nil
	}
	var res string
	filepath.WalkDir(now, func(path string, d fs.DirEntry, err error) error {
		if d.Name() == "apiScheme.json" && res == "" {
			res = path
		}
		return nil
	})
	if res != "" {
		return res, nil
	}
	return res, fmt.Errorf("ApiScheme didn`t Found")
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

func NewApiSchemeProvider() (*ApiSchemeProvider, error) {
	path, err := getSchemePath()
	if err != nil {
		return nil, err
	}
	schemePath := path
	asp := &ApiSchemeProvider{schemePath: schemePath, scheme: &ApiScheme{}}
	asp.Update()
	return asp, nil
}
