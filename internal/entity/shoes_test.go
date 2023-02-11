package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewShoes(t *testing.T) {
	shoes, err := NewShoes("Air J", "ni", 600, 8)
	assert.Nil(t, err)
	assert.NotNil(t, shoes)
	assert.Equal(t, "Air J", shoes.Name)
	assert.Equal(t, "ni", shoes.Brand)
	assert.Equal(t, 600.0, shoes.Price)
	assert.Equal(t, 8.0, shoes.Size)
}
