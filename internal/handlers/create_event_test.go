package handlers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateEvent(t *testing.T) {

	var storage *Storage
	var address = "google.com"
	var price = 100

	actualObj := &EventDB{
		Id:      12,
		Address: "google.com",
		Price:   100,
	}
	resp, err := storage.Create(address, price)
	if err != nil {
		t.Log(err)
	}
	assert.Equal(t, resp, actualObj)

}
