package sha256

import (
	"crypto/sha256"
	"encoding/hex"
)

// CalculateSHA256 generates a SHA-256 hash for a given string.
// It takes a string as input and returns the hexadecimal representation
// of the hash.
func CalculateSHA256(input string) string {
	// Create a new SHA-256 hash object.
	hasher := sha256.New()

	// Write the input string's byte slice to the hasher.
	// The Write method never returns an error, so we can ignore the return value.
	hasher.Write([]byte(input))

	// Sum the hash and return the byte slice of the hash sum.
	// This finalizes the hash computation.
	hashBytes := hasher.Sum(nil)

	// Encode the hash byte slice to a hexadecimal string and return it.
	return hex.EncodeToString(hashBytes)
}
