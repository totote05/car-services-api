package usecases_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"car-services-api.totote05.ar/domain/adapters"
	"car-services-api.totote05.ar/domain/entities"
	"car-services-api.totote05.ar/domain/usecases"
	"car-services-api.totote05.ar/tests/dsl"
	"car-services-api.totote05.ar/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestUpdateKm(t *testing.T) {
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
		list         []entities.Km
		expected     *entities.Km
		shouldGet    bool
		shouldSave   bool
		err          error
		vehicleError error
		getAllError  error
		saveError    error
	}{
		{
			name: "update should fail when km is invalid",
			km:   dsl.NewInvalidKmFive(),
			err:  entities.ErrKmHasZeroValue,
		},
		{
			name:         "update should fail when vehicle service fails on get",
			km:           km,
			shouldGet:    true,
			vehicleError: adapters.ErrNotFound,
			err:          adapters.ErrNotFound,
		},
		{
			name:        "update should fail when km service fails on get all",
			vehicle:     &vehicle,
			km:          km,
			shouldGet:   true,
			getAllError: anError,
			err:         anError,
		},
		{
			name:      "update should fail when km service fails on save",
			vehicle:   &vehicle,
			km:        km,
			shouldGet: true,
			list: []entities.Km{
				dsl.NewValidKmOne(),
			},
			shouldSave: true,
			saveError:  anError,
			err:        anError,
		},
		{
			name:      "update update should fail when km is not found",
			vehicle:   &vehicle,
			km:        km,
			shouldGet: true,
			list: []entities.Km{
				dsl.NewValidKmTwo(),
			},
			err: usecases.ErrKmNotFound,
		},
		{
			name:    "update with invalid older km should fail",
			vehicle: &vehicle,
			km:      dsl.UpdateWith(km, 1050, -1*time.Hour),
			list: []entities.Km{
				km,
				km2,
			},
			shouldGet: true,
			err:       usecases.ErrInvalidKmData,
		},
		{
			name:    "update with invalid newer km should fail",
			vehicle: &vehicle,
			km:      dsl.UpdateWith(km2, 1000, time.Hour),
			list: []entities.Km{
				km,
				km2,
			},
			shouldGet: true,
			err:       usecases.ErrInvalidKmData,
		},
		{
			name:    "update with same date and different value success",
			vehicle: &vehicle,
			km:      dsl.UpdateWith(km, 1010, 0),
			list: []entities.Km{
				km,
				km2,
			},
			shouldGet:  true,
			shouldSave: true,
			expected: &entities.Km{
				ID:    km.ID,
				Value: 1010,
				Date:  km.Date,
			},
		},
		{
			name:    "update with same value and different date success",
			vehicle: &vehicle,
			km:      dsl.UpdateWith(km, 1000, 30*time.Minute),
			list: []entities.Km{
				km,
				km2,
			},
			shouldGet:  true,
			shouldSave: true,
			expected: &entities.Km{
				ID:    km.ID,
				Value: 1000,
				Date:  km.Date.Add(30 * time.Minute),
			},
		},
	}

	for _, test := range suite {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			vehicleAdaper := mocks.NewVehicle(t)
			kmAdapter := mocks.NewKm(t)
			usecase := usecases.NewUpdateKm(kmAdapter, vehicleAdaper)

			if test.shouldGet {
				vehicleAdaper.On("Get", ctx, vehicle.ID).Return(test.vehicle, test.vehicleError)
			}

			if test.vehicleError == nil && test.shouldGet {
				kmAdapter.On("GetAll", ctx, vehicle.ID).Return(test.list, test.getAllError)
			}

			if test.shouldSave {
				kmAdapter.On("Save", ctx, vehicle.ID, test.km).Return(test.saveError)
			}

			result, err := usecase.Execute(ctx, vehicle.ID, km.ID, test.km)

			assert.Equal(test.expected, result)
			assert.Equal(test.err, err)
		})
	}
}
