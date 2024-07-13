// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/user/domain/gamification.go
//
// Generated by this command:
//
//	mockgen -source=pkg/user/domain/gamification.go -destination=internal/mock/user/gamification.go -package=mockuser
//

// Package mockuser is a generated GoMock package.
package mockuser

import (
	reflect "reflect"

	domain "github.com/namhq1989/vocab-booster-server-app/pkg/user/domain"
	appcontext "github.com/namhq1989/vocab-booster-utilities/appcontext"
	gomock "go.uber.org/mock/gomock"
)

// MockGamificationHub is a mock of GamificationHub interface.
type MockGamificationHub struct {
	ctrl     *gomock.Controller
	recorder *MockGamificationHubMockRecorder
}

// MockGamificationHubMockRecorder is the mock recorder for MockGamificationHub.
type MockGamificationHubMockRecorder struct {
	mock *MockGamificationHub
}

// NewMockGamificationHub creates a new mock instance.
func NewMockGamificationHub(ctrl *gomock.Controller) *MockGamificationHub {
	mock := &MockGamificationHub{ctrl: ctrl}
	mock.recorder = &MockGamificationHubMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGamificationHub) EXPECT() *MockGamificationHubMockRecorder {
	return m.recorder
}

// GetUserStats mocks base method.
func (m *MockGamificationHub) GetUserStats(ctx *appcontext.AppContext, userID string) (*domain.GamificationUserStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserStats", ctx, userID)
	ret0, _ := ret[0].(*domain.GamificationUserStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserStats indicates an expected call of GetUserStats.
func (mr *MockGamificationHubMockRecorder) GetUserStats(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserStats", reflect.TypeOf((*MockGamificationHub)(nil).GetUserStats), ctx, userID)
}
