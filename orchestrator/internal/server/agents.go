package server

import (
	"net/http"

	"github.com/eonias189/calculationService/orchestrator/internal/service"
	"github.com/go-chi/render"
)

type AgentsGetter interface {
	Agents(limit, offset int) ([]service.Agent, error)
}

type GetAgentsResp struct {
	Agents []service.Agent `json:"agents"`
}

func handleGetAgents(ag AgentsGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit := GetIntFromQuery(r.URL.Query(), "limit", 5)
		offset := GetIntFromQuery(r.URL.Query(), "offser", 0)
		agents, err := ag.Agents(limit, offset)

		if err != nil {
			HandleError(w, r, err, http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, GetAgentsResp{Agents: agents})

	}
}
