package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kylep342/mendel/db"
	"github.com/kylep342/mendel/model"
)

type PlantSpeciesHandler struct {
	Repo *db.PlantSpeciesTable
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
	species, err := h.Repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(species)
}

func (h *PlantSpeciesHandler) Create(w http.ResponseWriter, r *http.Request) {
	var ps model.PlantSpecies
	if err := json.NewDecoder(r.Body).Decode(&ps); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Repo.Create(&ps); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ps)
}

func (h *PlantSpeciesHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ps, err := h.Repo.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	json.NewEncoder(w).Encode(ps)
}

func (h *PlantSpeciesHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var ps model.PlantSpecies
	if err := json.NewDecoder(r.Body).Decode(&ps); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ps.Id = id
	if err := h.Repo.Update(&ps); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(ps)
}

func (h *PlantSpeciesHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := h.Repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
