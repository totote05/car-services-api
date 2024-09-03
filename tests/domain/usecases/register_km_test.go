package usecases_test

import (
	"context"
	"errors"
	"testing"

	"car-services-api.totote05.ar/domain/adapters"
	"car-services-api.totote05.ar/domain/entities"
	"car-services-api.totote05.ar/domain/usecases"
	"car-services-api.totote05.ar/tests/dsl"
	"car-services-api.totote05.ar/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterKm(t *testing.T) {
	anError := errors.New("an error")
	assert := assert.New(t)
	ctx := context.Background()
	vehicle := dsl.NewValidVehicleOne()
	km := dsl.NewValidKmOne()
	km2 := dsl.NewValidKmTwo()
	suite := []struct {
		name         string
		vehicle      *entities.Vehicle
		km           entities.Km
		expected     *entities.Km
		list         []entities.Km
		shouldGet    bool
		shouldSave   bool
		err          error
		serviceError error
		vehicleError error
		saveError    error
	}{
		{
			name:         "register service should fail when km service fails",
			vehicle:      &vehicle,
			km:           km,
			list:         []entities.Km{},
			shouldGet:    true,
			err:          adapters.ErrNotFound,
			serviceError: adapters.ErrNotFound,
		},
		{
			name:       "register first km successfully",
			vehicle:    &vehicle,
			km:         km,
			expected:   &km,
			list:       []entities.Km{},
			shouldGet:  true,
			shouldSave: true,
		},
		{
			name:       "register second km successfully",
			vehicle:    &vehicle,
			km:         km2,
			expected:   &km2,
			list:       []entities.Km{km},
			shouldGet:  true,
			shouldSave: true,
		},
		{
			name:    "register older km with newer value should fail",
			vehicle: &vehicle,
			km:      dsl.NewInvalidKm(),
			list: []entities.Km{
				km,
				km2,
			},
			shouldGet: true,
			err:       usecases.ErrInvalidKmData,
		},
		{
			name:         "register should fail when vehicle service fails",
			km:           km,
			shouldGet:    true,
			err:          adapters.ErrNotFound,
			vehicleError: adapters.ErrNotFound,
		},
		{
			name:    "register should fail when km service fails",
			vehicle: &vehicle,
			km:      dsl.NewInvalidKmTwo(),
			list: []entities.Km{
				km,
				km2,
			},
			shouldGet: true,
			err:       usecases.ErrInvalidKmData,
		},
		{
			name:    "register same km value should fail",
			vehicle: &vehicle,
			km:      dsl.NewInvalidKmThree(),
			list: []entities.Km{
				km,
			},
			shouldGet: true,
			err:       usecases.ErrInvalidKmData,
		},
		{
			name:    "register same date value should fail",
			vehicle: &vehicle,
			km:      dsl.NewInvalidKmFour(),
			list: []entities.Km{
				km,
			},
			shouldGet: true,
			err:       usecases.ErrInvalidKmData,
		},
		{
			name:       "register should fail when save fails",
			vehicle:    &vehicle,
			km:         km,
			list:       []entities.Km{},
			shouldGet:  true,
			shouldSave: true,
			err:        anError,
			saveError:  anError,
		},
		{
			name:    "register should fail when km is invalid",
			vehicle: &vehicle,
			km:      dsl.NewInvalidKmFive(),
			list:    nil,
			err:     entities.ErrKmHasZeroValue,
		},
	}

	for _, test := range suite {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			vehicleAdapter := mocks.NewVehicle(t)
			if test.shouldGet {
				vehicleAdapter.On("Get", ctx, vehicle.ID).Return(test.vehicle, test.vehicleError)
			}

			kmAdapter := mocks.NewKm(t)
			if test.shouldSave {
				kmAdapter.On("Save", ctx, vehicle.ID, mock.Anything).Return(test.err)
			}

			if test.vehicleError == nil && test.shouldGet {
				kmAdapter.On("GetAll", ctx, vehicle.ID).Return(test.list, test.serviceError)
			}

			usecase := usecases.NewRegisterKm(kmAdapter, vehicleAdapter)

			result, err := usecase.Execute(ctx, vehicle.ID, test.km)

			assert.Equal(test.err, err)
			if test.expected != nil {
				assert.Equal(test.expected.Value, result.Value)
				assert.Equal(test.expected.Date, result.Date)
			}
		})
	}
}
