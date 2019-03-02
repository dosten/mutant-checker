package util

import (
	"testing"
)

func TestEncodeDNA(t *testing.T) {
	tests := []struct {
		matrix []string
		result string
	}{
		{
			matrix: []string{
				"A",
			},
			result: "559aead08264d5795d3909718cdd05abd49572e84fe55590eef31a88a08fdffd",
		},
		{
			matrix: []string{
				"AG",
				"GA",
			},
			result: "c5f775c57db5bb200f463db8e27c9fa837c38fa0936b82ec4a13c9da6e03cc47",
		},
		{
			matrix: []string{
				"AAA",
				"CAG",
				"TGT",
			},
			result: "14654bd34a65a657131e2c7f9231716a8e6f59d8c693c6cd43f88e6c9431d40a",
		},
		{
			matrix: []string{
				"AAAA",
				"CAGT",
				"TGTT",
				"AGAA",
			},
			result: "aca122913c1edc6d7f55ab55936dfd095f3e02ceff7a6dc32228d764ddf65849",
		},
	}

	for _, test := range tests {
		result := EncodeDNA(test.matrix)
		if test.result != result {
			t.Errorf("expected '%s' but got '%s'", test.result, result)
		}
	}
}
