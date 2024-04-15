package use_auth

import (
	"time"

	"github.com/eonias189/calculationService/backend/internal/service"
	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/crypto/bcrypt"
)

type Executor struct {
	us        UserService
	tokenAuth *jwtauth.JWTAuth
	expTime   time.Duration
}

func (e *Executor) Register(body RegisterBody) (RegisterResponse, error) {
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return RegisterResponse{}, err
	}

	user := service.User{Login: body.Login, HashedPassword: string(hashedPasswordBytes)}
	_, err = e.us.Add(user)
	if err != nil {
		return RegisterResponse{}, err
	}

	return RegisterResponse{}, nil
}

func (e *Executor) Login(body LoginBody) (LoginResponse, error) {
	user, err := e.us.GetByLogin(body.Login)
	if err != nil {
		return LoginResponse{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(body.Password))
	if err != nil {
		return LoginResponse{}, err
	}

	now := time.Now()
	_, token, err := e.tokenAuth.Encode(map[string]interface{}{
		"user_id": user.Id,
		"login":   user.Login,
		"iat":     now.Unix(),
		"nbf":     now.Unix(),
		"exp":     now.Add(e.expTime).Unix(),
	})
	if err != nil {
		return LoginResponse{}, err
	}

	return LoginResponse{Token: token}, nil
}

func NewExecutor(us UserService, tokenAuth *jwtauth.JWTAuth, expTime time.Duration) *Executor {
	return &Executor{us: us, tokenAuth: tokenAuth, expTime: expTime}
}
