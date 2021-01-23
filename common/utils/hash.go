package utils

import (
	"crypto/sha256"
)

//--- FUNCTIONS

func Hash(input []byte) ([]byte, error) {
	hasher := sha256.New()
	_, err := hasher.Write(input)
	return hasher.Sum(nil), err
}
