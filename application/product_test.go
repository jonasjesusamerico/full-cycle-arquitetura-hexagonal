package application_test

import (
	"github.com/jonasjesusamerico/full-cycle-arquitetura-hexagonal/application"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()

	require.Nil(t, err)

	product.Price = 0

	err = product.Enable()

	require.Equal(t, "the price most be greater than zero to enable the product", err.Error())
}

func TestProduct_Disabled(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disabled()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disabled()

	require.Equal(t, "the price most be zero in order to have the product disabled", err.Error())
}
