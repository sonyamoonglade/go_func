package go_func

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var topAndBottomTests = []struct {
	ss     []float64
	n      int
	top    []float64
	bottom []float64
}{
	{
		nil,
		1,
		nil,
		nil,
	},
	{
		[]float64{},
		1,
		nil,
		nil,
	},
	{
		[]float64{1.23, 2.34},
		1,
		[]float64{1.23},
		[]float64{2.34},
	},
	{
		[]float64{1.23, 2.34},
		3,
		[]float64{1.23, 2.34},
		[]float64{2.34, 1.23},
	},
	{
		[]float64{1.23, 2.34},
		0,
		nil,
		nil,
	},
	{
		[]float64{1.23, 2.34},
		-1,
		nil,
		nil,
	},
}

func TestBottom(t *testing.T) {
	for _, test := range topAndBottomTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.bottom, Bottom(test.ss, test.n))
		})
	}
}
