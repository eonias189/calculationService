package use_auth

import (
	"net/http"
	"time"

	"github.com/eonias189/calculationService/backend/internal/lib/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
)

func MakeHandler(userService UserService, tokenAuth *jwtauth.JWTAuth, expTime time.Duration) http.Handler {
	mux := chi.NewRouter()
	ex := NewExecutor(userService, tokenAuth, expTime)
	mux.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		var body RegisterBody
		err := utils.BindAndValidate(w, r, &body)
		if err != nil {
			utils.HandleError(w, r, err, http.StatusBadRequest)
			return
		}

		resp, err := ex.Register(body)
		if err != nil {
			utils.HandleError(w, r, err, http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, resp)
	})

	mux.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		var body LoginBody
		err := utils.BindAndValidate(w, r, &body)
		if err != nil {
			utils.HandleError(w, r, err, http.StatusUnauthorized)
			return
		}

		resp, err := ex.Login(body)
		if err != nil {
			utils.HandleError(w, r, err, http.StatusUnauthorized)
			return
		}

		render.JSON(w, r, resp)
	})

	return mux
}
