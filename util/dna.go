package util

import (
	"crypto/sha256"
	"fmt"
)

// EncodeDNA encodes a DNA matrix to a SHA-256 hash
func EncodeDNA(rows []string) string {
	h := sha256.New()
	for _, row := range rows {
		h.Write([]byte(row))
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
