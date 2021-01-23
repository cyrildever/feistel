package feistel

import (
	"errors"
	"strings"

	"github.com/cyrildever/feistel/common/utils"
	"github.com/cyrildever/feistel/padding"
	utls "github.com/cyrildever/go-utls/common/utils"
	"github.com/cyrildever/go-utls/common/xor"
)

//--- TYPES

// Cipher ...
type Cipher struct {
	Key    string
	Rounds int
}

//--- METHODS

// Encrypt ...
func (c Cipher) Encrypt(src []byte) (ciphered []byte, err error) {
	if len(src) == 0 {
		return
	}
	// Apply the Feistel cipher
	data := padding.Apply(src)
	left, right, err := c.split(string(data))
	if err != nil {
		return
	}
	parts := []string{left, right}
	for i := 0; i < c.Rounds; i++ {
		left = right
		tmp, e := c.round(parts[1], i)
		if e != nil {
			err = e
			return
		}
		right, err = xor.String(parts[0], tmp)
		if err != nil {
			return
		}
		parts = []string{left, right}
	}
	ciphered = []byte(parts[0] + parts[1])
	return
}

// Feistel implementation

// add adds two strings in the sense that each charCode are added
func (c Cipher) add(str1, str2 string) (string, error) {
	if len(str1) != len(str2) {
		return "", errors.New("to be added, byte arrays must be of the same length")
	}
	added := ""
	for i := 0; i < len(str1); i++ {
		added += string(str1[i] + str2[i])
	}
	return added, nil
}

// extract returns an extraction of the passed string of the desired length from the passed start index.
// If the desired length is too long, the key string is repeated.
func (c Cipher) extract(from string, startIndex, desiredLength int) string {
	startIndex = startIndex % len(from)
	lengthNeeded := startIndex + desiredLength
	repetitions := lengthNeeded/len(from) + 1
	repeated := strings.Repeat(from, repetitions)
	return repeated[startIndex : startIndex+desiredLength]
}

// round is the function applied at each round of the obfuscation process to the right side of the Feistel cipher
func (c Cipher) round(item string, index int) (string, error) {
	addition, err := c.add(item, c.extract(c.Key, index, len(item)))
	if err != nil {
		return "", err
	}
	hashed, err := utils.Hash([]byte(addition))
	if err != nil {
		return "", err
	}
	hexHashed := utls.ToHex(hashed)
	return c.extract(hexHashed, index, len(item)), nil
}

// split splits a byte array in two equal parts
func (c Cipher) split(item string) (left, right string, err error) {
	if len(item)%2 != 0 {
		err = errors.New("invalid string length: cannot be split")
		return
	}
	half := len(item) / 2
	left = item[:half]
	right = item[half:]
	return
}

//--- FUNCTIONS

// NewCipher ...
func NewCipher(key string, rounds int) *Cipher {
	return &Cipher{
		Key:    key,
		Rounds: rounds,
	}
}
