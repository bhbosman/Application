// Code generated by MockGen. DO NOT EDIT.
// Source: IInterConnector.go

// Package PubSub is a generated GoMock package.
package PubSub

import (
	Messages "github.com/bhbosman/Application/Messages"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIInterConnector is a mock of IInterConnector interface
type MockIInterConnector struct {
	ctrl     *gomock.Controller
	recorder *MockIInterConnectorMockRecorder
}

// MockIInterConnectorMockRecorder is the mock recorder for MockIInterConnector
type MockIInterConnectorMockRecorder struct {
	mock *MockIInterConnector
}

// NewMockIInterConnector creates a new mock instance
func NewMockIInterConnector(ctrl *gomock.Controller) *MockIInterConnector {
	mock := &MockIInterConnector{ctrl: ctrl}
	mock.recorder = &MockIInterConnectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIInterConnector) EXPECT() *MockIInterConnectorMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockIInterConnector) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockIInterConnectorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIInterConnector)(nil).Close))
}

// Key mocks base method
func (m *MockIInterConnector) Key() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Key")
	ret0, _ := ret[0].(string)
	return ret0
}

// Key indicates an expected call of Key
func (mr *MockIInterConnectorMockRecorder) Key() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Key", reflect.TypeOf((*MockIInterConnector)(nil).Key))
}

// receiver mocks base method
func (m *MockIInterConnector) receiver() ISubKeyBucketReceiver {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "receiver")
	ret0, _ := ret[0].(ISubKeyBucketReceiver)
	return ret0
}

// receiver indicates an expected call of receiver
func (mr *MockIInterConnectorMockRecorder) receiver() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "receiver", reflect.TypeOf((*MockIInterConnector)(nil).receiver))
}

// receiverDescription mocks base method
func (m *MockIInterConnector) receiverDescription() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "receiverDescription")
	ret0, _ := ret[0].(string)
	return ret0
}

// receiverDescription indicates an expected call of receiverDescription
func (mr *MockIInterConnectorMockRecorder) receiverDescription() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "receiverDescription", reflect.TypeOf((*MockIInterConnector)(nil).receiverDescription))
}

// MockISubKeyBucketReceiver is a mock of ISubKeyBucketReceiver interface
type MockISubKeyBucketReceiver struct {
	ctrl     *gomock.Controller
	recorder *MockISubKeyBucketReceiverMockRecorder
}

// MockISubKeyBucketReceiverMockRecorder is the mock recorder for MockISubKeyBucketReceiver
type MockISubKeyBucketReceiverMockRecorder struct {
	mock *MockISubKeyBucketReceiver
}

// NewMockISubKeyBucketReceiver creates a new mock instance
func NewMockISubKeyBucketReceiver(ctrl *gomock.Controller) *MockISubKeyBucketReceiver {
	mock := &MockISubKeyBucketReceiver{ctrl: ctrl}
	mock.recorder = &MockISubKeyBucketReceiverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockISubKeyBucketReceiver) EXPECT() *MockISubKeyBucketReceiverMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockISubKeyBucketReceiver) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockISubKeyBucketReceiverMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockISubKeyBucketReceiver)(nil).Close))
}

// Handle mocks base method
func (m *MockISubKeyBucketReceiver) Handle(waitGroup Messages.IWaitGroup, data interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle", waitGroup, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Handle indicates an expected call of Handle
func (mr *MockISubKeyBucketReceiverMockRecorder) Handle(waitGroup, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockISubKeyBucketReceiver)(nil).Handle), waitGroup, data)
}
