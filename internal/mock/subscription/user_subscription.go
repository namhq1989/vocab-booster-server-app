// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/subscription/domain/user_subscription.go
//
// Generated by this command:
//
//	mockgen -source=pkg/subscription/domain/user_subscription.go -destination=internal/mock/subscription/user_subscription.go -package=mocksubscription
//

// Package mocksubscription is a generated GoMock package.
package mocksubscription

import (
	reflect "reflect"

	appcontext "github.com/namhq1989/vocab-booster-server-app/core/appcontext"
	domain "github.com/namhq1989/vocab-booster-server-app/pkg/subscription/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockUserSubscriptionRepository is a mock of UserSubscriptionRepository interface.
type MockUserSubscriptionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserSubscriptionRepositoryMockRecorder
}

// MockUserSubscriptionRepositoryMockRecorder is the mock recorder for MockUserSubscriptionRepository.
type MockUserSubscriptionRepositoryMockRecorder struct {
	mock *MockUserSubscriptionRepository
}

// NewMockUserSubscriptionRepository creates a new mock instance.
func NewMockUserSubscriptionRepository(ctrl *gomock.Controller) *MockUserSubscriptionRepository {
	mock := &MockUserSubscriptionRepository{ctrl: ctrl}
	mock.recorder = &MockUserSubscriptionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserSubscriptionRepository) EXPECT() *MockUserSubscriptionRepositoryMockRecorder {
	return m.recorder
}

// FindUserSubscriptionByUserID mocks base method.
func (m *MockUserSubscriptionRepository) FindUserSubscriptionByUserID(ctx *appcontext.AppContext, userID string) (*domain.UserSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserSubscriptionByUserID", ctx, userID)
	ret0, _ := ret[0].(*domain.UserSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserSubscriptionByUserID indicates an expected call of FindUserSubscriptionByUserID.
func (mr *MockUserSubscriptionRepositoryMockRecorder) FindUserSubscriptionByUserID(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserSubscriptionByUserID", reflect.TypeOf((*MockUserSubscriptionRepository)(nil).FindUserSubscriptionByUserID), ctx, userID)
}

// UpsertUserSubscription mocks base method.
func (m *MockUserSubscriptionRepository) UpsertUserSubscription(ctx *appcontext.AppContext, subscription domain.UserSubscription) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertUserSubscription", ctx, subscription)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertUserSubscription indicates an expected call of UpsertUserSubscription.
func (mr *MockUserSubscriptionRepositoryMockRecorder) UpsertUserSubscription(ctx, subscription any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertUserSubscription", reflect.TypeOf((*MockUserSubscriptionRepository)(nil).UpsertUserSubscription), ctx, subscription)
}