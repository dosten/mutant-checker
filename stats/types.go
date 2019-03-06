package stats

// Stats represents the available information about
// the performed mutant checks
type Stats struct {
	Mutants int     `json:"count_mutant_dna"`
	Humans  int     `json:"count_human_dna"`
	Ratio   float64 `json:"ratio"`
}
