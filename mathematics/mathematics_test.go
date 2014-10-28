package mathematics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	var sum int
	sum = Sum(30, 17)
	assert.Equal(t, sum, 47, "Sum should be 47")
}
