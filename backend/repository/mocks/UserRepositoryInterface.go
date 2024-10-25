// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	models "github.com/R-Thibault/OrgaJobSearch/backend/models"
	mock "github.com/stretchr/testify/mock"
)

// UserRepositoryInterface is an autogenerated mock type for the UserRepositoryInterface type
type UserRepositoryInterface struct {
	mock.Mock
}

// GetUserByEmail provides a mock function with given fields: email
func (_m *UserRepositoryInterface) GetUserByEmail(email string) (*models.User, error) {
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

// GetUserByID provides a mock function with given fields: ID
func (_m *UserRepositoryInterface) GetUserByID(ID uint) (*models.User, error) {
	ret := _m.Called(ID)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByID")
	}

	var r0 *models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*models.User, error)); ok {
		return rf(ID)
	}
	if rf, ok := ret.Get(0).(func(uint) *models.User); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByUUID provides a mock function with given fields: uuid
func (_m *UserRepositoryInterface) GetUserByUUID(uuid string) (*models.User, error) {
	ret := _m.Called(uuid)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByUUID")
	}

	var r0 *models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*models.User, error)); ok {
		return rf(uuid)
	}
	if rf, ok := ret.Get(0).(func(string) *models.User); ok {
		r0 = rf(uuid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PreRegisterUser provides a mock function with given fields: user
func (_m *UserRepositoryInterface) PreRegisterUser(user models.User) (*models.User, error) {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for PreRegisterUser")
	}

	var r0 *models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(models.User) (*models.User, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(models.User) *models.User); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(models.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveUser provides a mock function with given fields: user
func (_m *UserRepositoryInterface) SaveUser(user models.User) error {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for SaveUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateEmail provides a mock function with given fields: email
func (_m *UserRepositoryInterface) ValidateEmail(email string) error {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for ValidateEmail")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserRepositoryInterface creates a new instance of UserRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepositoryInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepositoryInterface {
	mock := &UserRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
