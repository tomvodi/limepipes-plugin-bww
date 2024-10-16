// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	symbols "github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/symbols"
)

// SymbolMerger is an autogenerated mock type for the SymbolMerger type
type SymbolMerger struct {
	mock.Mock
}

type SymbolMerger_Expecter struct {
	mock *mock.Mock
}

func (_m *SymbolMerger) EXPECT() *SymbolMerger_Expecter {
	return &SymbolMerger_Expecter{mock: &_m.Mock}
}

// MergeSymbols provides a mock function with given fields: left, right
func (_m *SymbolMerger) MergeSymbols(left *symbols.Symbol, right *symbols.Symbol) bool {
	ret := _m.Called(left, right)

	if len(ret) == 0 {
		panic("no return value specified for MergeSymbols")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(*symbols.Symbol, *symbols.Symbol) bool); ok {
		r0 = rf(left, right)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SymbolMerger_MergeSymbols_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MergeSymbols'
type SymbolMerger_MergeSymbols_Call struct {
	*mock.Call
}

// MergeSymbols is a helper method to define mock.On call
//   - left *symbols.Symbol
//   - right *symbols.Symbol
func (_e *SymbolMerger_Expecter) MergeSymbols(left interface{}, right interface{}) *SymbolMerger_MergeSymbols_Call {
	return &SymbolMerger_MergeSymbols_Call{Call: _e.mock.On("MergeSymbols", left, right)}
}

func (_c *SymbolMerger_MergeSymbols_Call) Run(run func(left *symbols.Symbol, right *symbols.Symbol)) *SymbolMerger_MergeSymbols_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*symbols.Symbol), args[1].(*symbols.Symbol))
	})
	return _c
}

func (_c *SymbolMerger_MergeSymbols_Call) Return(_a0 bool) *SymbolMerger_MergeSymbols_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SymbolMerger_MergeSymbols_Call) RunAndReturn(run func(*symbols.Symbol, *symbols.Symbol) bool) *SymbolMerger_MergeSymbols_Call {
	_c.Call.Return(run)
	return _c
}

// NewSymbolMerger creates a new instance of SymbolMerger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSymbolMerger(t interface {
	mock.TestingT
	Cleanup(func())
}) *SymbolMerger {
	mock := &SymbolMerger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
