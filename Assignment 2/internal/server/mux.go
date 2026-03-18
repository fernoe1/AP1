package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fernoe1/Assignment2_TemirlanSapargali/internal/model"
	"github.com/fernoe1/Assignment2_TemirlanSapargali/internal/store"
)

func NewMux(store *store.Store[string, string], stats *model.Stats) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /data", addItem(store, stats))
	mux.HandleFunc("GET /data", getItems(store, stats))
	mux.HandleFunc("GET /data/{key}", getSpecificItem(store, stats))
	mux.HandleFunc("DELETE /data/{key}", deleteItem(store, stats))
	mux.HandleFunc("GET /stats", getStats(stats))

	return mux
}

func getStats(stats *model.Stats) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(stats.Snapshot())
	}
}

func deleteItem(store *store.Store[string, string], stats *model.Stats) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stats.IncRequests()
		key := r.PathValue("key")
		ok := store.Delete(key)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func getSpecificItem(store *store.Store[string, string], stats *model.Stats) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stats.IncRequests()
		key := r.PathValue("key")
		item, ok := store.Get(key)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(item)
	}
}

func getItems(store *store.Store[string, string], stats *model.Stats) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stats.IncRequests()
		items := store.Snapshot()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(items)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func addItem(store *store.Store[string, string], stats *model.Stats) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stats.IncRequests()
		var item model.Item
		err := json.NewDecoder(r.Body).Decode(&item)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		if item.Name == "" {
			http.Error(w, "item name is required", http.StatusBadRequest)

			return
		}

		amount, err := strconv.Atoi(item.Amount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		if amount < 0 {
			http.Error(w, "item amount must be positive", http.StatusBadRequest)

			return
		}

		store.Set(item.Name, item.Amount)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(item)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
