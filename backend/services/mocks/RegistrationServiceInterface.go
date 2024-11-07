// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	models "github.com/R-Thibault/OrgaJobSearch/backend/models"
	mock "github.com/stretchr/testify/mock"
)

// RegistrationServiceInterface is an autogenerated mock type for the RegistrationServiceInterface type
type RegistrationServiceInterface struct {
	mock.Mock
}

// UserRegistration provides a mock function with given fields: creds
func (_m *RegistrationServiceInterface) UserRegistration(creds models.Credentials) error {
	ret := _m.Called(creds)

	if len(ret) == 0 {
		panic("no return value specified for UserRegistration")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Credentials) error); ok {
		r0 = rf(creds)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRegistrationServiceInterface creates a new instance of RegistrationServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRegistrationServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *RegistrationServiceInterface {
	mock := &RegistrationServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
