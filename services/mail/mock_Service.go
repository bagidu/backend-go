// Code generated by mockery v2.3.0. DO NOT EDIT.

package mail

import mock "github.com/stretchr/testify/mock"

// MockService is an autogenerated mock type for the Service type
type MockService struct {
	mock.Mock
}

// Send provides a mock function with given fields: _a0
func (_m *MockService) Send(_a0 *Mail) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*Mail) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
