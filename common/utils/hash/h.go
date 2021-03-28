package hash

import (
	"github.com/cyrildever/feistel/common/utils"
	"github.com/cyrildever/feistel/exception"
	keccak "github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/sha3"
)

//--- TYPES

// Engine defines a 256-bits hash algorithm to use
type Engine string

const (
	BLAKE2b Engine = "blake-2b-256"
	KECCAK  Engine = "keccak-256"
	SHA_256 Engine = "sha-256"
	SHA_3   Engine = "sha3-256"
)

func IsAvailableEngine(engine Engine) bool {
	// TODO Enrich with new engines
	return engine == BLAKE2b || engine == KECCAK || engine == SHA_256 || engine == SHA_3
}

//--- FUNCTIONS

func H(input []byte, using Engine) ([]byte, error) {
	switch using {
	case BLAKE2b:
		hasher, err := blake2b.New256(nil)
		if err != nil {
			return nil, err
		}
		_, err = hasher.Write(input)
		return hasher.Sum(nil), err
	case KECCAK:
		return keccak.Keccak256(input), nil
	case SHA_256:
		return utils.Hash(input)
	case SHA_3:
		hasher := sha3.New256()
		_, err := hasher.Write(input)
		return hasher.Sum(nil), err
	default:
		return nil, exception.NewUnkownEngineError()
	}
}
