// Copyright 2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.

// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import (
	log "github.com/chikei-development/session-manager-plugin/src/log"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// IWebSocketChannel is an autogenerated mock type for the IWebSocketChannel type
type IWebSocketChannel struct {
	mock.Mock
}

// Close provides a mock function with given fields: _a0
func (_m *IWebSocketChannel) Close(_a0 log.T) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(log.T) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetChannelToken provides a mock function with given fields:
func (_m *IWebSocketChannel) GetChannelToken() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetStreamUrl provides a mock function with given fields:
func (_m *IWebSocketChannel) GetStreamUrl() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Initialize provides a mock function with given fields: _a0, channelUrl, channelToken
func (_m *IWebSocketChannel) Initialize(_a0 log.T, channelUrl string, channelToken string) {
	_m.Called(_a0, channelUrl, channelToken)
}

// Open provides a mock function with given fields: _a0
func (_m *IWebSocketChannel) Open(_a0 log.T) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(log.T) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendMessage provides a mock function with given fields: _a0, input, inputType
func (_m *IWebSocketChannel) SendMessage(_a0 log.T, input []byte, inputType int) error {
	ret := _m.Called(_a0, input, inputType)

	var r0 error
	if rf, ok := ret.Get(0).(func(log.T, []byte, int) error); ok {
		r0 = rf(_a0, input, inputType)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetChannelToken provides a mock function with given fields: _a0
func (_m *IWebSocketChannel) SetChannelToken(_a0 string) {
	_m.Called(_a0)
}

// SetOnError provides a mock function with given fields: onErrorHandler
func (_m *IWebSocketChannel) SetOnError(onErrorHandler func(error)) {
	_m.Called(onErrorHandler)
}

// SetOnMessage provides a mock function with given fields: onMessageHandler
func (_m *IWebSocketChannel) SetOnMessage(onMessageHandler func([]byte)) {
	_m.Called(onMessageHandler)
}

// StartPings provides a mock function with given fields: _a0, pingInterval
func (_m *IWebSocketChannel) StartPings(_a0 log.T, pingInterval time.Duration) {
	_m.Called(_a0, pingInterval)
}
