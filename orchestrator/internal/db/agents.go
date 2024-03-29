package db

import (
	c "backend/contract"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func (db *DB) AddAgent(url string) error {
	query := fmt.Sprintf(`INSERT INTO agents (ping, maxThreads, execThreads, url) VALUES (%v, %v, %v, "%v")`, -1, -1, -1, url)
	_, err := db.db.Exec(query)
	return err
}

func (db *DB) GetAgents() ([]*c.AgentData, error) {
	agents := []*c.AgentData{}
	query := `SELECT * FROM agents`
	rows, err := db.db.Query(query)
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
		agents = append(agents, &agent)
	}
	return agents, nil
}

func (db *DB) GetAgentByUrl(url string) (*c.AgentData, error) {
	agent := &c.AgentData{Status: &c.AgentStatus{}}
	query := fmt.Sprintf(`SELECT * FROM agents WHERE url="%v"`, url)
	row, err := db.db.Query(query)
	if err != nil {
		return agent, err
	}
	defer row.Close()
	if !row.Next() {
		return agent, fmt.Errorf("AgentNotFound")
	}
	err = row.Scan(&agent.Id, &agent.Ping, &agent.Status.MaxThreads, &agent.Status.ExecutingThreads, &agent.Url)
	return agent, err
}

func (db *DB) GetAgent(id int) (*c.AgentData, error) {
	agent := &c.AgentData{Status: &c.AgentStatus{}}
	query := fmt.Sprintf(`SELECT * FROM agents WHERE id=%v`, id)
	row, err := db.db.Query(query)
	if err != nil {
		return agent, err
	}
	defer row.Close()
	if !row.Next() {
		return agent, fmt.Errorf("AgentNotFound")
	}
	err = row.Scan(&agent.Id, &agent.Ping, &agent.Status.MaxThreads, &agent.Status.ExecutingThreads, &agent.Url)
	return agent, err
}

func (db *DB) DeleteAgent(id int) error {
	query := fmt.Sprintf(`DELETE from agents WHERE id=%v`, id)
	_, err := db.db.Exec(query)
	return err
}

func (db *DB) UpdateAgent(id int, newAgent *c.AgentData) error {
	query := fmt.Sprintf(`UPDATE agents
	SET ping=%v, maxThreads=%v, execThreads=%v
	WHERE id=%v`, newAgent.Ping, newAgent.Status.MaxThreads, newAgent.Status.ExecutingThreads, id)
	_, err := db.db.Exec(query)
	return err
}
