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
