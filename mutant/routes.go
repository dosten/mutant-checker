package mutant

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

// Routes returns a new subrouter to be mounted in
// the main router
func Routes(checker Checker) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", checkMutant(checker))
	return r
}

func checkMutant(checker Checker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload CheckPayload
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		isMutant, err := checker.IsMutant(payload.DNA)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if isMutant {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusForbidden)
		}
	}
}
