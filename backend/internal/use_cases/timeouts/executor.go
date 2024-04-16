package use_timeouts

import (
	"errors"

	errs "github.com/eonias189/calculationService/backend/internal/errors"
	"github.com/eonias189/calculationService/backend/internal/service"
)

type Executor struct {
	timeoutsService TimeoutsSerice
}

func (e *Executor) GetTimeouts(userId int64) (GetTimeoutsResp, error) {
	timeouts, err := e.timeoutsService.GetForUser(userId)
	if errors.Is(err, errs.ErrNotFound) {
		err = nil
		timeouts = service.DefaultTimeouts
	}

	if err != nil {
		return GetTimeoutsResp{}, err
	}

	return GetTimeoutsResp{Timeouts: TimeoutsSource{
		Add: &timeouts.Add,
		Sub: &timeouts.Sub,
		Mul: &timeouts.Mul,
		Div: &timeouts.Div,
	}}, nil
}

func (e *Executor) PatchTimeouts(body PatchTimeoutsBody, userId int64) (PatchTimeoutsResp, error) {
	timeouts, err := e.timeoutsService.GetForUser(userId)
	if errors.Is(err, errs.ErrNotFound) {
		timeouts = service.DefaultTimeouts
		timeouts.UserId = userId
		err = nil
	}

	if err != nil {
		return PatchTimeoutsResp{}, err
	}

	if body.Add != nil {
		timeouts.Add = *body.Add
	}
	if body.Sub != nil {
		timeouts.Sub = *body.Sub
	}
	if body.Mul != nil {
		timeouts.Mul = *body.Mul
	}
	if body.Div != nil {
		timeouts.Div = *body.Div
	}

	err = e.timeoutsService.Put(timeouts)
	if err != nil {
		return PatchTimeoutsResp{}, err
	}

	return PatchTimeoutsResp{Timeouts: TimeoutsSource{
		Add: &timeouts.Add,
		Sub: &timeouts.Sub,
		Mul: &timeouts.Mul,
		Div: &timeouts.Div,
	}}, nil
}

func NewExecutor(ts TimeoutsSerice) *Executor {
	return &Executor{timeoutsService: ts}
}
