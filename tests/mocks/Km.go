// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "car-services-api.totote05.ar/domain/entities"
	mock "github.com/stretchr/testify/mock"
)

// Km is an autogenerated mock type for the Km type
type Km struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, vehicleID, kmID
func (_m *Km) Get(ctx context.Context, vehicleID entities.VehicleID, kmID entities.KmID) (*entities.Km, error) {
	ret := _m.Called(ctx, vehicleID, kmID)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *entities.Km
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entities.VehicleID, entities.KmID) (*entities.Km, error)); ok {
		return rf(ctx, vehicleID, kmID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entities.VehicleID, entities.KmID) *entities.Km); ok {
		r0 = rf(ctx, vehicleID, kmID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Km)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, entities.VehicleID, entities.KmID) error); ok {
		r1 = rf(ctx, vehicleID, kmID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: ctx, vehicleID
func (_m *Km) GetAll(ctx context.Context, vehicleID entities.VehicleID) ([]entities.Km, error) {
	ret := _m.Called(ctx, vehicleID)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []entities.Km
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entities.VehicleID) ([]entities.Km, error)); ok {
		return rf(ctx, vehicleID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entities.VehicleID) []entities.Km); ok {
		r0 = rf(ctx, vehicleID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.Km)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, entities.VehicleID) error); ok {
		r1 = rf(ctx, vehicleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, vehicleID, km
func (_m *Km) Save(ctx context.Context, vehicleID entities.VehicleID, km entities.Km) error {
	ret := _m.Called(ctx, vehicleID, km)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entities.VehicleID, entities.Km) error); ok {
		r0 = rf(ctx, vehicleID, km)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, vehicleID, km
func (_m *Km) Update(ctx context.Context, vehicleID entities.VehicleID, km entities.Km) error {
	ret := _m.Called(ctx, vehicleID, km)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entities.VehicleID, entities.Km) error); ok {
		r0 = rf(ctx, vehicleID, km)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewKm creates a new instance of Km. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewKm(t interface {
	mock.TestingT
	Cleanup(func())
}) *Km {
	mock := &Km{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
