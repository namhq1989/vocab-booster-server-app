// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/subscription/domain/service.go
//
// Generated by this command:
//
//	mockgen -source=pkg/subscription/domain/service.go -destination=internal/mock/subscription/service.go -package=mocksubscription
//

// Package mocksubscription is a generated GoMock package.
package mocksubscription

import (
	reflect "reflect"

	domain "github.com/namhq1989/vocab-booster-server-app/pkg/subscription/domain"
	appcontext "github.com/namhq1989/vocab-booster-utilities/appcontext"
	gomock "go.uber.org/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// GetUserSubscription mocks base method.
func (m *MockService) GetUserSubscription(ctx *appcontext.AppContext, userID string) (*domain.UserSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserSubscription", ctx, userID)
	ret0, _ := ret[0].(*domain.UserSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserSubscription indicates an expected call of GetUserSubscription.
func (mr *MockServiceMockRecorder) GetUserSubscription(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserSubscription", reflect.TypeOf((*MockService)(nil).GetUserSubscription), ctx, userID)
}
