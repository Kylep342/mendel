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
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": items})
}

func (h *CRUDHandler[T, PT]) Create(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.Env.Server.WriteTimeout)

	defer cancel()

	item := h.New()
	if err := json.NewDecoder(c.Request.Body).Decode(item); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	if err := h.Table.Create(ctx, item); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": item})
}

func (h *CRUDHandler[T, PT]) GetByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.Env.Server.ReadTimeout)

	defer cancel()

	id := c.Param("id")
	item, err := h.Table.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": item})
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
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": ok})
		return
	}

	if err := h.Table.Update(ctx, item); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": item})
}

func (h *CRUDHandler[T, PT]) Delete(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.Env.Server.WriteTimeout)

	defer cancel()

	id := c.Param("id")
	if err := h.Table.Delete(ctx, id); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}
