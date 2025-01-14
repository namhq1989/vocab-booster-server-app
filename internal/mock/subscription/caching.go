// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/subscription/domain/caching.go
//
// Generated by this command:
//
//	mockgen -source=pkg/subscription/domain/caching.go -destination=internal/mock/subscription/caching.go -package=mocksubscription
//

// Package mocksubscription is a generated GoMock package.
package mocksubscription

import (
	reflect "reflect"

	domain "github.com/namhq1989/vocab-booster-server-app/pkg/subscription/domain"
	appcontext "github.com/namhq1989/vocab-booster-utilities/appcontext"
	gomock "go.uber.org/mock/gomock"
)

// MockCachingRepository is a mock of CachingRepository interface.
type MockCachingRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCachingRepositoryMockRecorder
}

// MockCachingRepositoryMockRecorder is the mock recorder for MockCachingRepository.
type MockCachingRepositoryMockRecorder struct {
	mock *MockCachingRepository
}

// NewMockCachingRepository creates a new mock instance.
func NewMockCachingRepository(ctrl *gomock.Controller) *MockCachingRepository {
	mock := &MockCachingRepository{ctrl: ctrl}
	mock.recorder = &MockCachingRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCachingRepository) EXPECT() *MockCachingRepositoryMockRecorder {
	return m.recorder
}

// GetUserSubscription mocks base method.
func (m *MockCachingRepository) GetUserSubscription(ctx *appcontext.AppContext, userID string) (*domain.UserSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserSubscription", ctx, userID)
	ret0, _ := ret[0].(*domain.UserSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserSubscription indicates an expected call of GetUserSubscription.
func (mr *MockCachingRepositoryMockRecorder) GetUserSubscription(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserSubscription", reflect.TypeOf((*MockCachingRepository)(nil).GetUserSubscription), ctx, userID)
}

// SetUserSubscription mocks base method.
func (m *MockCachingRepository) SetUserSubscription(ctx *appcontext.AppContext, userID string, plan domain.UserSubscription) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUserSubscription", ctx, userID, plan)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUserSubscription indicates an expected call of SetUserSubscription.
func (mr *MockCachingRepositoryMockRecorder) SetUserSubscription(ctx, userID, plan any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUserSubscription", reflect.TypeOf((*MockCachingRepository)(nil).SetUserSubscription), ctx, userID, plan)
}
