package feistel_test

import (
	"fmt"
	"testing"

	"github.com/cyrildever/feistel"
	"github.com/cyrildever/feistel/common/utils/hash"
	utls "github.com/cyrildever/go-utls/common/utils"
	"gotest.tools/assert"
)

// TestFPEEncrypt ...
func TestFPEEncrypt(t *testing.T) {
	expected := "2a5d07024f5a501409"
	cipher := feistel.NewFPECipher(hash.SHA_256, "8ed9dcc1701c064f0fd7ae235f15143f989920e0ee9658bb7882c8d7d5f05692", 10)
	found, err := cipher.Encrypt("Edgewhere")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, len(found), 9) // len(Edgewhere)
	assert.Equal(t, utls.ToHex(found), expected)
	assert.Equal(t, len(utls.Must(utls.FromHex("2a5d07024f5a501409"))), 9)
}

// TestFPEDecrypt ...
func TestFPEDecrypt(t *testing.T) {
	expected := "Edgewhere"
	cipher := feistel.NewFPECipher(hash.SHA_256, "8ed9dcc1701c064f0fd7ae235f15143f989920e0ee9658bb7882c8d7d5f05692", 10)
	found, err := cipher.Decrypt(utls.Must(utls.FromHex("3d7c0a0f51415a521054")))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, found, expected)

	expected = "Edgewhere"
	fmt.Println(utls.Must(utls.FromHex("2a5d07024f5a501409")))
	found, err = cipher.Decrypt(utls.Must(utls.FromHex("2a5d07024f5a501409")))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, found, expected)
}
