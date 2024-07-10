// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/subscription/domain/queue.go
//
// Generated by this command:
//
//	mockgen -source=pkg/subscription/domain/queue.go -destination=internal/mock/subscription/queue.go -package=mocksubscription
//

// Package mocksubscription is a generated GoMock package.
package mocksubscription

import (
	reflect "reflect"

	domain "github.com/namhq1989/vocab-booster-server-app/pkg/subscription/domain"
	appcontext "github.com/namhq1989/vocab-booster-utilities/appcontext"
	gomock "go.uber.org/mock/gomock"
)

// MockQueueRepository is a mock of QueueRepository interface.
type MockQueueRepository struct {
	ctrl     *gomock.Controller
	recorder *MockQueueRepositoryMockRecorder
}

// MockQueueRepositoryMockRecorder is the mock recorder for MockQueueRepository.
type MockQueueRepositoryMockRecorder struct {
	mock *MockQueueRepository
}

// NewMockQueueRepository creates a new mock instance.
func NewMockQueueRepository(ctrl *gomock.Controller) *MockQueueRepository {
	mock := &MockQueueRepository{ctrl: ctrl}
	mock.recorder = &MockQueueRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueueRepository) EXPECT() *MockQueueRepositoryMockRecorder {
	return m.recorder
}

// DowngradeUserSubscription mocks base method.
func (m *MockQueueRepository) DowngradeUserSubscription(ctx *appcontext.AppContext, payload domain.QueueDowngradeUserSubscriptionPayload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DowngradeUserSubscription", ctx, payload)
	ret0, _ := ret[0].(error)
	return ret0
}

// DowngradeUserSubscription indicates an expected call of DowngradeUserSubscription.
func (mr *MockQueueRepositoryMockRecorder) DowngradeUserSubscription(ctx, payload any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DowngradeUserSubscription", reflect.TypeOf((*MockQueueRepository)(nil).DowngradeUserSubscription), ctx, payload)
}