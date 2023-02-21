package database

import (
	"testing"

	"github.com/gabrielAnFran/api-go/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateNewShoes(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Shoes{})

	shoes, err := entity.NewShoes("air jorda", "naike", 152.55, 10)
	assert.NoError(t, err)

	shoesDB := NewShoes(db)
	err = shoesDB.Create(shoes)
	assert.NoError(t, err)
	assert.NotEmpty(t, shoes.ID)
}

func TestFindAllShoes(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(entity.Shoes{})
	name := "air jorda"
	for i := 1; i < 24; i++ {
		if i == 10 {
			name = "capybara"
		}
		shoes, err := entity.NewShoes(name, "naike", 152.55, 10)

		assert.NoError(t, err)
		db.Create(shoes)
	}

	shoesDB := NewShoes(db)

	shoes, err := shoesDB.FindAll(1, 10, "asc")

	assert.NoError(t, err)
	assert.Len(t, shoes, 10)
	assert.Equal(t, "air jorda", shoes[0].Name)
	assert.Equal(t, "capybara", shoes[9].Name)
}

func TestFindById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(entity.Shoes{})

	shoes, err := entity.NewShoes("a", "naike", 152.55, 10)
	assert.NoError(t, err)
	db.Create(shoes)

}
