package application_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"github.com/tiagonevestia/go-hexagonal/application"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{
		ID:     uuid.NewV4().String(),
		Name:   "Lápis",
		Status: application.DISABLED,
		Price:  10,
	}

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, application.ERROR_ENABLED, err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{
		ID:     uuid.NewV4().String(),
		Name:   "Lápis",
		Status: application.ENABLED,
		Price:  0,
	}

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, application.ERROR_DISABLED, err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{
		ID:     uuid.NewV4().String(),
		Name:   "Lápis",
		Status: application.DISABLED,
		Price:  10,
	}

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, application.ERROR_STATUS, err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, application.ERROR_PRICE, err.Error())
}
