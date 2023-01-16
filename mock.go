// Code generated by mockery v2.15.0. DO NOT EDIT.

package go_unleash

import (
	context "github.com/Unleash/unleash-client-go/v3/context"

	mock "github.com/stretchr/testify/mock"
)

// UnleashMock is an autogenerated mock type for the IUnleash type
type UnleashMock struct {
	mock.Mock
}

// IsEnabled provides a mock function with given fields: flagName
func (_m *UnleashMock) IsEnabled(flagName string) bool {
	ret := _m.Called(flagName)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(flagName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsEnabledContext provides a mock function with given fields: flagName, ctx
func (_m *UnleashMock) IsEnabledContext(flagName string, ctx context.Context) bool {
	ret := _m.Called(flagName, ctx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, context.Context) bool); ok {
		r0 = rf(flagName, ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewUnleashMock interface {
	mock.TestingT
	Cleanup(func())
}

// NewUnleashMock creates a new instance of UnleashMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUnleashMock(t mockConstructorTestingTNewUnleashMock) *UnleashMock {
	mock := &UnleashMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
