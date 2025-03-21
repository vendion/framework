// Code generated by mockery. DO NOT EDIT.

package log

import (
	log "github.com/goravel/framework/contracts/log"
	mock "github.com/stretchr/testify/mock"
)

// Hook is an autogenerated mock type for the Hook type
type Hook struct {
	mock.Mock
}

type Hook_Expecter struct {
	mock *mock.Mock
}

func (_m *Hook) EXPECT() *Hook_Expecter {
	return &Hook_Expecter{mock: &_m.Mock}
}

// Fire provides a mock function with given fields: _a0
func (_m *Hook) Fire(_a0 log.Entry) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Fire")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(log.Entry) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Hook_Fire_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Fire'
type Hook_Fire_Call struct {
	*mock.Call
}

// Fire is a helper method to define mock.On call
//   - _a0 log.Entry
func (_e *Hook_Expecter) Fire(_a0 interface{}) *Hook_Fire_Call {
	return &Hook_Fire_Call{Call: _e.mock.On("Fire", _a0)}
}

func (_c *Hook_Fire_Call) Run(run func(_a0 log.Entry)) *Hook_Fire_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(log.Entry))
	})
	return _c
}

func (_c *Hook_Fire_Call) Return(_a0 error) *Hook_Fire_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Hook_Fire_Call) RunAndReturn(run func(log.Entry) error) *Hook_Fire_Call {
	_c.Call.Return(run)
	return _c
}

// Levels provides a mock function with no fields
func (_m *Hook) Levels() []log.Level {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Levels")
	}

	var r0 []log.Level
	if rf, ok := ret.Get(0).(func() []log.Level); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]log.Level)
		}
	}

	return r0
}

// Hook_Levels_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Levels'
type Hook_Levels_Call struct {
	*mock.Call
}

// Levels is a helper method to define mock.On call
func (_e *Hook_Expecter) Levels() *Hook_Levels_Call {
	return &Hook_Levels_Call{Call: _e.mock.On("Levels")}
}

func (_c *Hook_Levels_Call) Run(run func()) *Hook_Levels_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Hook_Levels_Call) Return(_a0 []log.Level) *Hook_Levels_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Hook_Levels_Call) RunAndReturn(run func() []log.Level) *Hook_Levels_Call {
	_c.Call.Return(run)
	return _c
}

// NewHook creates a new instance of Hook. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHook(t interface {
	mock.TestingT
	Cleanup(func())
}) *Hook {
	mock := &Hook{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
