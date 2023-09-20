package utils

import (
	"errors"
	"strings"
)

// Add adds two strings in the sense that each charCode are added
func Add(str1, str2 string) (string, error) {
	if len(str1) != len(str2) {
		return "", errors.New("to be added, items must be of the same length")
	}
	var added strings.Builder
	for i := 0; i < len(str1); i++ {
		added.WriteRune(rune(str1[i] + str2[i]))
	}
	return added.String(), nil
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

// Split splits a byte array in two parts, the first part being one character shorter in case the passed item has odd length
func Split(item string) (left, right string, err error) {
	half := len(item) / 2
	left = item[:half]
	right = item[half:]
	return
}
