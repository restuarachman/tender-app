// Code generated by mockery v2.28.2. DO NOT EDIT.

package mocks

import (
	entity "myapp/src/user/entity"

	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// FindAll provides a mock function with given fields:
func (_m *UserRepository) FindAll() ([]entity.User, error) {
	ret := _m.Called()

	var r0 []entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]entity.User, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []entity.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByEmail provides a mock function with given fields: email
func (_m *UserRepository) FindByEmail(email string) (entity.User, error) {
	ret := _m.Called(email)

	var r0 entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (entity.User, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) entity.User); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindById provides a mock function with given fields: id
func (_m *UserRepository) FindById(id uint) (entity.User, error) {
	ret := _m.Called(id)

	var r0 entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (entity.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) entity.User); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: _a0
func (_m *UserRepository) Save(_a0 entity.User) (entity.User, error) {
	ret := _m.Called(_a0)

	var r0 entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(entity.User) (entity.User, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(entity.User) entity.User); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	if rf, ok := ret.Get(1).(func(entity.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0, id
func (_m *UserRepository) Update(_a0 entity.User, id uint) (entity.User, error) {
	ret := _m.Called(_a0, id)

	var r0 entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(entity.User, uint) (entity.User, error)); ok {
		return rf(_a0, id)
	}
	if rf, ok := ret.Get(0).(func(entity.User, uint) entity.User); ok {
		r0 = rf(_a0, id)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	if rf, ok := ret.Get(1).(func(entity.User, uint) error); ok {
		r1 = rf(_a0, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
