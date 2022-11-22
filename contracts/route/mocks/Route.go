// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	http "github.com/goravel/framework/contracts/http"
	mock "github.com/stretchr/testify/mock"

	nethttp "net/http"

	route "github.com/goravel/framework/contracts/route"
)

// Route is an autogenerated mock type for the Route type
type Route struct {
	mock.Mock
}

// Any provides a mock function with given fields: _a0, _a1
func (_m *Route) Any(_a0 string, _a1 http.HandlerFunc) {
	_m.Called(_a0, _a1)
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *Route) Delete(_a0 string, _a1 http.HandlerFunc) {
	_m.Called(_a0, _a1)
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *Route) Get(_a0 string, _a1 http.HandlerFunc) {
	_m.Called(_a0, _a1)
}

// Group provides a mock function with given fields: _a0
func (_m *Route) Group(_a0 route.GroupFunc) {
	_m.Called(_a0)
}

// Middleware provides a mock function with given fields: _a0
func (_m *Route) Middleware(_a0 ...http.Middleware) route.Route {
	_va := make([]interface{}, len(_a0))
	for _i := range _a0 {
		_va[_i] = _a0[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 route.Route
	if rf, ok := ret.Get(0).(func(...http.Middleware) route.Route); ok {
		r0 = rf(_a0...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(route.Route)
		}
	}

	return r0
}

// Options provides a mock function with given fields: _a0, _a1
func (_m *Route) Options(_a0 string, _a1 http.HandlerFunc) {
	_m.Called(_a0, _a1)
}

// Patch provides a mock function with given fields: _a0, _a1
func (_m *Route) Patch(_a0 string, _a1 http.HandlerFunc) {
	_m.Called(_a0, _a1)
}

// Post provides a mock function with given fields: _a0, _a1
func (_m *Route) Post(_a0 string, _a1 http.HandlerFunc) {
	_m.Called(_a0, _a1)
}

// Prefix provides a mock function with given fields: addr
func (_m *Route) Prefix(addr string) route.Route {
	ret := _m.Called(addr)

	var r0 route.Route
	if rf, ok := ret.Get(0).(func(string) route.Route); ok {
		r0 = rf(addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(route.Route)
		}
	}

	return r0
}

// Put provides a mock function with given fields: _a0, _a1
func (_m *Route) Put(_a0 string, _a1 http.HandlerFunc) {
	_m.Called(_a0, _a1)
}

// Static provides a mock function with given fields: _a0, _a1
func (_m *Route) Static(_a0 string, _a1 string) {
	_m.Called(_a0, _a1)
}

// StaticFS provides a mock function with given fields: _a0, _a1
func (_m *Route) StaticFS(_a0 string, _a1 nethttp.FileSystem) {
	_m.Called(_a0, _a1)
}

// StaticFile provides a mock function with given fields: _a0, _a1
func (_m *Route) StaticFile(_a0 string, _a1 string) {
	_m.Called(_a0, _a1)
}

type NewRouteT interface {
	mock.TestingT
	Cleanup(func())
}

// NewRoute creates a new instance of Route. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRoute(t NewRouteT) *Route {
	mock := &Route{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
