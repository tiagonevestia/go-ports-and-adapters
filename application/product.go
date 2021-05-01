package application

import "errors"

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
	DISABLED = "disabled"
	ENABLED = "enabled"
	ERROR_ENABLED = "Para ativar o produto, o preço deve ser maior que zero."
)

type Product struct {
	ID string
	Name string
	Status string
	Price float64
}

func (p *Product) IsValid() (bool, error) {
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
	return nil
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



