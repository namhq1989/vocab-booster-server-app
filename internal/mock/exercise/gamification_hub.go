// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/exercise/domain/gamification_hub.go
//
// Generated by this command:
//
//	mockgen -source=pkg/exercise/domain/gamification_hub.go -destination=internal/mock/exercise/gamification_hub.go -package=mockexercise
//

// Package mockexercise is a generated GoMock package.
package mockexercise

import (
	reflect "reflect"
	time "time"

	domain "github.com/namhq1989/vocab-booster-server-app/pkg/exercise/domain"
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

// GetUserRecentPointsChart mocks base method.
func (m *MockGamificationHub) GetUserRecentPointsChart(ctx *appcontext.AppContext, userID string, from, to time.Time) ([]domain.UserAggregatedPoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserRecentPointsChart", ctx, userID, from, to)
	ret0, _ := ret[0].([]domain.UserAggregatedPoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserRecentPointsChart indicates an expected call of GetUserRecentPointsChart.
func (mr *MockGamificationHubMockRecorder) GetUserRecentPointsChart(ctx, userID, from, to any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserRecentPointsChart", reflect.TypeOf((*MockGamificationHub)(nil).GetUserRecentPointsChart), ctx, userID, from, to)
}