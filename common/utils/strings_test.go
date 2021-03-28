package utils_test

import (
	"testing"

	"github.com/cyrildever/feistel/common/utils"
	"gotest.tools/assert"
)

// TestSplit ...
func TestSplit(t *testing.T) {
	left := "edge"
	right := "where"
	edgewhere := left + right
	leftPart, rightPart, err := utils.Split(edgewhere)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, leftPart, left)
	assert.Equal(t, rightPart, right)
	assert.Assert(t, len(leftPart) != len(rightPart))
	assert.Assert(t, len(leftPart)+len(rightPart) == len(edgewhere))

	balanced := "balanced"
	leftPart, rightPart, err = utils.Split(balanced)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, leftPart, "bala")
	assert.Equal(t, rightPart, "nced")
	assert.Assert(t, len(balanced)%2 == 0)
	assert.Assert(t, len(leftPart) == len(rightPart))
}
