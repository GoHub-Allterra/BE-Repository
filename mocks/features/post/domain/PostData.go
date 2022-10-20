// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "gohub/features/post/domain"

	mock "github.com/stretchr/testify/mock"
)

// PostData is an autogenerated mock type for the PostData type
type PostData struct {
	mock.Mock
}

// DeletedId provides a mock function with given fields: param, token
func (_m *PostData) DeletedId(param int, token int) (int, error) {
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

// GetAll provides a mock function with given fields:
func (_m *PostData) GetAll() ([]domain.Post, error) {
	ret := _m.Called()

	var r0 []domain.Post
	if rf, ok := ret.Get(0).(func() []domain.Post); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Post)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllPostsByID provides a mock function with given fields: id
func (_m *PostData) GetAllPostsByID(id uint) ([]domain.Post, error) {
	ret := _m.Called(id)

	var r0 []domain.Post
	if rf, ok := ret.Get(0).(func(uint) []domain.Post); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Post)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: param
func (_m *PostData) GetById(param int) (domain.Post, error) {
	ret := _m.Called(param)

	var r0 domain.Post
	if rf, ok := ret.Get(0).(func(int) domain.Post); ok {
		r0 = rf(param)
	} else {
		r0 = ret.Get(0).(domain.Post)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: data, token
func (_m *PostData) Insert(data domain.Post, token int) (int, error) {
	ret := _m.Called(data, token)

	var r0 int
	if rf, ok := ret.Get(0).(func(domain.Post, int) int); ok {
		r0 = rf(data, token)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Post, int) error); ok {
		r1 = rf(data, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutPost provides a mock function with given fields: param, token, data
func (_m *PostData) PutPost(param int, token int, data domain.Post) (int, error) {
	ret := _m.Called(param, token, data)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, int, domain.Post) int); ok {
		r0 = rf(param, token, data)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, domain.Post) error); ok {
		r1 = rf(param, token, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPostData interface {
	mock.TestingT
	Cleanup(func())
}

// NewPostData creates a new instance of PostData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPostData(t mockConstructorTestingTNewPostData) *PostData {
	mock := &PostData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
