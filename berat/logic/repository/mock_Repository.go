// Code generated by mockery v2.10.0. DO NOT EDIT.

package repository

import (
	context "context"

	model "github.com/ctopher7/sirclo/v2/berat/model"
	mock "github.com/stretchr/testify/mock"

	sql "database/sql"
)

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

// GetAllBerat provides a mock function with given fields: ctx
func (_m *MockRepository) GetAllBerat(ctx context.Context) ([]model.Berat, error) {
	ret := _m.Called(ctx)

	var r0 []model.Berat
	if rf, ok := ret.Get(0).(func(context.Context) []model.Berat); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Berat)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBeratByID provides a mock function with given fields: ctx, id
func (_m *MockRepository) GetBeratByID(ctx context.Context, id int) (model.Berat, error) {
	ret := _m.Called(ctx, id)

	var r0 model.Berat
	if rf, ok := ret.Get(0).(func(context.Context, int) model.Berat); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.Berat)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDB provides a mock function with given fields:
func (_m *MockRepository) GetDB() *sql.DB {
	ret := _m.Called()

	var r0 *sql.DB
	if rf, ok := ret.Get(0).(func() *sql.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.DB)
		}
	}

	return r0
}

// InsertBerat provides a mock function with given fields: ctx, tx, req
func (_m *MockRepository) InsertBerat(ctx context.Context, tx *sql.Tx, req model.Berat) error {
	ret := _m.Called(ctx, tx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *sql.Tx, model.Berat) error); ok {
		r0 = rf(ctx, tx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
