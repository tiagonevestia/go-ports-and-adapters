package application_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tiagonevestia/go-hexagonal/application"
	mock_application "github.com/tiagonevestia/go-hexagonal/application/mocks"

	"github.com/golang/mock/gomock"
)

func TestProductService_Get(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{Persistence: persistence}

	res, err := service.Get("abc")
	require.Nil(t, err)
	require.Equal(t, product, res)

}

func TestProductService_Create(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{Persistence: persistence}

	res, err := service.Create("Produto 1", 10)
	require.Nil(t, err)
	require.Equal(t, product, res)

}

func TestProductService_Enable(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)

	product.EXPECT().Enable().Return(nil)

	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{Persistence: persistence}

	res, err := service.Enable(product)
	require.Nil(t, err)
	require.Equal(t, product, res)

}

func TestProductService_Disable(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)

	product.EXPECT().Disable().Return(nil)

	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{Persistence: persistence}

	res, err := service.Disable(product)
	require.Nil(t, err)
	require.Equal(t, product, res)

}
