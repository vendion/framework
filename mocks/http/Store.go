// Code generated by mockery. DO NOT EDIT.

package http

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

type Store_Expecter struct {
	mock *mock.Mock
}

func (_m *Store) EXPECT() *Store_Expecter {
	return &Store_Expecter{mock: &_m.Mock}
}

// Burst provides a mock function with given fields: ctx, key, tokens
func (_m *Store) Burst(ctx context.Context, key string, tokens uint64) error {
	ret := _m.Called(ctx, key, tokens)

	if len(ret) == 0 {
		panic("no return value specified for Burst")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, uint64) error); ok {
		r0 = rf(ctx, key, tokens)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Store_Burst_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Burst'
type Store_Burst_Call struct {
	*mock.Call
}

// Burst is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - tokens uint64
func (_e *Store_Expecter) Burst(ctx interface{}, key interface{}, tokens interface{}) *Store_Burst_Call {
	return &Store_Burst_Call{Call: _e.mock.On("Burst", ctx, key, tokens)}
}

func (_c *Store_Burst_Call) Run(run func(ctx context.Context, key string, tokens uint64)) *Store_Burst_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(uint64))
	})
	return _c
}

func (_c *Store_Burst_Call) Return(_a0 error) *Store_Burst_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Store_Burst_Call) RunAndReturn(run func(context.Context, string, uint64) error) *Store_Burst_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, key
func (_m *Store) Get(ctx context.Context, key string) (uint64, uint64, error) {
	ret := _m.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 uint64
	var r1 uint64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (uint64, uint64, error)); ok {
		return rf(ctx, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) uint64); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) uint64); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, key)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Store_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type Store_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
func (_e *Store_Expecter) Get(ctx interface{}, key interface{}) *Store_Get_Call {
	return &Store_Get_Call{Call: _e.mock.On("Get", ctx, key)}
}

func (_c *Store_Get_Call) Run(run func(ctx context.Context, key string)) *Store_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Store_Get_Call) Return(tokens uint64, remaining uint64, err error) *Store_Get_Call {
	_c.Call.Return(tokens, remaining, err)
	return _c
}

func (_c *Store_Get_Call) RunAndReturn(run func(context.Context, string) (uint64, uint64, error)) *Store_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Set provides a mock function with given fields: ctx, key, tokens, interval
func (_m *Store) Set(ctx context.Context, key string, tokens uint64, interval time.Duration) error {
	ret := _m.Called(ctx, key, tokens, interval)

	if len(ret) == 0 {
		panic("no return value specified for Set")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, uint64, time.Duration) error); ok {
		r0 = rf(ctx, key, tokens, interval)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Store_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type Store_Set_Call struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - tokens uint64
//   - interval time.Duration
func (_e *Store_Expecter) Set(ctx interface{}, key interface{}, tokens interface{}, interval interface{}) *Store_Set_Call {
	return &Store_Set_Call{Call: _e.mock.On("Set", ctx, key, tokens, interval)}
}

func (_c *Store_Set_Call) Run(run func(ctx context.Context, key string, tokens uint64, interval time.Duration)) *Store_Set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(uint64), args[3].(time.Duration))
	})
	return _c
}

func (_c *Store_Set_Call) Return(_a0 error) *Store_Set_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Store_Set_Call) RunAndReturn(run func(context.Context, string, uint64, time.Duration) error) *Store_Set_Call {
	_c.Call.Return(run)
	return _c
}

// Take provides a mock function with given fields: ctx, key
func (_m *Store) Take(ctx context.Context, key string) (uint64, uint64, uint64, bool, error) {
	ret := _m.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for Take")
	}

	var r0 uint64
	var r1 uint64
	var r2 uint64
	var r3 bool
	var r4 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (uint64, uint64, uint64, bool, error)); ok {
		return rf(ctx, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) uint64); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) uint64); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string) uint64); ok {
		r2 = rf(ctx, key)
	} else {
		r2 = ret.Get(2).(uint64)
	}

	if rf, ok := ret.Get(3).(func(context.Context, string) bool); ok {
		r3 = rf(ctx, key)
	} else {
		r3 = ret.Get(3).(bool)
	}

	if rf, ok := ret.Get(4).(func(context.Context, string) error); ok {
		r4 = rf(ctx, key)
	} else {
		r4 = ret.Error(4)
	}

	return r0, r1, r2, r3, r4
}

// Store_Take_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Take'
type Store_Take_Call struct {
	*mock.Call
}

// Take is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
func (_e *Store_Expecter) Take(ctx interface{}, key interface{}) *Store_Take_Call {
	return &Store_Take_Call{Call: _e.mock.On("Take", ctx, key)}
}

func (_c *Store_Take_Call) Run(run func(ctx context.Context, key string)) *Store_Take_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Store_Take_Call) Return(tokens uint64, remaining uint64, reset uint64, ok bool, err error) *Store_Take_Call {
	_c.Call.Return(tokens, remaining, reset, ok, err)
	return _c
}

func (_c *Store_Take_Call) RunAndReturn(run func(context.Context, string) (uint64, uint64, uint64, bool, error)) *Store_Take_Call {
	_c.Call.Return(run)
	return _c
}

// NewStore creates a new instance of Store. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *Store {
	mock := &Store{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
