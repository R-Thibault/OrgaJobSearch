// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	models "github.com/R-Thibault/OrgaJobSearch/backend/models"
	mock "github.com/stretchr/testify/mock"
)

// OTPRepositoryInterface is an autogenerated mock type for the OTPRepositoryInterface type
type OTPRepositoryInterface struct {
	mock.Mock
}

// GetOTPByCode provides a mock function with given fields: otpCode, otpType
func (_m *OTPRepositoryInterface) GetOTPByCode(otpCode string, otpType string) (*models.OTP, error) {
	ret := _m.Called(otpCode, otpType)

	if len(ret) == 0 {
		panic("no return value specified for GetOTPByCode")
	}

	var r0 *models.OTP
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*models.OTP, error)); ok {
		return rf(otpCode, otpType)
	}
	if rf, ok := ret.Get(0).(func(string, string) *models.OTP); ok {
		r0 = rf(otpCode, otpType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.OTP)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(otpCode, otpType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOTPCodeByUserIDandType provides a mock function with given fields: userID, otpType
func (_m *OTPRepositoryInterface) GetOTPCodeByUserIDandType(userID uint, otpType string) (*models.OTP, error) {
	ret := _m.Called(userID, otpType)

	if len(ret) == 0 {
		panic("no return value specified for GetOTPCodeByUserIDandType")
	}

	var r0 *models.OTP
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, string) (*models.OTP, error)); ok {
		return rf(userID, otpType)
	}
	if rf, ok := ret.Get(0).(func(uint, string) *models.OTP); ok {
		r0 = rf(userID, otpType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.OTP)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, string) error); ok {
		r1 = rf(userID, otpType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveOTP provides a mock function with given fields: otp
func (_m *OTPRepositoryInterface) SaveOTP(otp models.OTP) (string, error) {
	ret := _m.Called(otp)

	if len(ret) == 0 {
		panic("no return value specified for SaveOTP")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(models.OTP) (string, error)); ok {
		return rf(otp)
	}
	if rf, ok := ret.Get(0).(func(models.OTP) string); ok {
		r0 = rf(otp)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(models.OTP) error); ok {
		r1 = rf(otp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateOTPCode provides a mock function with given fields: otpID, otpCode, otpType
func (_m *OTPRepositoryInterface) UpdateOTPCode(otpID uint, otpCode string, otpType string) (*models.OTP, error) {
	ret := _m.Called(otpID, otpCode, otpType)

	if len(ret) == 0 {
		panic("no return value specified for UpdateOTPCode")
	}

	var r0 *models.OTP
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, string, string) (*models.OTP, error)); ok {
		return rf(otpID, otpCode, otpType)
	}
	if rf, ok := ret.Get(0).(func(uint, string, string) *models.OTP); ok {
		r0 = rf(otpID, otpCode, otpType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.OTP)
		}
	}

	if rf, ok := ret.Get(1).(func(uint, string, string) error); ok {
		r1 = rf(otpID, otpCode, otpType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewOTPRepositoryInterface creates a new instance of OTPRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOTPRepositoryInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *OTPRepositoryInterface {
	mock := &OTPRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
