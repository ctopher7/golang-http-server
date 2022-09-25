// Code generated by mockery v2.10.0. DO NOT EDIT.

package usecase

import (
	context "context"

	model "github.com/ctopher7/sirclo/v2/berat/model"
	mock "github.com/stretchr/testify/mock"
)

// MockUsecase is an autogenerated mock type for the Usecase type
type MockUsecase struct {
	mock.Mock
}

// GetBeratData provides a mock function with given fields: ctx, id
func (_m *MockUsecase) GetBeratData(ctx context.Context, id int) (model.GetBeratRes, error) {
	ret := _m.Called(ctx, id)

	var r0 model.GetBeratRes
	if rf, ok := ret.Get(0).(func(context.Context, int) model.GetBeratRes); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.GetBeratRes)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertBerat provides a mock function with given fields: ctx, req
func (_m *MockUsecase) InsertBerat(ctx context.Context, req model.Berat) (error, int) {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Berat) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, model.Berat) int); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Get(1).(int)
	}

	return r0, r1
}
