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

type PlantSpeciesHandler struct {
	Table *db.PlantSpeciesTable
}

func NewPlantSpeciesHandler(table *db.PlantSpeciesTable) *PlantSpeciesHandler {
	return &PlantSpeciesHandler{
		Table: table,
	}
}

func (h *PlantSpeciesHandler) RegisterRoutes(r chi.Router) {
	r.Route("/species", func(r chi.Router) {
		r.Get("/", h.GetAll)
		r.Post("/", h.Create)
		r.Get("/{id}", h.GetByID)
		r.Put("/{id}", h.Update)
		r.Delete("/{id}", h.Delete)
	})
}

func (h *PlantSpeciesHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	species, err := h.Table.GetAll()
	if err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	json.NewEncoder(w).Encode(species)
}

func (h *PlantSpeciesHandler) Create(w http.ResponseWriter, r *http.Request) {
	var ps models.PlantSpecies
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

func (h *PlantSpeciesHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ps, err := h.Table.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			responses.RespondWithError(w, http.StatusNotFound, "Not Found")
		} else {
			responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	json.NewEncoder(w).Encode(ps)
}

func (h *PlantSpeciesHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var ps models.PlantSpecies
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

func (h *PlantSpeciesHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := h.Table.Delete(id); err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
