// Code generated by MockGen. DO NOT EDIT.
// Source: IMessageFactory.go

// Package Messages is a generated GoMock package.
package Messages

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIMessageFactory is a mock of IMessageFactory interface
type MockIMessageFactory struct {
	ctrl     *gomock.Controller
	recorder *MockIMessageFactoryMockRecorder
}

// MockIMessageFactoryMockRecorder is the mock recorder for MockIMessageFactory
type MockIMessageFactoryMockRecorder struct {
	mock *MockIMessageFactory
}

// NewMockIMessageFactory creates a new mock instance
func NewMockIMessageFactory(ctrl *gomock.Controller) *MockIMessageFactory {
	mock := &MockIMessageFactory{ctrl: ctrl}
	mock.recorder = &MockIMessageFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIMessageFactory) EXPECT() *MockIMessageFactoryMockRecorder {
	return m.recorder
}

// Message mocks base method
func (m *MockIMessageFactory) Message() (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Message")
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Message indicates an expected call of Message
func (mr *MockIMessageFactoryMockRecorder) Message() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Message", reflect.TypeOf((*MockIMessageFactory)(nil).Message))
}

// MessageType mocks base method
func (m *MockIMessageFactory) MessageType() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MessageType")
	ret0, _ := ret[0].(int)
	return ret0
}

// MessageType indicates an expected call of MessageType
func (mr *MockIMessageFactoryMockRecorder) MessageType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MessageType", reflect.TypeOf((*MockIMessageFactory)(nil).MessageType))
}

// Length mocks base method
func (m *MockIMessageFactory) Length() uint16 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(uint16)
	return ret0
}

// Length indicates an expected call of Length
func (mr *MockIMessageFactoryMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockIMessageFactory)(nil).Length))
}
