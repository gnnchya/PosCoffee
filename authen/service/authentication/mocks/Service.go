// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	autenitcationin "github.com/gnnchya/PosCoffee/authen/service/authentication/authenticationin"

	mock "github.com/stretchr/testify/mock"

	out "github.com/gnnchya/PosCoffee/authen/service/authentication/out"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// GenerateToken provides a mock function with given fields: input
func (_m *Service) GenerateToken(input *autenitcationin.LoginInput) (*out.Token, error) {
	ret := _m.Called(input)

	var r0 *out.Token
	if rf, ok := ret.Get(0).(func(*autenitcationin.LoginInput) *out.Token); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*out.Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*autenitcationin.LoginInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Logout provides a mock function with given fields: input
func (_m *Service) Logout(input string) error {
	ret := _m.Called(input)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VerifyToken provides a mock function with given fields: accessToken
func (_m *Service) VerifyToken(accessToken string) (*string, error) {
	ret := _m.Called(accessToken)

	var r0 *string
	if rf, ok := ret.Get(0).(func(string) *string); ok {
		r0 = rf(accessToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(accessToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
