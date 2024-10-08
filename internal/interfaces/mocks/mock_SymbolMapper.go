// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	barline "github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/barline"

	measure "github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/measure"

	mock "github.com/stretchr/testify/mock"

	symbols "github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols"
)

// SymbolMapper is an autogenerated mock type for the SymbolMapper type
type SymbolMapper struct {
	mock.Mock
}

type SymbolMapper_Expecter struct {
	mock *mock.Mock
}

func (_m *SymbolMapper) EXPECT() *SymbolMapper_Expecter {
	return &SymbolMapper_Expecter{mock: &_m.Mock}
}

// BarlineForToken provides a mock function with given fields: token
func (_m *SymbolMapper) BarlineForToken(token string) (*barline.Barline, error) {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for BarlineForToken")
	}

	var r0 *barline.Barline
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*barline.Barline, error)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(string) *barline.Barline); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*barline.Barline)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SymbolMapper_BarlineForToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BarlineForToken'
type SymbolMapper_BarlineForToken_Call struct {
	*mock.Call
}

// BarlineForToken is a helper method to define mock.On call
//   - token string
func (_e *SymbolMapper_Expecter) BarlineForToken(token interface{}) *SymbolMapper_BarlineForToken_Call {
	return &SymbolMapper_BarlineForToken_Call{Call: _e.mock.On("BarlineForToken", token)}
}

func (_c *SymbolMapper_BarlineForToken_Call) Run(run func(token string)) *SymbolMapper_BarlineForToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *SymbolMapper_BarlineForToken_Call) Return(_a0 *barline.Barline, _a1 error) *SymbolMapper_BarlineForToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SymbolMapper_BarlineForToken_Call) RunAndReturn(run func(string) (*barline.Barline, error)) *SymbolMapper_BarlineForToken_Call {
	_c.Call.Return(run)
	return _c
}

// IsTimeSignature provides a mock function with given fields: token
func (_m *SymbolMapper) IsTimeSignature(token string) bool {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for IsTimeSignature")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SymbolMapper_IsTimeSignature_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsTimeSignature'
type SymbolMapper_IsTimeSignature_Call struct {
	*mock.Call
}

// IsTimeSignature is a helper method to define mock.On call
//   - token string
func (_e *SymbolMapper_Expecter) IsTimeSignature(token interface{}) *SymbolMapper_IsTimeSignature_Call {
	return &SymbolMapper_IsTimeSignature_Call{Call: _e.mock.On("IsTimeSignature", token)}
}

func (_c *SymbolMapper_IsTimeSignature_Call) Run(run func(token string)) *SymbolMapper_IsTimeSignature_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *SymbolMapper_IsTimeSignature_Call) Return(_a0 bool) *SymbolMapper_IsTimeSignature_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SymbolMapper_IsTimeSignature_Call) RunAndReturn(run func(string) bool) *SymbolMapper_IsTimeSignature_Call {
	_c.Call.Return(run)
	return _c
}

// SymbolForToken provides a mock function with given fields: token
func (_m *SymbolMapper) SymbolForToken(token string) (*symbols.Symbol, error) {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for SymbolForToken")
	}

	var r0 *symbols.Symbol
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*symbols.Symbol, error)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(string) *symbols.Symbol); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*symbols.Symbol)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SymbolMapper_SymbolForToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SymbolForToken'
type SymbolMapper_SymbolForToken_Call struct {
	*mock.Call
}

// SymbolForToken is a helper method to define mock.On call
//   - token string
func (_e *SymbolMapper_Expecter) SymbolForToken(token interface{}) *SymbolMapper_SymbolForToken_Call {
	return &SymbolMapper_SymbolForToken_Call{Call: _e.mock.On("SymbolForToken", token)}
}

func (_c *SymbolMapper_SymbolForToken_Call) Run(run func(token string)) *SymbolMapper_SymbolForToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *SymbolMapper_SymbolForToken_Call) Return(_a0 *symbols.Symbol, _a1 error) *SymbolMapper_SymbolForToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SymbolMapper_SymbolForToken_Call) RunAndReturn(run func(string) (*symbols.Symbol, error)) *SymbolMapper_SymbolForToken_Call {
	_c.Call.Return(run)
	return _c
}

// TimeSigForToken provides a mock function with given fields: token
func (_m *SymbolMapper) TimeSigForToken(token string) (*measure.TimeSignature, error) {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for TimeSigForToken")
	}

	var r0 *measure.TimeSignature
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*measure.TimeSignature, error)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(string) *measure.TimeSignature); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*measure.TimeSignature)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SymbolMapper_TimeSigForToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TimeSigForToken'
type SymbolMapper_TimeSigForToken_Call struct {
	*mock.Call
}

// TimeSigForToken is a helper method to define mock.On call
//   - token string
func (_e *SymbolMapper_Expecter) TimeSigForToken(token interface{}) *SymbolMapper_TimeSigForToken_Call {
	return &SymbolMapper_TimeSigForToken_Call{Call: _e.mock.On("TimeSigForToken", token)}
}

func (_c *SymbolMapper_TimeSigForToken_Call) Run(run func(token string)) *SymbolMapper_TimeSigForToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *SymbolMapper_TimeSigForToken_Call) Return(_a0 *measure.TimeSignature, _a1 error) *SymbolMapper_TimeSigForToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SymbolMapper_TimeSigForToken_Call) RunAndReturn(run func(string) (*measure.TimeSignature, error)) *SymbolMapper_TimeSigForToken_Call {
	_c.Call.Return(run)
	return _c
}

// NewSymbolMapper creates a new instance of SymbolMapper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSymbolMapper(t interface {
	mock.TestingT
	Cleanup(func())
}) *SymbolMapper {
	mock := &SymbolMapper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
