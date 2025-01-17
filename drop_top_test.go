package go_func

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var dropTopTests = []struct {
	ss      []float64
	n       int
	dropTop []float64
}{
	{
		nil,
		1,
		nil,
	},
	{
		[]float64{},
		1,
		nil,
	},
	{
		[]float64{1.23, 2.34},
		-1,
		nil,
	},
	{
		[]float64{1.23, 2.34},
		0,
		[]float64{1.23, 2.34},
	},

	{
		[]float64{1.23, 2.34},
		1,
		[]float64{2.34},
	},
	{
		[]float64{1.23, 2.34},
		2,
		nil,
	},
	{
		[]float64{1.23, 2.34},
		3,
		nil,
	},
}

func TestDropTop(t *testing.T) {
	for _, test := range dropTopTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.dropTop, DropTop(test.ss, test.n))
		})
	}
}
