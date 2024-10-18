// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	models "github.com/R-Thibault/OrgaJobSearch/models"
	mock "github.com/stretchr/testify/mock"
)

// OtpGeneratorServiceInterface is an autogenerated mock type for the OtpGeneratorServiceInterface type
type OtpGeneratorServiceInterface struct {
	mock.Mock
}

// GenerateOTP provides a mock function with given fields: user
func (_m *OtpGeneratorServiceInterface) GenerateOTP(user *models.User) models.OTP {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for GenerateOTP")
	}

	var r0 models.OTP
	if rf, ok := ret.Get(0).(func(*models.User) models.OTP); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(models.OTP)
	}

	return r0
}

// NewOtpGeneratorServiceInterface creates a new instance of OtpGeneratorServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOtpGeneratorServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *OtpGeneratorServiceInterface {
	mock := &OtpGeneratorServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
