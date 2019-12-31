package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	c := &Calculator{}

	assert.Equal(t, 3, c.Sum(1, 2))
}

func TestPow(t *testing.T) {
	c := &Calculator{}

	assert.Equal(t, 8, c.Pow(2, 3))
	assert.Equal(t, 0, c.Pow(0, 5))
	assert.Equal(t, 1, c.Pow(3, 0))
	assert.Equal(t, 3, c.Pow(3, 1))
	assert.Equal(t, 1, c.Pow(1, 1000))
}

func BenchmarkPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := &Calculator{}
		c.Pow(1, 2)
	}
}
