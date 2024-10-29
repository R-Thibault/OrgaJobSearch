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

// JobSeekerRegistration provides a mock function with given fields: tokenBody, creds
func (_m *RegistrationServiceInterface) JobSeekerRegistration(tokenBody string, creds models.Credentials) error {
	ret := _m.Called(tokenBody, creds)

	if len(ret) == 0 {
		panic("no return value specified for JobSeekerRegistration")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, models.Credentials) error); ok {
		r0 = rf(tokenBody, creds)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PreRegisterJobSeeker provides a mock function with given fields: email, careerSuportID
func (_m *RegistrationServiceInterface) PreRegisterJobSeeker(email string, careerSuportID *uint) (*models.User, error) {
	ret := _m.Called(email, careerSuportID)

	if len(ret) == 0 {
		panic("no return value specified for PreRegisterJobSeeker")
	}

	var r0 *models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *uint) (*models.User, error)); ok {
		return rf(email, careerSuportID)
	}
	if rf, ok := ret.Get(0).(func(string, *uint) *models.User); ok {
		r0 = rf(email, careerSuportID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string, *uint) error); ok {
		r1 = rf(email, careerSuportID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterCareerCoach provides a mock function with given fields: creds
func (_m *RegistrationServiceInterface) RegisterCareerCoach(creds models.Credentials) error {
	ret := _m.Called(creds)

	if len(ret) == 0 {
		panic("no return value specified for RegisterCareerCoach")
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
