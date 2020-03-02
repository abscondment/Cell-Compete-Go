package groupofcells

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCompete(t *testing.T) {
	var original = []uint8{1, 0, 0, 0, 0, 1, 0, 0}
	var expected = []uint8{0, 1, 0, 0, 1, 0, 1, 0}
	group := NewGroupOfCells(original)

	group.Compete(1)
	assert.Equal(t, group.ToSlice(), expected, "group should match expected value after one iteration")

	group.Compete(2)
	assert.NotEqual(t, group.ToSlice(), expected, "group should NOT match expected value after multiple iterations")
}

func TestCompeteFailure(t *testing.T) {
	var original = []uint8{1, 0, 0, 0, 0, 1, 0, 0}
	var expected = []uint8{0} // this is intentionally wrong
	group := NewGroupOfCells(original)

	group.Compete(1)
	assert.Equal(t, group.ToSlice(), expected, "this test failed intentionally to demonstrate how tests work")
}
