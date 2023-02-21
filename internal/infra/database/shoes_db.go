package database

import (
	"github.com/gabrielAnFran/api-go/internal/entity"
	"gorm.io/gorm"
)

type Shoes struct {
	DB *gorm.DB
}

func NewShoes(db *gorm.DB) *Shoes {
	return &Shoes{DB: db}
}

func (s *Shoes) Create(shoes *entity.Shoes) error {
	return s.DB.Create(shoes).Error
}

func (s *Shoes) FindAll(page, limit int, sort string) ([]entity.Shoes, error) {
	var shoes []entity.Shoes
	var err error

	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page != 0 && limit != 0 {
		err = s.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&shoes).Error
	} else {
		err = s.DB.Order("created_at " + sort).Find(&shoes).Error
	}

	return shoes, err
}
func (s *Shoes) FindById(id string) (*entity.Shoes, error) {
	var shoes entity.Shoes

	err := s.DB.First(&shoes, "id = ?", id).Error

	return &shoes, err
}
func (s *Shoes) Update(shoes *entity.Shoes) error {
	_, err := s.FindById(shoes.ID.String())
	if err != nil {
		return err
	}

	return s.DB.Save(shoes).Error
}

func (s *Shoes) Delete(id string) error {
	shoes, err := s.FindById(id)
	if err != nil {
		return err
	}
	return s.DB.Delete(shoes).Error
}
