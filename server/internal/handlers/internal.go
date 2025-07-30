package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kylep342/mendel/internal/constants"
	"github.com/kylep342/mendel/pkg/responses"
)

type InternalHandler struct {
	pool      *pgxpool.Pool
	envConfig *constants.EnvConfig
}

const (
	keyDb   = "db"
	keyHttp = "http"
)

func NewInternalHandler(pool *pgxpool.Pool, envConfig *constants.EnvConfig) *InternalHandler {
	return &InternalHandler{
		pool:      pool,
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
		keyHttp: true,
		keyDb:   false,
	}

	// check readiness per component
	err := h.pool.Ping(ctx)
	if err == nil {
		componentStats[keyDb] = true
	}

	// Respond
	for key := range componentStats {
		if !componentStats[key] {
			responses.RespondError(c, componentStats, http.StatusInternalServerError)
		}
	}

	responses.RespondData(c, componentStats, http.StatusOK)
}

// EnvConfig responds to a request to expose the internal server config
// Returns 404 in production
func (h *InternalHandler) EnvCheck(c *gin.Context) {
	if h.envConfig.App.Environment == constants.EnvProduction {
		responses.RespondError(c, "not found", http.StatusNotFound)
	} else {
		responses.RespondData(c, h.envConfig, http.StatusOK)
	}
}
