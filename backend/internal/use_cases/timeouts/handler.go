package use_timeouts

import (
	"net/http"

	"github.com/eonias189/calculationService/backend/internal/lib/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
)

func GetTimeoutsHandler(e *Executor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, err := utils.GetUserIdFromClaims(r)
		if err != nil {
			utils.HandleError(w, r, err, http.StatusBadRequest)
			return
		}

		resp, err := e.GetTimeouts(userId)
		if err != nil {
			utils.HandleError(w, r, err, http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, resp)
	}
}

func PatchTimeoutsHandler(e *Executor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, err := utils.GetUserIdFromClaims(r)
		if err != nil {
			utils.HandleError(w, r, err, http.StatusBadRequest)
			return
		}

		var body PatchTimeoutsBody
		err = utils.BindAndValidate(w, r, &body)
		if err != nil {
			utils.HandleError(w, r, err, http.StatusBadRequest)
			return
		}

		resp, err := e.PatchTimeouts(body, userId)
		if err != nil {
			utils.HandleError(w, r, err, http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, resp)
	}
}

func MakeHandler(ts TimeoutsSerice, tokenAuth *jwtauth.JWTAuth) http.Handler {
	r := chi.NewRouter()
	e := NewExecutor(ts)

	r.Use(jwtauth.Verifier(tokenAuth))

	r.Get("/", GetTimeoutsHandler(e))
	r.Patch("/", PatchTimeoutsHandler(e))

	return r
}
