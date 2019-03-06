package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dosten/mutant-checker/mutant"
	"github.com/dosten/mutant-checker/stats"
	"github.com/dosten/mutant-checker/store"
	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
)

func main() {
	opts, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opts)

	storer := store.NewRedisStorer(client)
	checker := mutant.NewSimpleChecker(storer)

	r := chi.NewRouter()

	r.Mount("/mutant", mutant.Routes(checker))
	r.Mount("/stats", stats.Routes(storer))

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
