package handlers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPriceParser(t *testing.T) {
	var address = "https://edc.sale/ru/aleksandrov/real-estate/sale/land/prodazha-zemelnogo-uchastka-kfkh-812533.html"
	actualResult := 35000000

	res, err := PriceParser(address)

	if assert.NoError(t, err) {
		assert.Equal(t, res, actualResult)
	}
}
