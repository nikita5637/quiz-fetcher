// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/nikita5637/quiz-fetcher/internal/pkg/model"
	mock "github.com/stretchr/testify/mock"
)

// SyncLogFacade is an autogenerated mock type for the SyncLogFacade type
type SyncLogFacade struct {
	mock.Mock
}

type SyncLogFacade_Expecter struct {
	mock *mock.Mock
}

func (_m *SyncLogFacade) EXPECT() *SyncLogFacade_Expecter {
	return &SyncLogFacade_Expecter{mock: &_m.Mock}
}

// CreateSyncLog provides a mock function with given fields: ctx, syncLog
func (_m *SyncLogFacade) CreateSyncLog(ctx context.Context, syncLog model.SyncLog) (model.SyncLog, error) {
	ret := _m.Called(ctx, syncLog)

	var r0 model.SyncLog
	if rf, ok := ret.Get(0).(func(context.Context, model.SyncLog) model.SyncLog); ok {
		r0 = rf(ctx, syncLog)
	} else {
		r0 = ret.Get(0).(model.SyncLog)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.SyncLog) error); ok {
		r1 = rf(ctx, syncLog)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SyncLogFacade_CreateSyncLog_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateSyncLog'
type SyncLogFacade_CreateSyncLog_Call struct {
	*mock.Call
}

// CreateSyncLog is a helper method to define mock.On call
//  - ctx context.Context
//  - syncLog model.SyncLog
func (_e *SyncLogFacade_Expecter) CreateSyncLog(ctx interface{}, syncLog interface{}) *SyncLogFacade_CreateSyncLog_Call {
	return &SyncLogFacade_CreateSyncLog_Call{Call: _e.mock.On("CreateSyncLog", ctx, syncLog)}
}

func (_c *SyncLogFacade_CreateSyncLog_Call) Run(run func(ctx context.Context, syncLog model.SyncLog)) *SyncLogFacade_CreateSyncLog_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.SyncLog))
	})
	return _c
}

func (_c *SyncLogFacade_CreateSyncLog_Call) Return(_a0 model.SyncLog, _a1 error) *SyncLogFacade_CreateSyncLog_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetLastSync provides a mock function with given fields: ctx, name
func (_m *SyncLogFacade) GetLastSync(ctx context.Context, name string) (model.SyncLog, error) {
	ret := _m.Called(ctx, name)

	var r0 model.SyncLog
	if rf, ok := ret.Get(0).(func(context.Context, string) model.SyncLog); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(model.SyncLog)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SyncLogFacade_GetLastSync_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLastSync'
type SyncLogFacade_GetLastSync_Call struct {
	*mock.Call
}

// GetLastSync is a helper method to define mock.On call
//  - ctx context.Context
//  - name string
func (_e *SyncLogFacade_Expecter) GetLastSync(ctx interface{}, name interface{}) *SyncLogFacade_GetLastSync_Call {
	return &SyncLogFacade_GetLastSync_Call{Call: _e.mock.On("GetLastSync", ctx, name)}
}

func (_c *SyncLogFacade_GetLastSync_Call) Run(run func(ctx context.Context, name string)) *SyncLogFacade_GetLastSync_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *SyncLogFacade_GetLastSync_Call) Return(_a0 model.SyncLog, _a1 error) *SyncLogFacade_GetLastSync_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// PatchSyncLog provides a mock function with given fields: ctx, syncLog
func (_m *SyncLogFacade) PatchSyncLog(ctx context.Context, syncLog model.SyncLog) (model.SyncLog, error) {
	ret := _m.Called(ctx, syncLog)

	var r0 model.SyncLog
	if rf, ok := ret.Get(0).(func(context.Context, model.SyncLog) model.SyncLog); ok {
		r0 = rf(ctx, syncLog)
	} else {
		r0 = ret.Get(0).(model.SyncLog)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.SyncLog) error); ok {
		r1 = rf(ctx, syncLog)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SyncLogFacade_PatchSyncLog_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PatchSyncLog'
type SyncLogFacade_PatchSyncLog_Call struct {
	*mock.Call
}

// PatchSyncLog is a helper method to define mock.On call
//  - ctx context.Context
//  - syncLog model.SyncLog
func (_e *SyncLogFacade_Expecter) PatchSyncLog(ctx interface{}, syncLog interface{}) *SyncLogFacade_PatchSyncLog_Call {
	return &SyncLogFacade_PatchSyncLog_Call{Call: _e.mock.On("PatchSyncLog", ctx, syncLog)}
}

func (_c *SyncLogFacade_PatchSyncLog_Call) Run(run func(ctx context.Context, syncLog model.SyncLog)) *SyncLogFacade_PatchSyncLog_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.SyncLog))
	})
	return _c
}

func (_c *SyncLogFacade_PatchSyncLog_Call) Return(_a0 model.SyncLog, _a1 error) *SyncLogFacade_PatchSyncLog_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewSyncLogFacade interface {
	mock.TestingT
	Cleanup(func())
}

// NewSyncLogFacade creates a new instance of SyncLogFacade. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSyncLogFacade(t mockConstructorTestingTNewSyncLogFacade) *SyncLogFacade {
	mock := &SyncLogFacade{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
