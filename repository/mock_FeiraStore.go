// Code generated by mockery 2.7.5. DO NOT EDIT.

package repository

import (
	domain "github.com/MAAARKIN/unico/domain"
	mock "github.com/stretchr/testify/mock"
)

// MockFeiraStore is an autogenerated mock type for the FeiraStore type
type MockFeiraStore struct {
	mock.Mock
}

// Create provides a mock function with given fields: item
func (_m *MockFeiraStore) Create(item domain.FeiraLivre) (uint64, error) {
	ret := _m.Called(item)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(domain.FeiraLivre) uint64); ok {
		r0 = rf(item)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.FeiraLivre) error); ok {
		r1 = rf(item)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *MockFeiraStore) Delete(id uint64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteByRegistro provides a mock function with given fields: registro
func (_m *MockFeiraStore) DeleteByRegistro(registro string) error {
	ret := _m.Called(registro)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(registro)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: id
func (_m *MockFeiraStore) Get(id uint64) (*domain.FeiraLivre, error) {
	ret := _m.Called(id)

	var r0 *domain.FeiraLivre
	if rf, ok := ret.Get(0).(func(uint64) *domain.FeiraLivre); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.FeiraLivre)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: op
func (_m *MockFeiraStore) GetAll(op domain.FeiraFiltro) ([]domain.FeiraLivre, error) {
	ret := _m.Called(op)

	var r0 []domain.FeiraLivre
	if rf, ok := ret.Get(0).(func(domain.FeiraFiltro) []domain.FeiraLivre); ok {
		r0 = rf(op)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.FeiraLivre)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.FeiraFiltro) error); ok {
		r1 = rf(op)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByRegistro provides a mock function with given fields: registro
func (_m *MockFeiraStore) GetByRegistro(registro string) (*domain.FeiraLivre, error) {
	ret := _m.Called(registro)

	var r0 *domain.FeiraLivre
	if rf, ok := ret.Get(0).(func(string) *domain.FeiraLivre); ok {
		r0 = rf(registro)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.FeiraLivre)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(registro)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, item
func (_m *MockFeiraStore) Update(id uint64, item domain.FeiraLivre) (*domain.FeiraLivre, error) {
	ret := _m.Called(id, item)

	var r0 *domain.FeiraLivre
	if rf, ok := ret.Get(0).(func(uint64, domain.FeiraLivre) *domain.FeiraLivre); ok {
		r0 = rf(id, item)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.FeiraLivre)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64, domain.FeiraLivre) error); ok {
		r1 = rf(id, item)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
