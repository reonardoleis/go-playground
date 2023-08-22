package mst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTableHeader(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		name     string
		c        int
		expected string
	}{
		{"should return A", 0, "A"},
		{"should return B", 1, "B"},
		{"should return Z", 25, "Z"},
		{"should return AA", 26, "AA"},
		{"should return AB", 27, "AB"},
		{"should return AZ", 51, "AZ"},
		{"should return AAA", 52, "AAA"},
		{"should return AAB", 53, "AAB"},
		{"should return AAAA", 78, "AAAA"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := getTableHeader(test.c)
			assert.Equal(test.expected, result, "result not equal")
		})
	}
}
