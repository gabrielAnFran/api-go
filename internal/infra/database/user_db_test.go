package database

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/gabrielAnFran/api-go/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("Fran", "f@f.com", "123231")
	userDB := NewUser(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	var foundUser entity.User

	err = db.
		First(&foundUser).
		Where("id = ?", user.ID).Error

	spew.Dump(foundUser)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, foundUser.ID)
	assert.Equal(t, user.Name, foundUser.Name)
	assert.Equal(t, user.Email, foundUser.Email)
	assert.NotNil(t, foundUser.Password)

}
