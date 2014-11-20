package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwap(t *testing.T) {
	first, second := Swap("hello", "world")
	assert.Equal(t, first, "world", "swap should be equal 'world'")
	assert.Equal(t, second, "hello", "swap not should be equal 'hello'")
}
