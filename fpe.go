package feistel

import (
	"encoding/binary"
	"math/bits"

	"github.com/cyrildever/feistel/common/utils"
	"github.com/cyrildever/feistel/common/utils/base256"
	"github.com/cyrildever/feistel/common/utils/hash"
	"github.com/cyrildever/feistel/exception"
	utls "github.com/cyrildever/go-utls/common/utils"
	"github.com/cyrildever/go-utls/common/xor"
)

//--- TYPES

// FPECipher builds a format-preserving encrypted cipher using the key with the passed hash engine at each round.
// NB: There must be at least 2 rounds.
type FPECipher struct {
	Engine hash.Engine
	Key    string
	Rounds int
}

//--- METHODS

// Encrypt ...
func (f FPECipher) Encrypt(src string) (ciphered base256.Readable, err error) {
	if len(f.Key) == 0 || f.Rounds < 2 || !hash.IsAvailableEngine(f.Engine) {
		err = exception.NewWrongCipherParametersError()
		return
	}
	if len(src) == 0 {
		return
	}
	// Apply the FPE Feistel cipher
	left, right, err := utils.Split(src)
	if err != nil {
		return
	}
	parts := []string{left, right}
	for i := 0; i < f.Rounds; i++ {
		left = right
		if len(parts[1]) < len(parts[0]) {
			neutral := xor.Neutral("0")
			parts[1] += string(neutral)
		}
		rnd, e := f.round(parts[1], i)
		if e != nil {
			err = e
			return
		}
		tmp := parts[0]
		crop := false
		if len(tmp)+1 == len(rnd) {
			neutral := xor.Neutral(rnd[len(tmp):])
			tmp += string(neutral)
			crop = true
		}
		right, err = xor.String(tmp, rnd)
		if err != nil {
			return
		}
		if crop {
			right = right[:len(right)-1]
		}
		parts = []string{left, right}
	}
	ciphered = base256.ToBase256Readable([]byte(parts[0] + parts[1]))
	return
}

// EncryptNumber ...
func (f FPECipher) EncryptNumber(src uint64) (ciphered base256.Readable, err error) {
	if len(f.Key) == 0 || f.Rounds < 2 || !hash.IsAvailableEngine(f.Engine) {
		err = exception.NewWrongCipherParametersError()
		return
	}

	if src < 256 {
		bytes := make([]byte, 1)
		bytes = append(bytes, uint64ToBytes(src)...)
		res, e := f.Encrypt(string(bytes))
		if e != nil {
			err = e
			return
		}
		ciphered = res
		err = exception.NewTooSmallToPreserveLengthError()
		return
	}

	bytes := uint64ToBytes(src)
	str := string(bytes)

	return f.Encrypt(str)
}

// EncryptString ...
func (f FPECipher) EncryptString(src string) (ciphered base256.Readable, err error) {
	return f.Encrypt(src)
}

// Decrypt ...
func (f FPECipher) Decrypt(ciphered base256.Readable) (string, error) {
	if len(f.Key) == 0 || f.Rounds < 2 || !hash.IsAvailableEngine(f.Engine) {
		return "", exception.NewWrongCipherParametersError()
	}
	if ciphered.IsEmpty() {
		return "", nil
	}
	// Apply the FPE Feistel cipher
	left, right, err := utils.Split(string(ciphered.Bytes()))
	if err != nil {
		return "", err
	}
	// Compensating the way Split() works by moving the first byte at right to the end of left if using an odd number of rounds
	if f.Rounds%2 != 0 && len(left) != len(right) {
		left += right[:1]
		right = right[1:]
	}
	for i := 0; i < f.Rounds; i++ {
		leftRound := left
		if len(left) < len(right) {
			neutral := xor.Neutral("0")
			leftRound += string(neutral)
		}
		rnd, err := f.round(leftRound, f.Rounds-i-1)
		if err != nil {
			return "", err
		}
		rightRound := right
		extended := false
		if len(rightRound)+1 == len(rnd) {
			rightRound += left[len(left)-1:]
			extended = true
		}
		tmp, err := xor.String(rightRound, rnd)
		if err != nil {
			return "", err
		}
		right = left
		if extended {
			tmp = tmp[:len(tmp)-1]
		}
		left = tmp
	}
	return string([]byte(left + right)), nil
}

// DecryptNumber ...
func (f FPECipher) DecryptNumber(ciphered base256.Readable) (uint64, error) {
	deciphered, err := f.Decrypt(ciphered)
	if err != nil {
		return 0, err
	}
	return bytesToUint64([]byte(deciphered))
}

// DecryptString ...
func (f FPECipher) DecryptString(ciphered base256.Readable) (string, error) {
	return f.Decrypt(ciphered)
}

// Feistel implementation

// round is the function applied at each round of the obfuscation process to the right side of the Feistel cipher
func (f FPECipher) round(item string, index int) (string, error) {
	addition, err := utils.Add(item, utils.Extract(f.Key, index, len(item)))
	if err != nil {
		return "", err
	}
	hashed, err := hash.H([]byte(addition), f.Engine)
	if err != nil {
		return "", err
	}
	hexHashed := utls.ToHex(hashed)
	return utils.Extract(hexHashed, index, len(item)), nil
}

//--- FUNCTIONS

// NewFPECipher ...
func NewFPECipher(engine hash.Engine, key string, rounds int) *FPECipher {
	return &FPECipher{
		Engine: engine,
		Key:    key,
		Rounds: rounds,
	}
}

//--- utilities

func uint64ToBytes(x uint64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, x)
	return buf[bits.LeadingZeros64(x)>>3:]
}

func bytesToUint64(x []byte) (uint64, error) {
	if len(x) > 8 {
		return 0, exception.NewNotUint64Error()
	}
	buf := make([]byte, 8)
	for i, b := range x {
		buf[i+8-len(x)] = b
	}
	return binary.BigEndian.Uint64(buf), nil
}
