package use_auth

import (
	"net/http"
	"time"

	"github.com/eonias189/calculationService/backend/internal/lib/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
)

func RegisterHandler(e *Executor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body RegisterBody
		err := utils.BindAndValidate(w, r, &body)
		if err != nil {
			utils.HandleError(w, r, err, http.StatusBadRequest)
			return
		}

		resp, err := e.Register(body)
		if err != nil {
			utils.HandleError(w, r, err, http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, resp)
	}
}

func LoginHandler(e *Executor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body LoginBody
		err := utils.BindAndValidate(w, r, &body)
		if err != nil {
			utils.HandleError(w, r, err, http.StatusUnauthorized)
			return
		}

		resp, err := e.Login(body)
		if err != nil {
			utils.HandleError(w, r, err, http.StatusUnauthorized)
			return
		}

		render.JSON(w, r, resp)
	}
}
func MakeHandler(userService UserService, tokenAuth *jwtauth.JWTAuth, expTime time.Duration) http.Handler {
	r := chi.NewRouter()
	e := NewExecutor(userService, tokenAuth, expTime)
	r.Post("/register", RegisterHandler(e))
	r.Post("/login", LoginHandler(e))
	return r
}
