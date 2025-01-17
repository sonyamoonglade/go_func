package go_func

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var jsonTests = []struct {
	ss         []float64
	jsonString string
}{
	{
		nil,
		`[]`, // Instead of null.
	},
	{
		[]float64{},
		`[]`,
	},
	{
		[]float64{12.3},
		`[12.3]`,
	},
	{
		[]float64{23, -2.5, 3424, 12.3},
		`[23,-2.5,3424,12.3]`,
	},
}

func TestJSONBytes(t *testing.T) {
	for _, test := range jsonTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, []byte(test.jsonString), JSONBytes(test.ss))
		})
	}
}
