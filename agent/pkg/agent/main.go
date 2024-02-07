package agent

import (
	"fmt"
)

type Agent struct {
}

func (o *Agent) Ping() (bool, error) {
	fmt.Println("this is Agent")
	return true, nil
}

func (o *Agent) Run(url string) {
	fmt.Println("starting at", url)
}

func New() *Agent {
	return &Agent{}
}
