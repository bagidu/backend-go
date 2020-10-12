// Code generated by mockery v2.3.0. DO NOT EDIT.

package generated

import mock "github.com/stretchr/testify/mock"

// MockResolverRoot is an autogenerated mock type for the ResolverRoot type
type MockResolverRoot struct {
	mock.Mock
}

// Mutation provides a mock function with given fields:
func (_m *MockResolverRoot) Mutation() MutationResolver {
	ret := _m.Called()

	var r0 MutationResolver
	if rf, ok := ret.Get(0).(func() MutationResolver); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(MutationResolver)
		}
	}

	return r0
}

// Query provides a mock function with given fields:
func (_m *MockResolverRoot) Query() QueryResolver {
	ret := _m.Called()

	var r0 QueryResolver
	if rf, ok := ret.Get(0).(func() QueryResolver); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(QueryResolver)
		}
	}

	return r0
}
