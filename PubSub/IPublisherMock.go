// Code generated by MockGen. DO NOT EDIT.
// Source: IPublisher.go

// Package PubSub is a generated GoMock package.
package PubSub

import (
	Messages "github.com/bhbosman/Application/Messages"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIPublisher is a mock of IPublisher interface
type MockIPublisher struct {
	ctrl     *gomock.Controller
	recorder *MockIPublisherMockRecorder
}

// MockIPublisherMockRecorder is the mock recorder for MockIPublisher
type MockIPublisherMockRecorder struct {
	mock *MockIPublisher
}

// NewMockIPublisher creates a new mock instance
func NewMockIPublisher(ctrl *gomock.Controller) *MockIPublisher {
	mock := &MockIPublisher{ctrl: ctrl}
	mock.recorder = &MockIPublisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIPublisher) EXPECT() *MockIPublisherMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockIPublisher) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockIPublisherMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIPublisher)(nil).Close))
}

// Publish mocks base method
func (m *MockIPublisher) Publish(key, subKey string, waitGroup Messages.IWaitGroup, data interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", key, subKey, waitGroup, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish
func (mr *MockIPublisherMockRecorder) Publish(key, subKey, waitGroup, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockIPublisher)(nil).Publish), key, subKey, waitGroup, data)
}

// Register mocks base method
func (m *MockIPublisher) Register(key, subKey string, receiver ISubKeyBucketReceiver) (IInterConnector, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", key, subKey, receiver)
	ret0, _ := ret[0].(IInterConnector)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register
func (mr *MockIPublisherMockRecorder) Register(key, subKey, receiver interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockIPublisher)(nil).Register), key, subKey, receiver)
}

// UnRegister mocks base method
func (m *MockIPublisher) UnRegister(key, subKey, receiverKey string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnRegister", key, subKey, receiverKey)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnRegister indicates an expected call of UnRegister
func (mr *MockIPublisherMockRecorder) UnRegister(key, subKey, receiverKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnRegister", reflect.TypeOf((*MockIPublisher)(nil).UnRegister), key, subKey, receiverKey)
}

// UnRegisterReceiver mocks base method
func (m *MockIPublisher) UnRegisterReceiver(receiver ISubKeyBucketReceiver) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UnRegisterReceiver", receiver)
}

// UnRegisterReceiver indicates an expected call of UnRegisterReceiver
func (mr *MockIPublisherMockRecorder) UnRegisterReceiver(receiver interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnRegisterReceiver", reflect.TypeOf((*MockIPublisher)(nil).UnRegisterReceiver), receiver)
}