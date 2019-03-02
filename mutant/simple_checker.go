package mutant

import (
	"fmt"
	"strconv"

	"github.com/dosten/mutant-checker/store"
	"github.com/dosten/mutant-checker/util"
)

const mutantCountKey = "mutant:count"
const humanCountKey = "human:count"

// SimpleChecker is a implementation of a mutant checker,
// it try to find mutations looking for each combination
// possible (horizontal, vertical, diagonal up, diagonal down)
// for each nitrogenous base in the matrix.
// since the given matrix can be any square matrix of any size
// it try to return early as soon the condition to be a mutant is met
type SimpleChecker struct {
	storer store.Storer
}

func (c *SimpleChecker) isMutant(rows []string) (bool, error) {
	// since we consider a human being be a mutant if he
	// have more than one mutation, we need to mantain a
	// mutations count and try to return early as soon as
	// the count is greater than 1
	mutationsCounter := 0

	// since we are going to access different elements of
	// the matrix at runtime we need to make sure the matrix is NxN
	// to avoid out of bounds array access
	matrixSize := len(rows)
	for y, row := range rows {
		if rowSize := len(row); rowSize != matrixSize {
			return false, fmt.Errorf("length mismatch at row %d, expected %d but got %d", y+1, matrixSize, rowSize)
		}
	}

	for y, row := range rows {
		for x, nb := range row {
			// for each iteration we check if the condition
			// to be a mutant is met to avoid unnecessary work
			if mutationsCounter > 1 {
				return true, nil
			}

			nbb := byte(nb)

			// check for a horizontal sequence
			if x+3 < matrixSize && nbb == row[x+1] && nbb == row[x+2] && nbb == row[x+3] {
				mutationsCounter++
			}

			// check for a vertical sequence
			if y+3 < matrixSize && nbb == rows[y+1][x] && nbb == rows[y+2][x] && nbb == rows[y+3][x] {
				mutationsCounter++
			}

			// check for a diagonal down sequence
			if x+3 < matrixSize && y+3 < matrixSize && nbb == rows[y+1][x+1] && nbb == rows[y+2][x+2] && nbb == rows[y+3][x+3] {
				mutationsCounter++
			}

			// check for a diagonal up sequence
			if x-3 >= 0 && y+3 < matrixSize && nbb == rows[y+1][x-1] && nbb == rows[y+2][x-2] && nbb == rows[y+3][x-3] {
				mutationsCounter++
			}
		}
	}

	return mutationsCounter > 1, nil
}

// IsMutant checks if a DNA sequence corresponds to a mutant
func (c *SimpleChecker) IsMutant(rows []string) (bool, error) {
	// this key is used to save the check result in the store
	key := "dna:" + util.EncodeDNA(rows)

	// try to read DNA analysis result from store
	value, err := c.storer.Get(key)
	if err == nil {
		if mutant, err := strconv.ParseBool(value); err == nil {
			return mutant, nil
		}
	}

	// check if a DNA sequence belongs to a mutant
	mutant, err := c.isMutant(rows)
	if err != nil {
		return false, err
	}

	// write result to store
	c.storer.Set(key, strconv.FormatBool(mutant))

	// increment the corresponding key based on the result
	if mutant {
		c.storer.Increment(mutantCountKey)
	} else {
		c.storer.Increment(humanCountKey)
	}

	return mutant, nil
}

// NewSimpleChecker returns a pointer to a SimpleChecker
func NewSimpleChecker(storer store.Storer) *SimpleChecker {
	return &SimpleChecker{storer}
}
