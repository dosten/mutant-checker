package stats

import (
	"encoding/json"
	"net/http"

	"github.com/dosten/mutant-checker/store"
	"github.com/go-chi/chi"
)

// Routes returns a new subrouter to be mounted in
// the main router
func Routes(storer store.Storer) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", getStats(storer))
	return r
}

func getStats(storer store.Storer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stats := GetStats(storer)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(stats)
	}
}
