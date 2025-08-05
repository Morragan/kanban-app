package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		resp := map[string]string{"message": "pong"}

		err := json.NewEncoder(w).Encode(resp)

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
