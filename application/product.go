package application

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	getID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

const (
	DISABLED       = "disabled"
	ENABLED        = "enabled"
	ERROR_ENABLED  = "Para ativar o produto, o preço deve ser maior que zero."
	ERROR_DISABLED = "Para desativar o produto, o preço deve ser igual a zero."
	ERROR_STATUS   = "O status deve ser ENABLED ou DISABLED"
	ERROR_PRICE    = "O preço deve ser maior ou igual a zero"
)

type Product struct {
	ID     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Price  float64 `valid:"float,optional"`
	Status string  `valid:"required"`
}

func (p *Product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLED
	}

	if p.Status != ENABLED && p.Status != DISABLED {
		return false, errors.New(ERROR_STATUS)
	}

	if p.Price < 0 {
		return false, errors.New(ERROR_PRICE)
	}

	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}
	return errors.New(ERROR_ENABLED)
}

func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Status = DISABLED
		return nil
	}
	return errors.New(ERROR_DISABLED)
}

func (p *Product) GetID() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
