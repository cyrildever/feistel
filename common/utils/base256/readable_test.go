package base256_test

import (
	"testing"

	"github.com/cyrildever/feistel/common/utils/base256"
	utls "github.com/cyrildever/go-utls/common/utils"
	"gotest.tools/assert"
)

// TestBase256 ...
func TestBase256(t *testing.T) {
	assert.Equal(t, len(base256.CHARSET), 256)
	assert.Equal(t, base256.CharAt(0), "!")
	assert.Equal(t, base256.CharAt(255), "ǿ")

	ref := base256.Readable("g¨«©½¬©·©")
	data := []byte("Edgewhere")
	b256 := base256.ToBase256Readable(data)
	assert.Equal(t, b256, ref)
	hex := "456467657768657265"
	assert.Equal(t, b256.ToHex(), hex)

	fpeEdgewhere := base256.Readable("K¡(#q|r5*")
	fpeBytes := []byte{42, 93, 7, 2, 79, 90, 80, 20, 9}
	assert.DeepEqual(t, fpeBytes, utls.Must(utls.FromHex("2a5d07024f5a501409")))
	fpeB256 := base256.ToBase256Readable(fpeBytes)
	assert.DeepEqual(t, fpeB256, fpeEdgewhere)
	assert.DeepEqual(t, fpeB256.Bytes(), fpeBytes)
	hex = "2a5d07024f5a501409"
	assert.Equal(t, fpeB256.ToHex(), hex)
}
