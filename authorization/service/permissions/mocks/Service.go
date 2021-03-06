// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/gnnchya/PosCoffee/authorize/domain"
	mock "github.com/stretchr/testify/mock"

	out "github.com/gnnchya/PosCoffee/authorize/service/permissions/out"

	permissionsin "github.com/gnnchya/PosCoffee/authorize/service/permissions/permissionsin"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, input
func (_m *Service) Create(ctx context.Context, input *permissionsin.CreateInput) (string, error) {
	ret := _m.Called(ctx, input)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *permissionsin.CreateInput) string); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *permissionsin.CreateInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, input
func (_m *Service) Delete(ctx context.Context, input *permissionsin.DeleteInput) error {
	ret := _m.Called(ctx, input)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *permissionsin.DeleteInput) error); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List provides a mock function with given fields: ctx, opt
func (_m *Service) List(ctx context.Context, opt *domain.PageOption) (int, []*out.PermissionsView, error) {
	ret := _m.Called(ctx, opt)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, *domain.PageOption) int); ok {
		r0 = rf(ctx, opt)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 []*out.PermissionsView
	if rf, ok := ret.Get(1).(func(context.Context, *domain.PageOption) []*out.PermissionsView); ok {
		r1 = rf(ctx, opt)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*out.PermissionsView)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *domain.PageOption) error); ok {
		r2 = rf(ctx, opt)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Read provides a mock function with given fields: ctx, input
func (_m *Service) Read(ctx context.Context, input *permissionsin.ReadInput) (*out.PermissionsView, error) {
	ret := _m.Called(ctx, input)

	var r0 *out.PermissionsView
	if rf, ok := ret.Get(0).(func(context.Context, *permissionsin.ReadInput) *out.PermissionsView); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*out.PermissionsView)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *permissionsin.ReadInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, input
func (_m *Service) Update(ctx context.Context, input *permissionsin.UpdateInput) error {
	ret := _m.Called(ctx, input)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *permissionsin.UpdateInput) error); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
