package feistel_test

import (
	"testing"

	"github.com/cyrildever/feistel"
	"github.com/cyrildever/feistel/common/utils/base256"
	"github.com/cyrildever/feistel/common/utils/hash"
	utls "github.com/cyrildever/go-utls/common/utils"
	"gotest.tools/assert"
)

// TestFPEEncrypt ...
func TestFPEEncrypt(t *testing.T) {
	expected := "K¡(#q|r5*"
	cipher := feistel.NewFPECipher(hash.SHA_256, "8ed9dcc1701c064f0fd7ae235f15143f989920e0ee9658bb7882c8d7d5f05692", 10)
	found, err := cipher.Encrypt("Edgewhere")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, found.Len(), len("Edgewhere"))
	assert.Equal(t, found.String(), expected)
	assert.Equal(t, found.ToHex(), "2a5d07024f5a501409")

	blake2, _ := feistel.NewFPECipher(hash.BLAKE2b, "8ed9dcc1701c064f0fd7ae235f15143f989920e0ee9658bb7882c8d7d5f05692", 10).Encrypt("Edgewhere")
	assert.Assert(t, blake2.String() != expected)
	expectedBlake2 := "¼u*$q0up¢"
	assert.Equal(t, blake2.String(), expectedBlake2)
}

// TestFPEDecrypt ...
func TestFPEDecrypt(t *testing.T) {
	nonFPE := "Edgewhere"
	cipher := feistel.NewFPECipher(hash.SHA_256, "8ed9dcc1701c064f0fd7ae235f15143f989920e0ee9658bb7882c8d7d5f05692", 10)
	found, err := cipher.Decrypt(base256.ToBase256Readable(utls.Must(utls.FromHex("3d7c0a0f51415a521054"))))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, found, nonFPE)

	ref := base256.Readable("K¡(#q|r5*")
	expected := "Edgewhere"
	found, err = cipher.Decrypt(base256.ToBase256Readable(utls.Must(utls.FromHex("2a5d07024f5a501409"))))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, found, expected)
	found, _ = cipher.Decrypt(ref)
	assert.Equal(t, found, expected)
	b256, err := base256.HexToReadable("2a5d07024f5a501409")
	if err != nil {
		t.Fatal(err)
	}
	assert.DeepEqual(t, b256, ref)

	fromBlake2 := base256.Readable("¼u*$q0up¢")
	blake2, _ := feistel.NewFPECipher(hash.BLAKE2b, "8ed9dcc1701c064f0fd7ae235f15143f989920e0ee9658bb7882c8d7d5f05692", 10).Decrypt(fromBlake2)
	assert.Equal(t, blake2, expected)
}

// TestFPEEncryptDecrypt ...
func TestFPEEncryptDecrypt(t *testing.T) {
	ref := "Edgewhere"
	cipher := feistel.NewFPECipher(hash.SHA_256, "8ed9dcc1701c064f0fd7ae235f15143f989920e0ee9658bb7882c8d7d5f05692", 255)
	encrypted, err := cipher.Encrypt(ref)
	if err != nil {
		t.Fatal(err)
	}
	decrypted, err := cipher.Decrypt(encrypted)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, decrypted, ref)
}

// TestReadableString ...
func TestReadableString(t *testing.T) {
	source := "my-source-data"
	cipher := feistel.NewFPECipher(hash.SHA_256, "some-32-byte-long-key-to-be-safe", 128)

	obfuscated, err := cipher.EncryptString(source)
	assert.DeepEqual(t, obfuscated.Bytes(), []byte{62, 125, 126, 123, 99, 124, 118, 109, 108, 121, 97, 49, 33, 101})
	assert.NilError(t, err)
	strAscii := obfuscated.String(true)
	assert.Equal(t, strAscii, ">}~{c|vmlya1!e") // Fully readable because, by some chance, the underlying byte slice has all bytes ranging from 33 to 126 (see above)
	assert.Equal(t, len(source), len(strAscii))
	assert.DeepEqual(t, len(obfuscated.Bytes()), len(source)) // Always true

	strNotAscii := obfuscated.String()
	assert.Equal(t, strNotAscii, "`ÃÄÁ§Â¼²±¿¥RB©")
	assert.Assert(t, len(source) != len(strNotAscii)) // Because of the use of multi-byte encoded characters to make them readable
	assert.Equal(t, len(source), len([]rune(strNotAscii)))

	deciphered, err := cipher.DecryptString(obfuscated)
	assert.NilError(t, err)
	assert.Equal(t, deciphered, source)
}

// TestReadableNumber ...
func TestReadableNumber(t *testing.T) {
	source := uint64(123456789)
	cipher := feistel.NewFPECipher(hash.SHA_256, "some-32-byte-long-key-to-be-safe", 128)

	obfuscated, err := cipher.EncryptNumber(source)
	assert.NilError(t, err)
	assert.Equal(t, obfuscated.Uint64(), uint64(22780178))
	assert.Equal(t, obfuscated.ToNumber(), "22780178")
	assert.Equal(t, obfuscated.ToNumber(9), "022780178")

	deobfuscated, err := cipher.DecryptNumber(obfuscated)
	assert.NilError(t, err)
	assert.Equal(t, deobfuscated, source)

	// Numbers below 256 don't preserve length during encryption
	source = uint64(123)

	obfuscated, err = cipher.EncryptNumber(source)
	assert.Error(t, err, "too small to respect length") // Hence the error
	assert.Equal(t, obfuscated.Uint64(), uint64(24359))
	assert.Equal(t, obfuscated.ToNumber(), "24359")

	obfuscated, _ = base256.NumberToReadable(24359)
	deobfuscated, err = cipher.DecryptNumber(obfuscated)
	assert.NilError(t, err)
	assert.Equal(t, deobfuscated, source)
}
