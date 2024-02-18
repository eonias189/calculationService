package orchestrator

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"orchestrator/internal/api"
	"orchestrator/internal/db"
	"orchestrator/internal/timeouts"
	"time"

	c "backend/contract"
	"backend/utils"
)

type Orchestrator struct {
	api *api.AgentApi
	db  *db.DB
}

func NewTaskId() string {
	u := make([]byte, 16)
	rand.Read(u)

	u[8] = (u[8] | 0x80) & 0xBF
	u[6] = (u[6] | 0x40) & 0x4F

	return hex.EncodeToString(u)
}

func (o *Orchestrator) AddTask(expression string) error {
	fmt.Println("adding Task", expression)
	newId := NewTaskId()
	_, err := o.db.GetTask(newId)

	for err == nil {
		newId = NewTaskId()
		_, err = o.db.GetTask(newId)
	}
	return o.db.AddTask(newId, expression)
}
func (o *Orchestrator) GetTask(agentId int) (*c.Task, *c.Timeouts, error) {
	task, err := o.db.GetPendingTask()
	if err != nil {
		return &c.Task{}, &c.Timeouts{}, err
	}
	newTask := &c.Task{Id: task.Id, Expression: task.Expression, Result: task.Result, AgentId: int64(agentId), Status: c.TaskStatus_execution}
	err = o.db.UpdateTask(task.Id, newTask)
	if err != nil {
		return &c.Task{}, &c.Timeouts{}, err
	}
	fmt.Println("sending task", task.Expression, "to", agentId)
	tmts, _ := timeouts.GetTimeouts()
	return task, tmts, nil
}
func (o *Orchestrator) GetTasks() ([]*c.Task, error) {
	fmt.Println("sending tasks")
	return o.db.GetTasks()
}
func (o *Orchestrator) GetAgents() ([]*c.AgentData, error) {
	fmt.Println("sending agents data")
	res := []*c.AgentData{}
	agents, err := o.db.GetAgents()
	if err != nil {
		return res, err
	}
	for _, agent := range agents {
		if agent.Id != -1 {
			res = append(res, agent)
		}
	}
	return res, nil
}
func (o *Orchestrator) GetTimeouts() (*c.Timeouts, error) {
	tmts, _ := timeouts.GetTimeouts()
	return tmts, nil
}
func (o *Orchestrator) SetTimeouts(tmts *c.Timeouts) error {
	fmt.Println("setting timeouts", tmts.Add, tmts.Substract, tmts.Multiply, tmts.Divide)
	return timeouts.SetTimeouts(tmts)
}

func (o *Orchestrator) SetResult(id string, res int, status c.TaskStatus) error {
	fmt.Println("setting result", id, res, status)
	task, err := o.db.GetTask(id)
	if err != nil {
		return err
	}
	task.Result = int64(res)
	task.Status = status
	task.AgentId = -1
	return o.db.UpdateTask(id, task)

}
func (o *Orchestrator) Register(url string) (int, error) {
	fmt.Println("registring", url)
	o.db.AddAgent(url)
	agent, err := o.db.GetAgentByUrl(url)
	return int(agent.Id), err
}

type getStatusTask struct {
	resp *c.AgentStatus
	err  error
}

func (o *Orchestrator) updateAgentData(agent *c.AgentData) utils.Task {
	return utils.NewTask(func() {
		newAgent := &c.AgentData{Id: agent.Id, Url: agent.Url, Ping: agent.Ping, Status: agent.Status}
		start := time.Now()
		resChan := make(chan getStatusTask)
		resp := &c.AgentStatus{}
		var err error
		go func() {
			status, err := o.api.GetStatus(agent.Url)
			resChan <- getStatusTask{resp: status, err: err}
		}()
		select {
		case res := <-resChan:
			resp = res.resp
			err = res.err
		case <-time.After(time.Second):
			err = fmt.Errorf("Timeout")
			resp.ExecutingThreads = 0
			resp.MaxThreads = 0
		}
		finish := time.Now()
		if err != nil {
			newAgent.Ping = 999
			newAgent.Status.ExecutingThreads = 0
		} else {
			pingDur := finish.Sub(start)
			ping := min(pingDur.Milliseconds(), 999)
			newAgent.Ping = ping
			newAgent.Status.ExecutingThreads = resp.ExecutingThreads
			newAgent.Status.MaxThreads = resp.MaxThreads
		}
		// fmt.Println("updating", newAgent.Id, "to", newAgent.Ping, newAgent.Status.ExecutingThreads, newAgent.Status.MaxThreads)
		o.db.UpdateAgent(int(agent.Id), newAgent)
	})
}

func (o *Orchestrator) updateAgentsData() {
	for {
		agents, err := o.db.GetAgents()
		if err != nil {
			fmt.Println(err)
			time.Sleep(time.Second * 5)
			continue
		}
		wp := utils.NewWorkerPool(len(agents))
		wp.Start()
		for _, agent := range agents {
			if agent.Id == -1 {
				continue
			}
			wp.AddTask(o.updateAgentData(agent))
		}
		wp.Close()
		time.Sleep(time.Second * 5)
	}
}

func (o *Orchestrator) searchDeadAgents() {
	for {
		agents, err := o.db.GetAgents()
		if err != nil {
			time.Sleep(time.Second * 10)
		}
		for _, agent := range agents {
			if agent.Id == -1 {
				continue
			}
			if agent.Ping == int64(999) {
				// fmt.Println("found dead agent", agent.Id)
				tasks, err := o.db.GetTasksOfAgent(agent.Id)
				if err != nil {
					continue
				}
				for _, task := range tasks {
					task.Status = c.TaskStatus_pending
					task.AgentId = -1
					o.db.UpdateTask(task.Id, task)
				}
			}
		}
		time.Sleep(time.Second * 10)
	}
}

func (o *Orchestrator) Run(url string) {
	fmt.Println("starting at", url)
	go o.updateAgentsData()
	go o.searchDeadAgents()
}

func NewOrchestrator(api *api.AgentApi, db *db.DB) *Orchestrator {
	return &Orchestrator{api: api, db: db}
}
