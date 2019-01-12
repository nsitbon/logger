// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import logger "github.com/instabledesign/logger"
import mock "github.com/stretchr/testify/mock"

// HandlerInterface is an autogenerated mock type for the HandlerInterface type
type HandlerInterface struct {
	mock.Mock
}

// Handle provides a mock function with given fields: e
func (_m *HandlerInterface) Handle(e logger.Entry) error {
	ret := _m.Called(e)

	var r0 error
	if rf, ok := ret.Get(0).(func(logger.Entry) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
