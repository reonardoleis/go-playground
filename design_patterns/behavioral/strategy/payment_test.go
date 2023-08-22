package strategy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPay(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		name        string
		countryCode string
		amount      float32
		expected    float32
	}{
		{"brazil payment test", "BR", 100, 330},
		{"monaco payment test", "MC", 100, 130},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			provider := PaymentProvider{}
			result := provider.Pay(test.countryCode, test.amount)
			assert.Equal(test.expected, result, "payment test failed")
		})
	}
}
