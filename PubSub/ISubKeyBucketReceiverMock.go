// Code generated by MockGen. DO NOT EDIT.
// Source: ISubKeyBucketReceiver.go

// Package PubSub is a generated GoMock package.
package PubSub

import (
	Messages "github.com/bhbosman/Application/Messages"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

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

// FeedStopped mocks base method
func (m *MockISubKeyBucketReceiver) FeedStopped() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FeedStopped")
	ret0, _ := ret[0].(error)
	return ret0
}

// FeedStopped indicates an expected call of FeedStopped
func (mr *MockISubKeyBucketReceiverMockRecorder) FeedStopped() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FeedStopped", reflect.TypeOf((*MockISubKeyBucketReceiver)(nil).FeedStopped))
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