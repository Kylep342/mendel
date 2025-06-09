package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kylep342/mendel/internal/constants"
	"github.com/kylep342/mendel/internal/db"
	"github.com/kylep342/mendel/internal/models"
	"github.com/kylep342/mendel/pkg/responses"
)

type CRUDHandler[T interface{}, PT interface {
	~*T
	models.Model
}] struct {
	Env   *constants.EnvConfig
	Table db.CRUDTable[T]
	New   func() PT
}

func NewCRUDHandler[T interface{}, PT interface {
	~*T
	models.Model
}](
	dbConn *sql.DB,
	env *constants.EnvConfig,
	newModelFunc func() PT,
	tableCreator func(d *sql.DB) db.CRUDTable[T],
) *CRUDHandler[T, PT] {
	return &CRUDHandler[T, PT]{
		Table: tableCreator(dbConn),
		New:   newModelFunc,
		Env:   env,
	}
}

func (h *CRUDHandler[T, PT]) RegisterRoutes(g *gin.Engine, basePath string) {
	rg := g.Group(basePath)
	rg.GET("/", h.GetAll)
}

func (h *CRUDHandler[T, PT]) GetAll(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.Env.Server.ReadTimeout)

	defer cancel()

	items, err := h.Table.GetAll(ctx)
	if err != nil {
		responses.RespondError(c, err, http.StatusInternalServerError)
		return
	}
	responses.RespondData(c, items)
}

func (h *CRUDHandler[T, PT]) Create(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.Env.Server.WriteTimeout)

	defer cancel()

	item := h.New()
	if err := json.NewDecoder(c.Request.Body).Decode(item); err != nil {
		responses.RespondError(c, err, http.StatusBadRequest)
		return
	}
	if err := h.Table.Create(ctx, item); err != nil {
		responses.RespondError(c, err, http.StatusInternalServerError)
		return
	}
	responses.RespondData(c, item)
}

func (h *CRUDHandler[T, PT]) GetByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.Env.Server.ReadTimeout)

	defer cancel()

	id := c.Param("id")
	item, err := h.Table.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			responses.RespondError(c, "not found", http.StatusNotFound)
		} else {
			responses.RespondError(c, err, http.StatusInternalServerError)
		}
		return
	}
	responses.RespondData(c, item)
}

// TODO: This method is not passing JSON params from HTTP put methods and populating the model (e.g. "name", "taxon")
func (h *CRUDHandler[T, PT]) Update(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.Env.Server.WriteTimeout)

	defer cancel()

	id := c.Param("id")
	item := h.New()
	if model, ok := any(item).(models.Model); ok {
		model.SetID(id)
	} else {
		responses.RespondError(c, ok, http.StatusInternalServerError)
		return
	}

	if err := h.Table.Update(ctx, item); err != nil {
		responses.RespondError(c, err, http.StatusInternalServerError)
		return
	}
	responses.RespondData(c, item)
}

func (h *CRUDHandler[T, PT]) Delete(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.Env.Server.WriteTimeout)

	defer cancel()

	id := c.Param("id")
	if err := h.Table.Delete(ctx, id); err != nil {
		responses.RespondError(c, err, http.StatusInternalServerError)
		return
	}
	responses.RespondData(c, id)
}
