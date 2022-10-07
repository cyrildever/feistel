package base256_test

import (
	"testing"

	"github.com/cyrildever/feistel"
	"github.com/cyrildever/feistel/common/utils/base256"
	"github.com/cyrildever/feistel/common/utils/hash"
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

	src := "my-source-data"
	cipher := feistel.NewFPECipher(hash.SHA_256, "some-32-byte-long-key-to-be-safe", 128)
	b256, err := cipher.Encrypt(src)
	assert.NilError(t, err)
	byteArr := []byte{62, 125, 126, 123, 99, 124, 118, 109, 108, 121, 97, 49, 33, 101}
	assert.DeepEqual(t, b256.Bytes(), byteArr)
	ascii := b256.String(true)
	assert.Equal(t, ascii, ">}~{c|vmlya1!e")
	assert.Equal(t, len(ascii), len(src))
	notAscii := b256.String()
	assert.Equal(t, notAscii, "`ÃÄÁ§Â¼²±¿¥RB©")
	// Both results are visually 14 characters long, but the latter uses characters that range over 127, thus resulting in a longer underlying byte slice as computed by the len() function
	assert.Assert(t, ascii != notAscii)
	// But, as you can see, if you transform both strings in runes, than they are both equal - and of length 14 like the original source)
	assert.Equal(t, len([]rune(ascii)), 14)
	assert.Equal(t, len([]rune(notAscii)), 14)
	assert.Equal(t, len(src), 14)
}
