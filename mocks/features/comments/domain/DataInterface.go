// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "gohub/features/comments/domain"

	mock "github.com/stretchr/testify/mock"
)

// DataInterface is an autogenerated mock type for the DataInterface type
type DataInterface struct {
	mock.Mock
}

// AddComment provides a mock function with given fields: data
func (_m *DataInterface) AddComment(data domain.Comments) (domain.Comments, error) {
	ret := _m.Called(data)

	var r0 domain.Comments
	if rf, ok := ret.Get(0).(func(domain.Comments) domain.Comments); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(domain.Comments)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Comments) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteComent provides a mock function with given fields: param, token
func (_m *DataInterface) DeleteComent(param int, token int) (int, error) {
	ret := _m.Called(param, token)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, int) int); ok {
		r0 = rf(param, token)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(param, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewDataInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewDataInterface creates a new instance of DataInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDataInterface(t mockConstructorTestingTNewDataInterface) *DataInterface {
	mock := &DataInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}