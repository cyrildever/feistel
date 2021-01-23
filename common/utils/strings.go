package utils

import (
	"errors"
	"strings"
)

// Add adds two strings in the sense that each charCode are added
func Add(str1, str2 string) (string, error) {
	if len(str1) != len(str2) {
		return "", errors.New("to be added, byte arrays must be of the same length")
	}
	added := ""
	for i := 0; i < len(str1); i++ {
		added += string(str1[i] + str2[i])
	}
	return added, nil
}

// Extract returns an extraction of the passed string of the desired length from the passed start index.
// If the desired length is too long, the string is repeated.
func Extract(from string, startIndex, desiredLength int) string {
	startIndex = startIndex % len(from)
	lengthNeeded := startIndex + desiredLength
	repetitions := lengthNeeded/len(from) + 1
	repeated := strings.Repeat(from, repetitions)
	return repeated[startIndex : startIndex+desiredLength]
}

// Split splits a byte array in two equal parts
func Split(item string) (left, right string, err error) {
	if len(item)%2 != 0 {
		err = errors.New("invalid string length: cannot be split")
		return
	}
	half := len(item) / 2
	left = item[:half]
	right = item[half:]
	return
}
