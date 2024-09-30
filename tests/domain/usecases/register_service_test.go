package usecases_test

import (
	"context"
	"testing"

	"car-services-api.totote05.ar/domain/entities"
	"car-services-api.totote05.ar/domain/usecases"
	"car-services-api.totote05.ar/tests/dsl"
	"car-services-api.totote05.ar/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestRegisterService(t *testing.T) {
	anError := assert.AnError
	assert := assert.New(t)
	ctx := context.Background()
	vehicle := dsl.NewValidVehicleOne()
	service := dsl.NewValidServiceOne()
	km := dsl.NewValidKm()
	serviceRegister := dsl.NewValidServiceRegister()
	suite := []struct {
		name         string
		vehicle      *entities.Vehicle
		vehicleErr   error
		service      *entities.Service
		serviceErr   error
		km           *entities.Km
		shouldSaveKm bool
		kmSaveErr    error
		kmList       entities.KmList
		kmListErr    error
		shouldSave   bool
		expected     *entities.ServiceRegister
		err          error
	}{
		{
			name:       "vehicle service fails",
			vehicleErr: anError,
			err:        anError,
		},
		{
			name:       "service service fails",
			vehicle:    &vehicle,
			serviceErr: anError,
			err:        anError,
		},
		{
			name:      "km service fails",
			vehicle:   &vehicle,
			service:   service,
			kmListErr: anError,
			err:       anError,
		},
		{
			name:    "invalid km should throw invalid km data error",
			vehicle: &vehicle,
			service: service,
			km:      &km,
			kmList:  entities.KmList{dsl.NewValidKmOne()},
			err:     usecases.ErrInvalidKmData,
		},
		{
			name:         "km service fails when saving",
			vehicle:      &vehicle,
			service:      service,
			km:           &km,
			kmList:       entities.KmList{},
			shouldSaveKm: true,
			kmSaveErr:    anError,
			err:          anError,
		},
		{
			name:         "service register service fails",
			vehicle:      &vehicle,
			service:      service,
			km:           &km,
			kmList:       entities.KmList{},
			shouldSaveKm: true,
			shouldSave:   true,
			err:          anError,
		},
		{
			name:         "service register with no km history should save successfully",
			vehicle:      &vehicle,
			service:      service,
			km:           &km,
			kmList:       entities.KmList{},
			shouldSaveKm: true,
			expected:     &serviceRegister,
			shouldSave:   true,
		},
		{
			name:         "service register with km history should save successfully",
			vehicle:      &vehicle,
			service:      service,
			km:           &km,
			shouldSaveKm: true,
			kmList:       entities.KmList{dsl.NewValidKmTwo()},
			expected:     &serviceRegister,
			shouldSave:   true,
		},
	}

	for _, test := range suite {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			vehicleService := mocks.NewVehicle(t)
			serviceService := mocks.NewService(t)
			kmService := mocks.NewKm(t)
			serviceRegisterService := mocks.NewServiceRegister(t)

			if test.vehicleErr != nil {
				vehicleService.On("Get", ctx, vehicle.ID).Return(nil, test.vehicleErr)
			}
			if test.vehicle != nil {
				vehicleService.On("Get", ctx, vehicle.ID).Return(test.vehicle, nil)
			}

			if test.serviceErr != nil {
				serviceService.On("Get", ctx, service.ID).Return(nil, test.serviceErr)
			}
			if test.service != nil {
				serviceService.On("Get", ctx, service.ID).Return(test.service, nil)
			}

			if test.kmListErr != nil {
				kmService.On("GetAll", ctx, vehicle.ID).Return(test.kmList, test.kmListErr)
			}
			if test.kmList != nil {
				kmService.On("GetAll", ctx, vehicle.ID).Return(test.kmList, nil)
			}

			if test.shouldSaveKm && test.kmSaveErr != nil {
				kmService.On("Save", ctx, vehicle.ID, dsl.AnythingOfType(km)).Return(test.kmSaveErr)
			}
			if test.shouldSaveKm && test.kmSaveErr == nil {
				kmService.On("Save", ctx, vehicle.ID, dsl.AnythingOfType(km)).Return(nil)
			}

			if test.shouldSave && test.err != nil {
				serviceRegisterService.On("Save", ctx, dsl.AnythingOfType(serviceRegister)).Return(test.err)
			}
			if test.shouldSave && test.err == nil {
				serviceRegisterService.On("Save", ctx, dsl.AnythingOfType(serviceRegister)).Return(nil)
			}

			usecase := usecases.NewRegisterService(
				serviceRegisterService,
				vehicleService,
				serviceService,
				kmService,
			)

			result, err := usecase.Execute(ctx, vehicle.ID, service.ID, km)

			assert.Equal(test.err, err)
			if test.expected != nil {
				assert.Equal(test.expected.VehicleID, result.VehicleID)
				assert.Equal(test.expected.ServiceID, result.ServiceID)
				assert.Equal(test.expected.Km.Date, result.Km.Date)
				assert.Equal(test.expected.Km.Value, result.Km.Value)
			}
		})
	}
}
