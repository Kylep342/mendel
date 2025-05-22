package handlers

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kylep342/mendel/constants"
	"github.com/kylep342/mendel/responses"
)

type InternalHandler struct {
	dbConn    *sql.DB
	envConfig *constants.EnvConfig
}

func NewInternalHandler(dbConn *sql.DB, envConfig *constants.EnvConfig) *InternalHandler {
	return &InternalHandler{
		dbConn:    dbConn,
		envConfig: envConfig,
	}
}

func (h *InternalHandler) RegisterRoutes(r chi.Router, basePath string) {
	r.Route(basePath, func(r chi.Router) {
		r.Get(constants.RouteHealth, h.Healthcheck)
		r.Get(constants.RouteEnv, h.EnvCheck)
	})
}

// Healthcheck responds to a request to report app health
//   - db
//   - http
func (h *InternalHandler) Healthcheck(w http.ResponseWriter, r *http.Request) {

	// initialize per-component stats
	componentStats := map[string]bool{
		"http": true,
		"db":   false,
	}

	// check readiness per component
	err := h.dbConn.Ping()
	if err == nil {
		componentStats["db"] = true
	}

	// Respond
	// first check for any unhealthy components and respond with error
	for key := range componentStats {
		if !componentStats[key] {
			responses.RespondWithError(w, http.StatusInternalServerError, componentStats)
		}
	}

	// Server is healthy
	responses.RespondWithData(w, http.StatusOK, componentStats)
}

// EnvConfig responds to a request to expose the internal server config
//   - environment variables for the given server
//
// Returns 404 in production
func (h *InternalHandler) EnvCheck(w http.ResponseWriter, r *http.Request) {
	if h.envConfig.App.Environment == constants.EnvProduction {
		responses.RespondWithError(w, http.StatusNotFound, "not found")
	} else {
		responses.RespondWithData(w, http.StatusOK, h.envConfig)
	}
}
