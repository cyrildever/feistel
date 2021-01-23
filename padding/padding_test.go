package padding_test

import (
	"testing"

	"github.com/cyrildever/feistel/padding"
	"gotest.tools/assert"
)

// TestPadding ...
func TestPadding(t *testing.T) {
	expected := "Edgewhere"
	padded := padding.Apply([]byte(expected))
	assert.Equal(t, len(padded), len(expected)+1)
	assert.Assert(t, len(padded)%2 == 0)
	assert.Equal(t, string(padded), "Edgewhere")
	found := string(padding.Unapply(padded))
	assert.Equal(t, expected, found)
}
