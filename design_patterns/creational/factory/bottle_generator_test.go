package factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInteractWithBottles(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		name                    string
		bottlesBrands           []string
		expectedCocaColaBottles int
		expectedPepsiBottles    int
	}{
		{"2 coca-cola bottles, 1 pepsi bottle", []string{CocaCola, CocaCola, Pepsi}, 2, 1},
		{"1 coca-cola bottle, 2 pepsi bottles", []string{CocaCola, Pepsi, Pepsi}, 1, 2},
		{"3 coca-cola bottles", []string{CocaCola, CocaCola, CocaCola}, 3, 0},
		{"3 pepsi bottles", []string{Pepsi, Pepsi, Pepsi}, 0, 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			generator := SodaBottleGenerator{}

			for _, bottleBrand := range test.bottlesBrands {
				generator.AddBottle(bottleBrand)
			}

			cocaColaBottles, pepsiBottles := InteractWithBottles(generator.Bottles)

			assert.Equal(test.expectedCocaColaBottles, cocaColaBottles, test.name)
			assert.Equal(test.expectedPepsiBottles, pepsiBottles, test.name)
		})
	}
}
