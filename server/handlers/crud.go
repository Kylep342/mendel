package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kylep342/mendel/db"
	"github.com/kylep342/mendel/models"
	"github.com/kylep342/mendel/responses"
)

type CRUDHandler[T interface{}, PT interface {
	~*T
	models.Model
}] struct {
	Table db.CRUDTable[T]
	New   func() PT
}

func NewCRUDHandler[T interface{}, PT interface {
	~*T
	models.Model
}](
	dbConn *sql.DB,
	newModelFunc func() PT,
	tableCreator func(d *sql.DB) db.CRUDTable[T],
) *CRUDHandler[T, PT] {
	return &CRUDHandler[T, PT]{
		Table: tableCreator(dbConn),
		New:   newModelFunc,
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

func (h *CRUDHandler[T, PT]) GetAll(w http.ResponseWriter, r *http.Request) {
	items, err := h.Table.GetAll()
	if err != nil {
		// log.Printf("Error in GetAll: %v", err)
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func (h *CRUDHandler[T, PT]) Create(w http.ResponseWriter, r *http.Request) {
	item := h.New()
	if err := json.NewDecoder(r.Body).Decode(item); err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.Table.Create(item); err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func (h *CRUDHandler[T, PT]) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	item, err := h.Table.GetByID(id)
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

func (h *CRUDHandler[T, PT]) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	item := h.New()
	if model, ok := any(item).(models.Model); ok {
		model.SetID(id)
	} else {
		responses.RespondWithError(w, http.StatusInternalServerError, "Item does not implement Model")
		return
	}

	if err := h.Table.Update(item); err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	json.NewEncoder(w).Encode(item)
}

func (h *CRUDHandler[T, PT]) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := h.Table.Delete(id); err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
