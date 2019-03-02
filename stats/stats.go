package stats

import (
	"math"
	"strconv"

	"github.com/dosten/mutant-checker/store"
)

const mutantCountKey = "mutant:count"
const humanCountKey = "human:count"

// GetStats returns a Stats object from a storer
func GetStats(storer store.Storer) *Stats {
	// get the mutants count
	value1, err := storer.Get(mutantCountKey)
	if err != nil {
		value1 = "0"
	}
	mutants, err := strconv.Atoi(value1)
	if err != nil {
		mutants = 0
	}

	// get the humans count
	value2, err := storer.Get(humanCountKey)
	if err != nil {
		value2 = "0"
	}
	humans, err := strconv.Atoi(value2)
	if err != nil {
		mutants = 0
	}

	// get the ratio between humans and mutants
	ratio := math.Round(float64(mutants)/float64(humans)*100) / 100

	return &Stats{
		Mutants: mutants,
		Humans:  humans,
		Ratio:   ratio,
	}
}
