// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	models "github.com/afrizal423/mygram/app/models"
	mock "github.com/stretchr/testify/mock"
)

// IPhotoRepository is an autogenerated mock type for the IPhotoRepository type
type IPhotoRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: req, userID
func (_m *IPhotoRepository) Create(req models.Photo, userID uint) (models.Photo, error) {
	ret := _m.Called(req, userID)

	var r0 models.Photo
	var r1 error
	if rf, ok := ret.Get(0).(func(models.Photo, uint) (models.Photo, error)); ok {
		return rf(req, userID)
	}
	if rf, ok := ret.Get(0).(func(models.Photo, uint) models.Photo); ok {
		r0 = rf(req, userID)
	} else {
		r0 = ret.Get(0).(models.Photo)
	}

	if rf, ok := ret.Get(1).(func(models.Photo, uint) error); ok {
		r1 = rf(req, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id, userID
func (_m *IPhotoRepository) Delete(id uint, userID uint) (models.Photo, error) {
	ret := _m.Called(id, userID)

	var r0 models.Photo
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, uint) (models.Photo, error)); ok {
		return rf(id, userID)
	}
	if rf, ok := ret.Get(0).(func(uint, uint) models.Photo); ok {
		r0 = rf(id, userID)
	} else {
		r0 = ret.Get(0).(models.Photo)
	}

	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(id, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllByUserId provides a mock function with given fields: userID
func (_m *IPhotoRepository) GetAllByUserId(userID uint) ([]models.Photo, error) {
	ret := _m.Called(userID)

	var r0 []models.Photo
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) ([]models.Photo, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(uint) []models.Photo); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Photo)
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
func (_m *IPhotoRepository) GetByUserId(userID uint, id uint) (models.Photo, error) {
	ret := _m.Called(userID, id)

	var r0 models.Photo
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, uint) (models.Photo, error)); ok {
		return rf(userID, id)
	}
	if rf, ok := ret.Get(0).(func(uint, uint) models.Photo); ok {
		r0 = rf(userID, id)
	} else {
		r0 = ret.Get(0).(models.Photo)
	}

	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(userID, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: req, id, userID
func (_m *IPhotoRepository) Update(req models.Photo, id uint, userID uint) (models.Photo, error) {
	ret := _m.Called(req, id, userID)

	var r0 models.Photo
	var r1 error
	if rf, ok := ret.Get(0).(func(models.Photo, uint, uint) (models.Photo, error)); ok {
		return rf(req, id, userID)
	}
	if rf, ok := ret.Get(0).(func(models.Photo, uint, uint) models.Photo); ok {
		r0 = rf(req, id, userID)
	} else {
		r0 = ret.Get(0).(models.Photo)
	}

	if rf, ok := ret.Get(1).(func(models.Photo, uint, uint) error); ok {
		r1 = rf(req, id, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIPhotoRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewIPhotoRepository creates a new instance of IPhotoRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIPhotoRepository(t mockConstructorTestingTNewIPhotoRepository) *IPhotoRepository {
	mock := &IPhotoRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
