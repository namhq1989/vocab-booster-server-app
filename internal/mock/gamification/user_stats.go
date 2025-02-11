// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/gamification/domain/user_stats.go
//
// Generated by this command:
//
//	mockgen -source=pkg/gamification/domain/user_stats.go -destination=internal/mock/gamification/user_stats.go -package=mockgamification
//

// Package mockgamification is a generated GoMock package.
package mockgamification

import (
	reflect "reflect"

	domain "github.com/namhq1989/vocab-booster-server-app/pkg/gamification/domain"
	appcontext "github.com/namhq1989/vocab-booster-utilities/appcontext"
	gomock "go.uber.org/mock/gomock"
)

// MockUserStatsRepository is a mock of UserStatsRepository interface.
type MockUserStatsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserStatsRepositoryMockRecorder
}

// MockUserStatsRepositoryMockRecorder is the mock recorder for MockUserStatsRepository.
type MockUserStatsRepositoryMockRecorder struct {
	mock *MockUserStatsRepository
}

// NewMockUserStatsRepository creates a new mock instance.
func NewMockUserStatsRepository(ctrl *gomock.Controller) *MockUserStatsRepository {
	mock := &MockUserStatsRepository{ctrl: ctrl}
	mock.recorder = &MockUserStatsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserStatsRepository) EXPECT() *MockUserStatsRepositoryMockRecorder {
	return m.recorder
}

// FindUserStats mocks base method.
func (m *MockUserStatsRepository) FindUserStats(ctx *appcontext.AppContext, userID string) (*domain.UserStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserStats", ctx, userID)
	ret0, _ := ret[0].(*domain.UserStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserStats indicates an expected call of FindUserStats.
func (mr *MockUserStatsRepositoryMockRecorder) FindUserStats(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserStats", reflect.TypeOf((*MockUserStatsRepository)(nil).FindUserStats), ctx, userID)
}

// IncreaseUserStats mocks base method.
func (m *MockUserStatsRepository) IncreaseUserStats(ctx *appcontext.AppContext, userID string, point int64, completionTime int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncreaseUserStats", ctx, userID, point, completionTime)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncreaseUserStats indicates an expected call of IncreaseUserStats.
func (mr *MockUserStatsRepositoryMockRecorder) IncreaseUserStats(ctx, userID, point, completionTime any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncreaseUserStats", reflect.TypeOf((*MockUserStatsRepository)(nil).IncreaseUserStats), ctx, userID, point, completionTime)
}
