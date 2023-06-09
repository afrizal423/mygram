// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	models "github.com/afrizal423/mygram/app/models"
	mock "github.com/stretchr/testify/mock"
)

// ICommentRepository is an autogenerated mock type for the ICommentRepository type
type ICommentRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: req, userID
func (_m *ICommentRepository) Create(req models.Comment, userID uint) (models.Comment, error) {
	ret := _m.Called(req, userID)

	var r0 models.Comment
	var r1 error
	if rf, ok := ret.Get(0).(func(models.Comment, uint) (models.Comment, error)); ok {
		return rf(req, userID)
	}
	if rf, ok := ret.Get(0).(func(models.Comment, uint) models.Comment); ok {
		r0 = rf(req, userID)
	} else {
		r0 = ret.Get(0).(models.Comment)
	}

	if rf, ok := ret.Get(1).(func(models.Comment, uint) error); ok {
		r1 = rf(req, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id, userID
func (_m *ICommentRepository) Delete(id uint, userID uint) (models.Comment, error) {
	ret := _m.Called(id, userID)

	var r0 models.Comment
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, uint) (models.Comment, error)); ok {
		return rf(id, userID)
	}
	if rf, ok := ret.Get(0).(func(uint, uint) models.Comment); ok {
		r0 = rf(id, userID)
	} else {
		r0 = ret.Get(0).(models.Comment)
	}

	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(id, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllByUserId provides a mock function with given fields: userID
func (_m *ICommentRepository) GetAllByUserId(userID uint) ([]models.Comment, error) {
	ret := _m.Called(userID)

	var r0 []models.Comment
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) ([]models.Comment, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(uint) []models.Comment); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Comment)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUserId provides a mock function with given fields: userID, id
func (_m *ICommentRepository) GetByUserId(userID uint, id uint) (models.Comment, error) {
	ret := _m.Called(userID, id)

	var r0 models.Comment
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, uint) (models.Comment, error)); ok {
		return rf(userID, id)
	}
	if rf, ok := ret.Get(0).(func(uint, uint) models.Comment); ok {
		r0 = rf(userID, id)
	} else {
		r0 = ret.Get(0).(models.Comment)
	}

	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(userID, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HitungFoto provides a mock function with given fields: id
func (_m *ICommentRepository) HitungFoto(id int) int64 {
	ret := _m.Called(id)

	var r0 int64
	if rf, ok := ret.Get(0).(func(int) int64); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// Update provides a mock function with given fields: req, id, userID
func (_m *ICommentRepository) Update(req models.Comment, id uint, userID uint) (models.Comment, error) {
	ret := _m.Called(req, id, userID)

	var r0 models.Comment
	var r1 error
	if rf, ok := ret.Get(0).(func(models.Comment, uint, uint) (models.Comment, error)); ok {
		return rf(req, id, userID)
	}
	if rf, ok := ret.Get(0).(func(models.Comment, uint, uint) models.Comment); ok {
		r0 = rf(req, id, userID)
	} else {
		r0 = ret.Get(0).(models.Comment)
	}

	if rf, ok := ret.Get(1).(func(models.Comment, uint, uint) error); ok {
		r1 = rf(req, id, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewICommentRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewICommentRepository creates a new instance of ICommentRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewICommentRepository(t mockConstructorTestingTNewICommentRepository) *ICommentRepository {
	mock := &ICommentRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
