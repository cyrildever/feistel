package hash_test

import (
	"testing"

	"github.com/cyrildever/feistel/common/utils/hash"
	utls "github.com/cyrildever/go-utls/common/utils"
	"gotest.tools/assert"
)

// TestH ...
func TestH(t *testing.T) {
	data := []byte("Edgewhere")

	blake2b := "e5ff44a9b2caa01099082dd6e9055ea5d002beea078e9251454494ccf6869b2f"
	found, err := hash.H(data, hash.BLAKE2b)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, blake2b, utls.ToHex(found))

	keccak := "ac501ee78bc9b9429f6b923953946606b260a8de141eb253567342b678bc5f10"
	found, err = hash.H(data, hash.KECCAK)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, keccak, utls.ToHex(found))

	sha256 := "c0c77f225dd222144bc4ef79dca00ab7d955f26da2b1e0f25df81f8a7e86917c"
	found, err = hash.H(data, hash.SHA_256)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, sha256, utls.ToHex(found))

	sha3 := "9d6bf5763cb18bceb7c15270ff8400ae70bf3cd71928463a30f02805d913409d"
	found, err = hash.H(data, hash.SHA_3)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, sha3, utls.ToHex(found))
}
