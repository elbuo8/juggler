package controllers

import (
	"encoding/json"
	"github.com/elbuo8/juggler/app"
	"github.com/elbuo8/juggler/app/models"
	"github.com/elbuo8/juggler/app/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func GetServices(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db := r.Context().Value("database").(*models.DB)
		services, err := db.GetServicesByPage(1, 127) // Get all for now
		if err != nil {
			app.Logger.Error(err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(services)
	}
}

func GetService(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := mux.Vars(r)["name"]
		db := r.Context().Value("database").(*models.DB)
		service, err := db.GetServiceByName(name)
		if err != nil {
			if err.Error() == "not found" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			app.Logger.Error(err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(service)
	}
}

func CreateService(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var s struct {
			Name string `json:"name"`
		}
		err := utils.ParseJSONFromReq(r, &s)
		if err != nil {
			http.Error(w, "invalid payload", http.StatusBadRequest)
			return
		}

		if s.Name == "" {
			http.Error(w, "missing name", http.StatusBadRequest)
			return
		}

		db := r.Context().Value("database").(*models.DB)
		service, err := db.NewService(s.Name)

		if err != nil {
			app.Logger.Error(err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(service)
	}
}

func DeleteService(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := mux.Vars(r)["name"]
		db := r.Context().Value("database").(*models.DB)
		err := db.DeleteServiceByName(name)
		if err != nil {
			if err.Error() == "not found" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			app.Logger.Error(err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
