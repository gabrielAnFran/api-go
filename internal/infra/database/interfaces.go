package database

import "github.com/gabrielAnFran/api-go/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(emaild string) (*entity.User, error)
}

type ShoesInterface interface {
	Create(product *entity.Shoes) error
	FindAll(page, limit int, sort string) ([]entity.Shoes, error)
	FindByID(id string) (entity.Shoes, error)
	Update(product *entity.Shoes) error
	Delete(id string) error
}
