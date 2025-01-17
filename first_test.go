package go_func

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var firstAndLastTests = []struct {
	ss    []float64
	first float64
	last  float64
}{
	{
		nil,
		0,
		0,
	},
	{
		[]float64{100},
		100,
		100,
	},
	{
		[]float64{1, 2},
		1,
		2,
	},
	{
		[]float64{1, 2, 3},
		1,
		3,
	},
}

func TestFirst(t *testing.T) {
	for _, test := range firstAndLastTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.first, First(test.ss))
		})
	}
}
