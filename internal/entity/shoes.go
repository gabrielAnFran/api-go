package entity

import (
	"errors"
)

type Shoes struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Size  float64 `json:"size"`
	Brand string  `json:"brand"`
}

var (
	ErrNameIsRequired    = errors.New("name is required")
	ErrPriceIsRequired   = errors.New("price is required")
	ErrSizeIsRequired    = errors.New("size is required")
	ErrBrandIsIdRequired = errors.New("brand is required")
	ErrPriceIsInvalid    = errors.New("price is invalid")
)

func NewShoes(name, brand string, price, size float64) (*Shoes, error) {
	shoes := Shoes{
		Name:  name,
		Price: price,
		Size:  size,
		Brand: brand,
	}

	err := shoes.Validate()

	if err != nil {
		return nil, err
	}

	return &shoes, nil

}

func (s *Shoes) Validate() error {
	if s.Name == "" {
		return ErrNameIsRequired
	}

	if s.Brand == "" {
		return ErrBrandIsIdRequired
	}

	if s.Price == 0 {
		return ErrPriceIsRequired
	}

	if s.Price < 0 {
		return ErrPriceIsInvalid
	}

	return nil
}
