package application_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tiagonevestia/go-hexagonal/application"
)


func TestProduct_Enable(t *testing.T) {
	product := application.Product{
		ID: "001",
		Name: "Lápis",
		Status: application.ENABLED,
		Price: 10,
	}

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, application.ERROR_ENABLED, err.Error())
}