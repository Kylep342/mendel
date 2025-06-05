package handlers

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5"
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

func (h *InternalHandler) RegisterRoutes(r chi.Router, basePath string) {
	r.Route(basePath, func(r chi.Router) {
		r.Get(constants.RouteHealth, h.Healthcheck)
		r.Get(constants.RouteEnv, h.EnvCheck)
	})
}

func (h *InternalHandler) RegisterRoutesGin(g *gin.Engine, basePath string) {
	rg := g.Group(basePath)
	rg.GET(constants.RouteHealth, h.HealthcheckGin)
	rg.GET(constants.RouteEnv, h.EnvCheckGin)
}

// Healthcheck responds to a request to report app health
//   - db
//   - http
func (h *InternalHandler) Healthcheck(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)

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
	// first check for any unhealthy components and respond with error
	for key := range componentStats {
		if !componentStats[key] {
			responses.RespondWithError(w, http.StatusInternalServerError, componentStats)
		}
	}

	// Server is healthy
	responses.RespondWithData(w, http.StatusOK, componentStats)
}

func (h *InternalHandler) HealthcheckGin(c *gin.Context) {
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
			c.AbortWithStatusJSON(http.StatusInternalServerError, componentStats)
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": componentStats})
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

// EnvConfig responds to a request to expose the internal server config
//   - environment variables for the given server
//
// Returns 404 in production
func (h *InternalHandler) EnvCheckGin(c *gin.Context) {
	if h.envConfig.App.Environment == constants.EnvProduction {
		c.AbortWithStatusJSON(http.StatusNotFound, "not found")
	} else {
		c.JSON(http.StatusOK, gin.H{"data": h.envConfig})
	}
}
