package feistel_test

import (
	"testing"

	"github.com/cyrildever/feistel"
	utls "github.com/cyrildever/go-utls/common/utils"
	"gotest.tools/assert"
)

// TestEncrypt ...
func TestEncrypt(t *testing.T) {
	expected := "3d7c0a0f51415a521054"
	cipher := feistel.NewCipher("8ed9dcc1701c064f0fd7ae235f15143f989920e0ee9658bb7882c8d7d5f05692", 10)
	found, err := cipher.Encrypt("Edgewhere")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, utls.ToHex(found), expected)
}

// TestDecrypt ...
func TestDecrypt(t *testing.T) {
	expected := "Edgewhere"
	cipher := feistel.NewCipher("8ed9dcc1701c064f0fd7ae235f15143f989920e0ee9658bb7882c8d7d5f05692", 10)
	found, err := cipher.Decrypt(utls.Must(utls.FromHex("3d7c0a0f51415a521054")))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, found, expected)
}
