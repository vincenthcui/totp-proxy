package proxy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLast(t *testing.T) {
	cases := [][3]int64{
		// now, interval, expected
		{1524486261, 10, 1524486269},
		{1524486264, 10, 1524486269},
		{1524486264, 60, 1524486299},
		{1524486288, 60, 1524486299},
	}

	for _, c := range cases {
		now, interval, expected := c[0], c[1], c[2]
		val := last(now, interval)
		assert.Equal(t, expected, val)
	}
}
