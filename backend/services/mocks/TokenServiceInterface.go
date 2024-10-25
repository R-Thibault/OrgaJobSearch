// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	models "github.com/R-Thibault/OrgaJobSearch/backend/models"
	mock "github.com/stretchr/testify/mock"
)

// TokenServiceInterface is an autogenerated mock type for the TokenServiceInterface type
type TokenServiceInterface struct {
	mock.Mock
}

// VerifyToken provides a mock function with given fields: tokenString
func (_m *TokenServiceInterface) VerifyToken(tokenString string) (*models.JWTToken, error) {
	ret := _m.Called(tokenString)

	if len(ret) == 0 {
		panic("no return value specified for VerifyToken")
	}

	var r0 *models.JWTToken
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*models.JWTToken, error)); ok {
		return rf(tokenString)
	}
	if rf, ok := ret.Get(0).(func(string) *models.JWTToken); ok {
		r0 = rf(tokenString)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.JWTToken)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(tokenString)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTokenServiceInterface creates a new instance of TokenServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTokenServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *TokenServiceInterface {
	mock := &TokenServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}