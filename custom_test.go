package feistel_test

import (
	"testing"

	"github.com/cyrildever/feistel"
	utls "github.com/cyrildever/go-utls/common/utils"
	"gotest.tools/assert"
)

// TestCustomCipher ...
func TestCustomCipher(t *testing.T) {
	ref := "Edgewhere"
	// Identical to cipher_test.go#TestEncrypt
	keys := []string{
		"8ed9dcc1701c064f0fd7ae235f15143f989920e0ee9658bb7882c8d7d5f05692",
		"8ed9dcc1701c064f0fd7ae235f15143f989920e0ee9658bb7882c8d7d5f05692",
		"8ed9dcc1701c064f0fd7ae235f15143f989920e0ee9658bb7882c8d7d5f05692",
		"8ed9dcc1701c064f0fd7ae235f15143f989920e0ee9658bb7882c8d7d5f05692",
		"8ed9dcc1701c064f0fd7ae235f15143f989920e0ee9658bb7882c8d7d5f05692",
		"8ed9dcc1701c064f0fd7ae235f15143f989920e0ee9658bb7882c8d7d5f05692",
		"8ed9dcc1701c064f0fd7ae235f15143f989920e0ee9658bb7882c8d7d5f05692",
		"8ed9dcc1701c064f0fd7ae235f15143f989920e0ee9658bb7882c8d7d5f05692",
		"8ed9dcc1701c064f0fd7ae235f15143f989920e0ee9658bb7882c8d7d5f05692",
		"8ed9dcc1701c064f0fd7ae235f15143f989920e0ee9658bb7882c8d7d5f05692",
	}
	expected := "3d7c0a0f51415a521054"
	cipher := feistel.NewCustomCipher(keys)
	found, err := cipher.Encrypt(ref)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, utls.ToHex(found), expected)

	// Another test with custom keys
	keys = []string{
		"1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
		"9876543210fedcba9876543210fedcba9876543210fedcba9876543210fedcba",
		"abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789",
	}
	expected = "445951465c5a19613633"
	cipher = feistel.NewCustomCipher(keys)
	found, err = cipher.Encrypt(ref)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, utls.ToHex(found), expected)

	deciphered, err := cipher.Decrypt(found)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, deciphered, ref)
}
