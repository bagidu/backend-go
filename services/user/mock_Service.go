// Code generated by mockery v2.3.0. DO NOT EDIT.

package user

import mock "github.com/stretchr/testify/mock"

// MockService is an autogenerated mock type for the Service type
type MockService struct {
	mock.Mock
}

// All provides a mock function with given fields: limit, offset
func (_m *MockService) All(limit int, offset int) ([]*User, error) {
	ret := _m.Called(limit, offset)

	var r0 []*User
	if rf, ok := ret.Get(0).(func(int, int) []*User); ok {
		r0 = rf(limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: _a0
func (_m *MockService) Create(_a0 *User) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*User) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindBy provides a mock function with given fields: field, val
func (_m *MockService) FindBy(field string, val string) (*User, error) {
	ret := _m.Called(field, val)

	var r0 *User
	if rf, ok := ret.Get(0).(func(string, string) *User); ok {
		r0 = rf(field, val)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(field, val)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOne provides a mock function with given fields: id
func (_m *MockService) FindOne(id string) (*User, error) {
	ret := _m.Called(id)

	var r0 *User
	if rf, ok := ret.Get(0).(func(string) *User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0
func (_m *MockService) Update(_a0 *User) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*User) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
