package interfaces

import (
	"time"
)

type OperationsTimeouts struct {
	Add      time.Duration `json:"add"`
	Subtract time.Duration `json:"substract"`
	Multiply time.Duration `json:"multiply"`
	Divide   time.Duration `json:"divide"`
}
