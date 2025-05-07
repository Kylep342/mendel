package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kylep342/mendel/db"
	"github.com/kylep342/mendel/models"
)

type CRUDHandler[T any] struct {
	Table db.CRUDTable[T]
	New   func() *T
}

func (h *CRUDHandler[T]) RegisterRoutes(r chi.Router, basePath string) {
	r.Route(basePath, func(r chi.Router) {
		r.Get("/", h.GetAll)
		r.Post("/", h.Create)
		r.Get("/{id}", h.GetByID)
		r.Put("/{id}", h.Update)
		r.Delete("/{id}", h.Delete)
	})
}

func (h *CRUDHandler[T]) GetAll(w http.ResponseWriter, r *http.Request) {
	items, err := h.Table.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(items)
}

func (h *CRUDHandler[T]) Create(w http.ResponseWriter, r *http.Request) {
	item := h.New()
	if err := json.NewDecoder(r.Body).Decode(item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Table.Create(item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func (h *CRUDHandler[T]) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	item, err := h.Table.GetByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	json.NewEncoder(w).Encode(item)
}

func (h *CRUDHandler[T]) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	item := h.New()
	if identifiable, ok := any(item).(models.Model); ok {
		identifiable.SetID(id)
	} else {
		http.Error(w, "Item does not implement Identifiable", http.StatusInternalServerError)
		return
	}

	if err := h.Table.Update(item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(item)
}

func (h *CRUDHandler[T]) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := h.Table.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
