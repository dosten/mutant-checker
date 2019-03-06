package mutant

import (
	"errors"
	"testing"

	"github.com/dosten/mutant-checker/store"
)

func TestIsMutation(t *testing.T) {
	tests := map[string]struct {
		matrix []string
		result bool
		err    error
	}{
		"1x1 matrix should not fail": {
			matrix: []string{
				"A",
			},
			result: false,
			err:    nil,
		},
		"2x2 matrix should not fail": {
			matrix: []string{
				"AT",
				"CA",
			},
			result: false,
			err:    nil,
		},
		"one diagonal mutation is not enough": {
			matrix: []string{
				"ATGC",
				"CAGT",
				"TTAT",
				"AGAA",
			},
			result: false,
			err:    nil,
		},
		"one vertical mutation is not enough": {
			matrix: []string{
				"ATGC",
				"CAGT",
				"TTGT",
				"AGGA",
			},
			result: false,
			err:    nil,
		},
		"one horizontal mutation is not enough": {
			matrix: []string{
				"AAAA",
				"CAGT",
				"TGTT",
				"AGAA",
			},
			result: false,
			err:    nil,
		},
		"malformed matrix": {
			matrix: []string{
				"AATAG",
				"CAGT",
				"TTTT",
				"AGAA",
			},
			result: false,
			err:    errors.New("length mismatch at row 1, expected 4 but got 5"),
		},
		"4x4 matrix with two mutations": {
			matrix: []string{
				"AAAA",
				"CAGT",
				"TTTT",
				"AGAA",
			},
			result: true,
			err:    nil,
		},
		"4x4 matrix with two perpendicular mutations": {
			matrix: []string{
				"AAAAG",
				"CAGTG",
				"TTATG",
				"AGAAG",
				"AAGAT",
			},
			result: true,
			err:    nil,
		},
		"4x4 matrix with horizontal and perpendicular up mutations": {
			matrix: []string{
				"AAAAG",
				"CAGGG",
				"TTGTG",
				"AGAAG",
				"AAGAT",
			},
			result: true,
			err:    nil,
		},
		"4x4 matrix with horizontal and perpendicular up mutations twice": {
			matrix: []string{
				"AAAAG",
				"CAGGG",
				"TTGTG",
				"AGAAG",
				"AAGAT",
			},
			result: true,
			err:    nil,
		},
	}

	storer := store.NewMemoryStorer()
	checker := NewSimpleChecker(storer)

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := checker.IsMutant(test.matrix)
			if test.result != result {
				t.Fatalf("expected result '%t' but got '%t'", test.result, result)
			}
			if test.err == err {
				t.SkipNow()
			}
			if test.err == nil && err != nil {
				t.Fatal("expected no error")
			}
			if test.err != nil && err == nil {
				t.Fatal("expected error")
			}
			if test.err.Error() != err.Error() {
				t.Fatalf("expected error '%s' but got '%s'", test.err, err)
			}
		})
	}
}
