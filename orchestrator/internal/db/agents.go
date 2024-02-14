package db

import (
	c "backend/contract"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type AgentsDB struct {
	db *sql.DB
}

func (adb *AgentsDB) AddAgent(a c.AgentData) error {
	query := fmt.Sprintf(`INSERT INTO agents (ping, maxThreads, execThreads, url) VALUES (%v, %v, %v, "%v")`, a.Ping, a.Status.MaxThreads, a.Status.ExecutingThreads, a.Url)
	_, err := adb.db.Exec(query)
	return err
}

func (adb *AgentsDB) GetAgents() ([]c.AgentData, error) {
	agents := []c.AgentData{}
	query := `SELECT * FROM agents`
	rows, err := adb.db.Query(query)
	if err != nil {
		return agents, err
	}
	defer rows.Close()
	for rows.Next() {
		agent := c.AgentData{Status: &c.AgentStatus{}}
		err = rows.Scan(&agent.Id, &agent.Ping, &agent.Status.MaxThreads, &agent.Status.ExecutingThreads, &agent.Url)
		if err != nil {
			continue
		}
		agents = append(agents, agent)
	}
	return agents, nil
}

func (adb *AgentsDB) GetAgent(id int) (c.AgentData, error) {
	agent := c.AgentData{Status: &c.AgentStatus{}}
	query := fmt.Sprintf(`SELECT * FROM agents WHERE id=%v`, id)
	row, err := adb.db.Query(query)
	if err != nil {
		return agent, err
	}
	defer row.Close()
	if !row.Next() {
		return agent, fmt.Errorf("Agent Not Found")
	}
	err = row.Scan(&agent.Id, &agent.Ping, &agent.Status.MaxThreads, &agent.Status.ExecutingThreads, &agent.Url)
	return agent, err
}

func (adb *AgentsDB) DeleteAgent(id int) error {
	query := fmt.Sprintf(`DELETE from agents WHERE id=%v`, id)
	_, err := adb.db.Exec(query)
	return err
}

func (adb *AgentsDB) UpdateAgent(id int, ping int, status c.AgentStatus) error {
	query := fmt.Sprintf(`UPDATE agents
	SET ping=%v, maxThreads=%v, execThreads=%v
	WHERE id=%v`, ping, status.MaxThreads, status.ExecutingThreads, id)
	_, err := adb.db.Exec(query)
	return err
}

func NewAgentsDB(path string) (*AgentsDB, error) {
	scheme := `CREATE TABLE IF NOT EXISTS agents (
		id INTEGER NOT NULL PRIMARY KEY,
		ping int,
		maxThreads int,
		execThreads int,
		url text UNIQUE
	);`

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(scheme)
	if err != nil {
		return nil, err
	}
	return &AgentsDB{db: db}, nil
}
