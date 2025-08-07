package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kylep342/mendel/internal/components"
	"github.com/kylep342/mendel/internal/constants"
	"github.com/kylep342/mendel/internal/db"
	"github.com/kylep342/mendel/pkg/responses"
)

// CRUDHandler exposes CRUD operations on Table over HTTP
//
//	Env: for config values
//	Table: the table to CRUD
//	New: Constructor CRUDTable[T]
type CRUDHandler[T any, PT interface {
	~*T
	components.Model
}] struct {
	Env   *constants.EnvConfig
	Table db.CRUDTable[T]
	New   func() PT
}

// NewCRUDHandler is the constructor for CRUDHandler
func NewCRUDHandler[T any, PT interface {
	~*T
	components.Model
}](
	db *pgxpool.Pool,
	env *constants.EnvConfig,
	newModelFunc func() PT,
	tableCreator func(d *pgxpool.Pool) db.CRUDTable[T],
) *CRUDHandler[T, PT] {
	return &CRUDHandler[T, PT]{
		Table: tableCreator(db),
		New:   newModelFunc,
		Env:   env,
	}
}

// RegisterRoutes connects the handlers to an HTTP server
func (h *CRUDHandler[T, PT]) RegisterRoutes(g *gin.Engine, basePath string) {
	rg := g.Group(basePath)
	rg.GET("/", h.GetAll)
	rg.GET("/:id", h.GetByID)
	rg.POST("/", h.Create)
	rg.PUT("/:id", h.Update)
	rg.DELETE("/:id", h.Delete)
}

// GetAll responds to a request with all records from CRUDTable[T]
func (h *CRUDHandler[T, PT]) GetAll(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.Env.Server.ReadTimeout)
	defer cancel()

	items, err := h.Table.GetAll(ctx)
	if err != nil {
		responses.RespondError(c, err.Error(), http.StatusInternalServerError)
		return
	}
	responses.RespondData(c, items, http.StatusOK)
}

// Create responds to a request to add a record to CRUDTable[T]
func (h *CRUDHandler[T, PT]) Create(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.Env.Server.WriteTimeout)
	defer cancel()

	item := h.New()
	if err := json.NewDecoder(c.Request.Body).Decode(item); err != nil {
		responses.RespondError(c, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Table.Create(ctx, item); err != nil {
		responses.RespondError(c, err.Error(), http.StatusInternalServerError)
		return
	}
	responses.RespondData(c, item, http.StatusOK)
}

// GetByID responds to a request with the requested record from CRUDTable[T]
func (h *CRUDHandler[T, PT]) GetByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.Env.Server.ReadTimeout)
	defer cancel()

	id := c.Param("id")
	item, err := h.Table.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			responses.RespondError(c, "not found", http.StatusNotFound)
		} else {
			responses.RespondError(c, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	responses.RespondData(c, item, http.StatusOK)
}

// Update responds to a request to change the requested record in CRUDTable[T]
func (h *CRUDHandler[T, PT]) Update(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.Env.Server.WriteTimeout)
	defer cancel()

	id := c.Param("id")
	item := h.New()
	if model, ok := any(item).(components.Model); ok {
		model.SetID(id)
	} else {
		responses.RespondError(c, "failed to set ID on model", http.StatusInternalServerError)
		return
	}

	if err := json.NewDecoder(c.Request.Body).Decode(item); err != nil {
		responses.RespondError(c, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Table.Update(ctx, item); err != nil {
		responses.RespondError(c, err.Error(), http.StatusInternalServerError)
		return
	}
	responses.RespondData(c, item, http.StatusOK)
}

// Delete responds to a request to remove the requested record from CRUDTable[T]
func (h *CRUDHandler[T, PT]) Delete(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.Env.Server.WriteTimeout)
	defer cancel()

	id := c.Param("id")
	if err := h.Table.Delete(ctx, id); err != nil {
		responses.RespondError(c, err.Error(), http.StatusInternalServerError)
		return
	}
	responses.RespondData(c, id, http.StatusOK)
}
