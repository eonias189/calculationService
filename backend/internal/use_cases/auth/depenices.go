package use_auth

import "github.com/eonias189/calculationService/backend/internal/service"

type UserService interface {
	Add(user service.User) (int64, error)
	GetByLogin(login string) (service.User, error)
}
