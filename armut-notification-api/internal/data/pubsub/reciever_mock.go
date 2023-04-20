// Code generated by MockGen. DO NOT EDIT.
// Source: ../../internal/data/pubsub/reciever.go

// Package pubsub is a generated GoMock package.
package pubsub

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIReceiver is a mock of IReceiver interface.
type MockIReceiver struct {
	ctrl     *gomock.Controller
	recorder *MockIReceiverMockRecorder
}

// MockIReceiverMockRecorder is the mock recorder for MockIReceiver.
type MockIReceiverMockRecorder struct {
	mock *MockIReceiver
}

// NewMockIReceiver creates a new mock instance.
func NewMockIReceiver(ctrl *gomock.Controller) *MockIReceiver {
	mock := &MockIReceiver{ctrl: ctrl}
	mock.recorder = &MockIReceiverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIReceiver) EXPECT() *MockIReceiverMockRecorder {
	return m.recorder
}

// InitReceivers mocks base method.
func (m *MockIReceiver) InitReceivers(count int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InitReceivers", count)
}

// InitReceivers indicates an expected call of InitReceivers.
func (mr *MockIReceiverMockRecorder) InitReceivers(count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitReceivers", reflect.TypeOf((*MockIReceiver)(nil).InitReceivers), count)
}

// MockIReceiverHandler is a mock of IReceiverHandler interface.
type MockIReceiverHandler struct {
	ctrl     *gomock.Controller
	recorder *MockIReceiverHandlerMockRecorder
}

// MockIReceiverHandlerMockRecorder is the mock recorder for MockIReceiverHandler.
type MockIReceiverHandlerMockRecorder struct {
	mock *MockIReceiverHandler
}

// NewMockIReceiverHandler creates a new mock instance.
func NewMockIReceiverHandler(ctrl *gomock.Controller) *MockIReceiverHandler {
	mock := &MockIReceiverHandler{ctrl: ctrl}
	mock.recorder = &MockIReceiverHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIReceiverHandler) EXPECT() *MockIReceiverHandlerMockRecorder {
	return m.recorder
}

// Handle mocks base method.
func (m *MockIReceiverHandler) Handle(ch chan error, model *ReceiverHandlerModel) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Handle", ch, model)
}

// Handle indicates an expected call of Handle.
func (mr *MockIReceiverHandlerMockRecorder) Handle(ch, model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockIReceiverHandler)(nil).Handle), ch, model)
}
