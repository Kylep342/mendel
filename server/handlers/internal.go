package handlers

import (
	"database/sql"
	"net/http"

	"github.com/kylep342/mendel/responses"
)

type InternalHandler struct {
	dbConn *sql.DB
	// logger
}

// Healthcheck is an API for evaluating server component readiness
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
	// first check for any unhealthy components and respond wit error
	for key := range componentStats {
		if !componentStats[key] {
			responses.RespondWithError(w, http.StatusInternalServerError, componentStats)
		}
	}

	// Server is healthy
	responses.RespondWithData(w, http.StatusOK, componentStats)
}
