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
