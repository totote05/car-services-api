package usecases_test

import (
	"context"
	"testing"

	"car-services-api.totote05.ar/domain/adapters"
	"car-services-api.totote05.ar/domain/entities"
	"car-services-api.totote05.ar/domain/usecases"
	"car-services-api.totote05.ar/tests/dsl"
	"car-services-api.totote05.ar/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetService(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	suite := []struct {
		description string
		input       entities.ServiceID
		expected    *entities.Service
		err         error
	}{
		{"failed repository should return error", dsl.NewValidServiceOne().ID, nil, adapters.ErrGetting},
		{"not existing service should return not found error", dsl.NewValidServiceOne().ID, nil, adapters.ErrNotFound},
		{"valid id should return valid service", dsl.NewValidServiceOne().ID, dsl.NewValidServiceOne(), nil},
	}

	for _, test := range suite {
		test := test
		t.Run(test.description, func(t *testing.T) {
			t.Parallel()
			serviceAdapter := mocks.NewService(t)
			serviceAdapter.On("Get", ctx, test.input).Return(test.expected, test.err)

			usecase := usecases.NewGetService(serviceAdapter)
			result, err := usecase.Execute(ctx, test.input)

			if test.err != nil {
				assert.Nil(result)
				assert.ErrorIs(err, test.err)
			} else {
				assert.Nil(err)
				assert.Equal(result, test.expected)
			}
		})
	}
}
