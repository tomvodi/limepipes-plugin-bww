// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	music_model "github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/musicmodel"
)

// TuneFixer is an autogenerated mock type for the TuneFixer type
type TuneFixer struct {
	mock.Mock
}

type TuneFixer_Expecter struct {
	mock *mock.Mock
}

func (_m *TuneFixer) EXPECT() *TuneFixer_Expecter {
	return &TuneFixer_Expecter{mock: &_m.Mock}
}

// Fix provides a mock function with given fields: muMo
func (_m *TuneFixer) Fix(muMo music_model.MusicModel) {
	_m.Called(muMo)
}

// TuneFixer_Fix_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Fix'
type TuneFixer_Fix_Call struct {
	*mock.Call
}

// Fix is a helper method to define mock.On call
//   - muMo music_model.MusicModel
func (_e *TuneFixer_Expecter) Fix(muMo interface{}) *TuneFixer_Fix_Call {
	return &TuneFixer_Fix_Call{Call: _e.mock.On("Fix", muMo)}
}

func (_c *TuneFixer_Fix_Call) Run(run func(muMo music_model.MusicModel)) *TuneFixer_Fix_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(music_model.MusicModel))
	})
	return _c
}

func (_c *TuneFixer_Fix_Call) Return() *TuneFixer_Fix_Call {
	_c.Call.Return()
	return _c
}

func (_c *TuneFixer_Fix_Call) RunAndReturn(run func(music_model.MusicModel)) *TuneFixer_Fix_Call {
	_c.Call.Return(run)
	return _c
}

// NewTuneFixer creates a new instance of TuneFixer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTuneFixer(t interface {
	mock.TestingT
	Cleanup(func())
}) *TuneFixer {
	mock := &TuneFixer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
