// Code generated by mockery. DO NOT EDIT.

package driver

import (
	database "github.com/goravel/framework/contracts/database"
	docker "github.com/goravel/framework/contracts/testing/docker"

	driver "github.com/goravel/framework/contracts/database/driver"

	mock "github.com/stretchr/testify/mock"
)

// Driver is an autogenerated mock type for the Driver type
type Driver struct {
	mock.Mock
}

type Driver_Expecter struct {
	mock *mock.Mock
}

func (_m *Driver) EXPECT() *Driver_Expecter {
	return &Driver_Expecter{mock: &_m.Mock}
}

// Docker provides a mock function with no fields
func (_m *Driver) Docker() (docker.DatabaseDriver, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Docker")
	}

	var r0 docker.DatabaseDriver
	var r1 error
	if rf, ok := ret.Get(0).(func() (docker.DatabaseDriver, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() docker.DatabaseDriver); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(docker.DatabaseDriver)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Driver_Docker_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Docker'
type Driver_Docker_Call struct {
	*mock.Call
}

// Docker is a helper method to define mock.On call
func (_e *Driver_Expecter) Docker() *Driver_Docker_Call {
	return &Driver_Docker_Call{Call: _e.mock.On("Docker")}
}

func (_c *Driver_Docker_Call) Run(run func()) *Driver_Docker_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Driver_Docker_Call) Return(_a0 docker.DatabaseDriver, _a1 error) *Driver_Docker_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Driver_Docker_Call) RunAndReturn(run func() (docker.DatabaseDriver, error)) *Driver_Docker_Call {
	_c.Call.Return(run)
	return _c
}

// Grammar provides a mock function with no fields
func (_m *Driver) Grammar() driver.Grammar {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Grammar")
	}

	var r0 driver.Grammar
	if rf, ok := ret.Get(0).(func() driver.Grammar); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.Grammar)
		}
	}

	return r0
}

// Driver_Grammar_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Grammar'
type Driver_Grammar_Call struct {
	*mock.Call
}

// Grammar is a helper method to define mock.On call
func (_e *Driver_Expecter) Grammar() *Driver_Grammar_Call {
	return &Driver_Grammar_Call{Call: _e.mock.On("Grammar")}
}

func (_c *Driver_Grammar_Call) Run(run func()) *Driver_Grammar_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Driver_Grammar_Call) Return(_a0 driver.Grammar) *Driver_Grammar_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Driver_Grammar_Call) RunAndReturn(run func() driver.Grammar) *Driver_Grammar_Call {
	_c.Call.Return(run)
	return _c
}

// Pool provides a mock function with no fields
func (_m *Driver) Pool() database.Pool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Pool")
	}

	var r0 database.Pool
	if rf, ok := ret.Get(0).(func() database.Pool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(database.Pool)
	}

	return r0
}

// Driver_Pool_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Pool'
type Driver_Pool_Call struct {
	*mock.Call
}

// Pool is a helper method to define mock.On call
func (_e *Driver_Expecter) Pool() *Driver_Pool_Call {
	return &Driver_Pool_Call{Call: _e.mock.On("Pool")}
}

func (_c *Driver_Pool_Call) Run(run func()) *Driver_Pool_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Driver_Pool_Call) Return(_a0 database.Pool) *Driver_Pool_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Driver_Pool_Call) RunAndReturn(run func() database.Pool) *Driver_Pool_Call {
	_c.Call.Return(run)
	return _c
}

// Processor provides a mock function with no fields
func (_m *Driver) Processor() driver.Processor {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Processor")
	}

	var r0 driver.Processor
	if rf, ok := ret.Get(0).(func() driver.Processor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.Processor)
		}
	}

	return r0
}

// Driver_Processor_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Processor'
type Driver_Processor_Call struct {
	*mock.Call
}

// Processor is a helper method to define mock.On call
func (_e *Driver_Expecter) Processor() *Driver_Processor_Call {
	return &Driver_Processor_Call{Call: _e.mock.On("Processor")}
}

func (_c *Driver_Processor_Call) Run(run func()) *Driver_Processor_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Driver_Processor_Call) Return(_a0 driver.Processor) *Driver_Processor_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Driver_Processor_Call) RunAndReturn(run func() driver.Processor) *Driver_Processor_Call {
	_c.Call.Return(run)
	return _c
}

// NewDriver creates a new instance of Driver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDriver(t interface {
	mock.TestingT
	Cleanup(func())
}) *Driver {
	mock := &Driver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
