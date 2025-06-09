package handlers

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kylep342/mendel/internal/constants"
	"github.com/kylep342/mendel/pkg/responses"
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

func (h *InternalHandler) RegisterRoutes(g *gin.Engine, basePath string) {
	rg := g.Group(basePath)
	rg.GET(constants.RouteHealth, h.Healthcheck)
	rg.GET(constants.RouteEnv, h.EnvCheck)
}

func (h *InternalHandler) Healthcheck(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.envConfig.Server.ReadTimeout)

	defer cancel()

	// initialize per-component stats
	componentStats := map[string]bool{
		"http": true,
		"db":   false,
	}

	// check readiness per component
	err := h.dbConn.PingContext(ctx)
	if err == nil {
		componentStats["db"] = true
	}

	// Respond
	for key := range componentStats {
		if !componentStats[key] {
			responses.RespondError(c, componentStats, http.StatusInternalServerError)
		}
	}

	responses.RespondData(c, componentStats)
}

// EnvConfig responds to a request to expose the internal server config
//   - environment variables for the given server
//
// Returns 404 in production
func (h *InternalHandler) EnvCheck(c *gin.Context) {
	if h.envConfig.App.Environment == constants.EnvProduction {
		responses.RespondError(c, "not found", http.StatusNotFound)
	} else {
		responses.RespondData(c, h.envConfig)
	}
}
