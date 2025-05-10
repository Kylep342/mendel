package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kylep342/mendel/db"
	"github.com/kylep342/mendel/models"
	"github.com/kylep342/mendel/responses"
)

type PlantCultivarHandler struct {
	Table *db.PlantCultivarTable
}

func NewPlantCultivarHandler(table *db.PlantCultivarTable) *PlantCultivarHandler {
	return &PlantCultivarHandler{
		Table: table,
	}
}

func (h *PlantCultivarHandler) RegisterRoutes(r chi.Router) {
	r.Route("/plant-cultivar", func(r chi.Router) {
		r.Get("/", h.GetAll)
		r.Post("/", h.Create)
		r.Get("/{id}", h.GetByID)
		r.Put("/{id}", h.Update)
		r.Delete("/{id}", h.Delete)
	})
}

func (h *PlantCultivarHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	Cultivar, err := h.Table.GetAll()
	if err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	json.NewEncoder(w).Encode(Cultivar)
}

func (h *PlantCultivarHandler) Create(w http.ResponseWriter, r *http.Request) {
	var ps models.PlantCultivar
	if err := json.NewDecoder(r.Body).Decode(&ps); err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.Table.Create(&ps); err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ps)
}

func (h *PlantCultivarHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ps, err := h.Table.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			responses.RespondWithError(w, http.StatusNotFound, "Not found")
		} else {
			responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	json.NewEncoder(w).Encode(ps)
}

func (h *PlantCultivarHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var ps models.PlantCultivar
	if err := json.NewDecoder(r.Body).Decode(&ps); err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	ps.Id = id
	if err := h.Table.Update(&ps); err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	json.NewEncoder(w).Encode(ps)
}

func (h *PlantCultivarHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := h.Table.Delete(id); err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
