package handlers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComparePrices(t *testing.T) {
	oldPrice1 := 1000
	newPrice1 := 1000
	newPrice2 := 200

	actualRes := ComparePrices(oldPrice1, newPrice1)

	badRes := ComparePrices(oldPrice1, newPrice2)
	assert.Equal(t, actualRes, false)
	assert.NotEqual(t, badRes, false)
}
