package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5"
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

func (h *CRUDHandler[T, PT]) RegisterRoutes(r chi.Router, basePath string) {
	r.Route(basePath, func(r chi.Router) {
		r.Get("/", h.GetAll)
		r.Post("/", h.Create)
		r.Get("/{id}", h.GetByID)
		r.Put("/{id}", h.Update)
		r.Delete("/{id}", h.Delete)
	})
}

func (h *CRUDHandler[T, PT]) RegisterRoutesGin(g *gin.Engine, basePath string) {
	rg := g.Group(basePath)
	rg.GET("/", h.GetAllGin)
}

func (h *CRUDHandler[T, PT]) GetAll(w http.ResponseWriter, r *http.Request) {
	//TODO: me learn how to pass the config settings in smoothly
	ctx, cancel := context.WithTimeout(r.Context(), 60*time.Second)

	defer cancel()

	items, err := h.Table.GetAll(ctx)
	if err != nil {
		// log.Printf("Error in GetAll: %v", err)
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func (h *CRUDHandler[T, PT]) Create(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 60*time.Second)

	defer cancel()

	item := h.New()
	if err := json.NewDecoder(r.Body).Decode(item); err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.Table.Create(ctx, item); err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func (h *CRUDHandler[T, PT]) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 60*time.Second)

	defer cancel()

	id := chi.URLParam(r, "id")
	item, err := h.Table.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			responses.RespondWithError(w, http.StatusNotFound, "Not found")
		} else {
			responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	json.NewEncoder(w).Encode(item)
}

// TODO: This method is not passing JSON params from HTTP put methods and populating the model (e.g. "name", "taxon")
func (h *CRUDHandler[T, PT]) Update(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 60*time.Second)

	defer cancel()

	id := chi.URLParam(r, "id")
	item := h.New()
	if model, ok := any(item).(models.Model); ok {
		model.SetID(id)
	} else {
		responses.RespondWithError(w, http.StatusInternalServerError, "Item does not implement Model")
		return
	}

	if err := h.Table.Update(ctx, item); err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	json.NewEncoder(w).Encode(item)
}

func (h *CRUDHandler[T, PT]) Delete(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 60*time.Second)

	defer cancel()

	id := chi.URLParam(r, "id")
	if err := h.Table.Delete(ctx, id); err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *CRUDHandler[T, PT]) GetAllGin(c *gin.Context) {
	//TODO: me learn how to pass the config settings in smoothly
	ctx, cancel := context.WithTimeout(context.Background(), h.Env.Server.ReadTimeout)

	defer cancel()

	items, err := h.Table.GetAll(ctx)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *CRUDHandler[T, PT]) CreateGin(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 60*time.Second)

	defer cancel()

	item := h.New()
	if err := json.NewDecoder(r.Body).Decode(item); err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.Table.Create(ctx, item); err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func (h *CRUDHandler[T, PT]) GetByIDGin(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 60*time.Second)

	defer cancel()

	id := chi.URLParam(r, "id")
	item, err := h.Table.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			responses.RespondWithError(w, http.StatusNotFound, "Not found")
		} else {
			responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	json.NewEncoder(w).Encode(item)
}

// TODO: This method is not passing JSON params from HTTP put methods and populating the model (e.g. "name", "taxon")
func (h *CRUDHandler[T, PT]) UpdateGin(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 60*time.Second)

	defer cancel()

	id := chi.URLParam(r, "id")
	item := h.New()
	if model, ok := any(item).(models.Model); ok {
		model.SetID(id)
	} else {
		responses.RespondWithError(w, http.StatusInternalServerError, "Item does not implement Model")
		return
	}

	if err := h.Table.Update(ctx, item); err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	json.NewEncoder(w).Encode(item)
}

func (h *CRUDHandler[T, PT]) DeleteGin(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 60*time.Second)

	defer cancel()

	id := chi.URLParam(r, "id")
	if err := h.Table.Delete(ctx, id); err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
