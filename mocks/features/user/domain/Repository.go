// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "gohub/features/user/domain"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ID
func (_m *Repository) Delete(ID uint) (domain.Core, error) {
	ret := _m.Called(ID)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func(uint) domain.Core); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Edit provides a mock function with given fields: input
func (_m *Repository) Edit(input domain.Core) (domain.Core, error) {
	ret := _m.Called(input)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func(domain.Core) domain.Core); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Core) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ID
func (_m *Repository) Get(ID uint) (domain.Core, error) {
	ret := _m.Called(ID)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func(uint) domain.Core); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUsername provides a mock function with given fields: name
func (_m *Repository) GetByUsername(name domain.Core) (domain.Core, int) {
	ret := _m.Called(name)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func(domain.Core) domain.Core); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(domain.Core) int); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Get(1).(int)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: newUser
func (_m *Repository) Insert(newUser domain.Core) (domain.Core, error) {
	ret := _m.Called(newUser)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func(domain.Core) domain.Core); ok {
		r0 = rf(newUser)
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Core) error); ok {
		r1 = rf(newUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: input
func (_m *Repository) Login(input domain.Core) (domain.Core, error) {
	ret := _m.Called(input)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func(domain.Core) domain.Core); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Core) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
