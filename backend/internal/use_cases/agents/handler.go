package use_agents

import (
	"net/http"

	"github.com/eonias189/calculationService/backend/internal/lib/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func HandleGetAgents(e *Executor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit := utils.GetIntQuery(r.URL.Query(), "limit", 10)
		offset := utils.GetIntQuery(r.URL.Query(), "offset", 0)
		resp, err := e.GetAgents(limit, offset)
		if err != nil {
			utils.HandleError(w, r, err, http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, resp)
	}
}

func MakeHandler(agentService AgentService) http.Handler {
	r := chi.NewRouter()
	e := NewExecutor(agentService)
	r.Get("/", HandleGetAgents(e))
	return r
}
