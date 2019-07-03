// Code generated by MockGen. DO NOT EDIT.
// Source: IInstrumentOrderBook.go

// Package FullMarketDepth is a generated GoMock package.
package FullMarketDepth

import (
	Managers "github.com/bhbosman/Application/Managers"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIInstrumentOrderBook is a mock of IInstrumentOrderBook interface
type MockIInstrumentOrderBook struct {
	ctrl     *gomock.Controller
	recorder *MockIInstrumentOrderBookMockRecorder
}

// MockIInstrumentOrderBookMockRecorder is the mock recorder for MockIInstrumentOrderBook
type MockIInstrumentOrderBookMockRecorder struct {
	mock *MockIInstrumentOrderBook
}

// NewMockIInstrumentOrderBook creates a new mock instance
func NewMockIInstrumentOrderBook(ctrl *gomock.Controller) *MockIInstrumentOrderBook {
	mock := &MockIInstrumentOrderBook{ctrl: ctrl}
	mock.recorder = &MockIInstrumentOrderBookMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIInstrumentOrderBook) EXPECT() *MockIInstrumentOrderBookMockRecorder {
	return m.recorder
}

// Push mocks base method
func (m *MockIInstrumentOrderBook) Push(message interface{}, messageSource Managers.IMessageSource) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Push", message, messageSource)
	ret0, _ := ret[0].(error)
	return ret0
}

// Push indicates an expected call of Push
func (mr *MockIInstrumentOrderBookMockRecorder) Push(message, messageSource interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockIInstrumentOrderBook)(nil).Push), message, messageSource)
}