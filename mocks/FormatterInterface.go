// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import logger "github.com/instabledesign/logger"
import mock "github.com/stretchr/testify/mock"

// FormatterInterface is an autogenerated mock type for the FormatterInterface type
type FormatterInterface struct {
	mock.Mock
}

// Format provides a mock function with given fields: entry
func (_m *FormatterInterface) Format(entry logger.Entry) interface{} {
	ret := _m.Called(entry)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(logger.Entry) interface{}); ok {
		r0 = rf(entry)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}