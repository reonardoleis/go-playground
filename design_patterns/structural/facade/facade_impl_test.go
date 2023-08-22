package facade

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReallyComplexFunctionFacade(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		name     string
		x        int
		a        int
		b        int
		c        int
		expected int
	}{
		{"should return 10", 1, 5, 2, 2, 10},
		{"should return 7", 2, 1, 3, 1, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReallyComplexFunctionFacade(tt.x, tt.a, tt.b, tt.c)
			assert.Equal(tt.expected, result)
		})
	}
}
