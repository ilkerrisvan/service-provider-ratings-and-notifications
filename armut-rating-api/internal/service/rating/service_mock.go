// Code generated by MockGen. DO NOT EDIT.
// Source: ../../internal/service/rating/service.go

// Package rating is a generated GoMock package.
package rating

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIRatingService is a mock of IRatingService interface.
type MockIRatingService struct {
	ctrl     *gomock.Controller
	recorder *MockIRatingServiceMockRecorder
}

// MockIRatingServiceMockRecorder is the mock recorder for MockIRatingService.
type MockIRatingServiceMockRecorder struct {
	mock *MockIRatingService
}

// NewMockIRatingService creates a new mock instance.
func NewMockIRatingService(ctrl *gomock.Controller) *MockIRatingService {
	mock := &MockIRatingService{ctrl: ctrl}
	mock.recorder = &MockIRatingServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRatingService) EXPECT() *MockIRatingServiceMockRecorder {
	return m.recorder
}

// AddRating mocks base method.
func (m *MockIRatingService) AddRating(ch chan *AddRatingServiceResponse, model *AddRatingServiceModel) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddRating", ch, model)
}

// AddRating indicates an expected call of AddRating.
func (mr *MockIRatingServiceMockRecorder) AddRating(ch, model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRating", reflect.TypeOf((*MockIRatingService)(nil).AddRating), ch, model)
}

// GetAverageService mocks base method.
func (m *MockIRatingService) GetAverageService(ch chan *GetAverageServiceResponse, model *GetAverageServiceRequestModel) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetAverageService", ch, model)
}

// GetAverageService indicates an expected call of GetAverageService.
func (mr *MockIRatingServiceMockRecorder) GetAverageService(ch, model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAverageService", reflect.TypeOf((*MockIRatingService)(nil).GetAverageService), ch, model)
}

// PublishPubSubMessage mocks base method.
func (m *MockIRatingService) PublishPubSubMessage(ch chan *PublishPubSubMessageServiceResponse, model *PublishPubSubMessageServiceModel) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PublishPubSubMessage", ch, model)
}

// PublishPubSubMessage indicates an expected call of PublishPubSubMessage.
func (mr *MockIRatingServiceMockRecorder) PublishPubSubMessage(ch, model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishPubSubMessage", reflect.TypeOf((*MockIRatingService)(nil).PublishPubSubMessage), ch, model)
}
