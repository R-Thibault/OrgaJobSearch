// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	models "github.com/R-Thibault/OrgaJobSearch/backend/models"
	mock "github.com/stretchr/testify/mock"
)

// UserServiceInterface is an autogenerated mock type for the UserServiceInterface type
type UserServiceInterface struct {
	mock.Mock
}

// EmailValidation provides a mock function with given fields: email
func (_m *UserServiceInterface) EmailValidation(email string) error {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for EmailValidation")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUserByEmail provides a mock function with given fields: email
func (_m *UserServiceInterface) GetUserByEmail(email string) (*models.User, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByEmail")
	}

	var r0 *models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*models.User, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) *models.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByID provides a mock function with given fields: userID
func (_m *UserServiceInterface) GetUserByID(userID uint) (*models.User, error) {
	ret := _m.Called(userID)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByID")
	}

	var r0 *models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*models.User, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(uint) *models.User); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// JobSeekerRegistration provides a mock function with given fields: tokenBody, creds
func (_m *UserServiceInterface) JobSeekerRegistration(tokenBody string, creds models.Credentials) error {
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
func (_m *UserServiceInterface) PreRegisterJobSeeker(email string, careerSuportID *uint) (*models.User, error) {
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

// RegisterUser provides a mock function with given fields: creds
func (_m *UserServiceInterface) RegisterUser(creds models.Credentials) error {
	ret := _m.Called(creds)

	if len(ret) == 0 {
		panic("no return value specified for RegisterUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Credentials) error); ok {
		r0 = rf(creds)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserServiceInterface creates a new instance of UserServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserServiceInterface {
	mock := &UserServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
