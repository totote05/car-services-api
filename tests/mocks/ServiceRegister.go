// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "car-services-api.totote05.ar/domain/entities"
	mock "github.com/stretchr/testify/mock"
)

// ServiceRegister is an autogenerated mock type for the ServiceRegister type
type ServiceRegister struct {
	mock.Mock
}

// GetAll provides a mock function with given fields: ctx, vehicleID
func (_m *ServiceRegister) GetAll(ctx context.Context, vehicleID entities.VehicleID) ([]entities.ServiceRegister, error) {
	ret := _m.Called(ctx, vehicleID)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []entities.ServiceRegister
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entities.VehicleID) ([]entities.ServiceRegister, error)); ok {
		return rf(ctx, vehicleID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entities.VehicleID) []entities.ServiceRegister); ok {
		r0 = rf(ctx, vehicleID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.ServiceRegister)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, entities.VehicleID) error); ok {
		r1 = rf(ctx, vehicleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, serviceRegister
func (_m *ServiceRegister) Save(ctx context.Context, serviceRegister entities.ServiceRegister) error {
	ret := _m.Called(ctx, serviceRegister)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entities.ServiceRegister) error); ok {
		r0 = rf(ctx, serviceRegister)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewServiceRegister creates a new instance of ServiceRegister. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServiceRegister(t interface {
	mock.TestingT
	Cleanup(func())
}) *ServiceRegister {
	mock := &ServiceRegister{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
