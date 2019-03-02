package mutant

// Checker checks if a DNA sequence corresponds to a mutant
type Checker interface {
	IsMutant([]string) (bool, error)
}
