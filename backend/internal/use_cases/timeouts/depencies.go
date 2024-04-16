package use_timeouts

import "github.com/eonias189/calculationService/backend/internal/service"

type TimeoutsSerice interface {
	GetForUser(userId int64) (service.Timeouts, error)
	Put(timeouts service.Timeouts) error
}
