package flagx

import (
	"slices"
	"strconv"
	"testing"

	"github.com/earthboundkid/assert"
)

func TestJoin(t *testing.T) {
	assert.Continues(t).
		Equal("", joinFunc(slices.Values[[]int](nil), ", ", strconv.Itoa)).
		Equal("1", joinFunc(slices.Values([]int{1}), ", ", strconv.Itoa)).
		Equal("1, 2", joinFunc(slices.Values([]int{1, 2}), ", ", strconv.Itoa))
}
