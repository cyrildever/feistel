package feistel

import (
	"errors"

	"github.com/cyrildever/feistel/common/padding"
	"github.com/cyrildever/feistel/common/utils"
	utls "github.com/cyrildever/go-utls/common/utils"
	"github.com/cyrildever/go-utls/common/xor"
)

//--- TYPES

// CustomCipher uses custom keys instead of the SHA-256 hashing function to provide a new key at each round.
// The number of rounds is then determined by the number of keys provided.
type CustomCipher struct {
	Keys []string
}

//--- METHODS

// Encrypt ...
func (cc CustomCipher) Encrypt(src string) (ciphered []byte, err error) {
	if len(src) == 0 {
		return
	}
	// Apply the Feistel cipher
	data := padding.Apply([]byte(src))
	left, right, err := utils.Split(string(data))
	if err != nil {
		return
	}
	parts := []string{left, right}
	for i := 0; i < len(cc.Keys); i++ {
		left = right
		rnd, e := cc.round(parts[1], i)
		if e != nil {
			err = e
			return
		}
		right, err = xor.String(parts[0], rnd)
		if err != nil {
			return
		}
		parts = []string{left, right}
	}
	ciphered = []byte(parts[0] + parts[1])
	return
}

// Decrypt ...
func (cc CustomCipher) Decrypt(ciphered []byte) (string, error) {
	if len(ciphered) == 0 {
		return "", nil
	}
	if len(ciphered)%2 != 0 {
		return "", errors.New("invalid obfuscated data")
	}
	// Apply Feistel cipher
	left, right, err := utils.Split(string(ciphered))
	if err != nil {
		return "", err
	}
	for i := 0; i < len(cc.Keys); i++ {
		rnd, err := cc.round(left, len(cc.Keys)-i-1)
		if err != nil {
			return "", err
		}
		tmp, err := xor.String(right, rnd)
		if err != nil {
			return "", err
		}
		right = left
		left = tmp
	}
	return string(padding.Unapply([]byte(left + right))), nil
}

// Feistel implementation

// round is the function applied at each round of the obfuscation process to the right side of the Feistel cipher
func (cc CustomCipher) round(item string, index int) (string, error) {
	addition, err := utils.Add(item, utils.Extract(cc.Keys[index], index, len(item)))
	if err != nil {
		return "", err
	}
	hashed, err := utils.Hash([]byte(addition))
	if err != nil {
		return "", err
	}
	hexHashed := utls.ToHex(hashed)
	return utils.Extract(hexHashed, index, len(item)), nil
}

//--- FUNCTIONS

// NewCustomCipher ...
func NewCustomCipher(keys []string) *CustomCipher {
	return &CustomCipher{
		Keys: keys,
	}
}
