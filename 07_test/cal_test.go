package cal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, 10, Add(5, 5))
}

func TestSub(t *testing.T) {
	assert.Equal(t, 4, Sub(5, 1))
}

func BenchmarkAdd2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(i, i+1)
	}
}

func TestMockAdd(t *testing.T) {
}
