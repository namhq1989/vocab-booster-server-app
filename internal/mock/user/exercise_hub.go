// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/user/domain/exercise_hub.go
//
// Generated by this command:
//
//	mockgen -source=pkg/user/domain/exercise_hub.go -destination=internal/mock/user/exercise_hub.go -package=mockuser
//

// Package mockuser is a generated GoMock package.
package mockuser

import (
	reflect "reflect"

	domain "github.com/namhq1989/vocab-booster-server-app/pkg/user/domain"
	appcontext "github.com/namhq1989/vocab-booster-utilities/appcontext"
	gomock "go.uber.org/mock/gomock"
)

// MockExerciseHub is a mock of ExerciseHub interface.
type MockExerciseHub struct {
	ctrl     *gomock.Controller
	recorder *MockExerciseHubMockRecorder
}

// MockExerciseHubMockRecorder is the mock recorder for MockExerciseHub.
type MockExerciseHubMockRecorder struct {
	mock *MockExerciseHub
}

// NewMockExerciseHub creates a new mock instance.
func NewMockExerciseHub(ctrl *gomock.Controller) *MockExerciseHub {
	mock := &MockExerciseHub{ctrl: ctrl}
	mock.recorder = &MockExerciseHubMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExerciseHub) EXPECT() *MockExerciseHubMockRecorder {
	return m.recorder
}

// GetUserStats mocks base method.
func (m *MockExerciseHub) GetUserStats(ctx *appcontext.AppContext, userID string) (*domain.ExerciseUserStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserStats", ctx, userID)
	ret0, _ := ret[0].(*domain.ExerciseUserStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserStats indicates an expected call of GetUserStats.
func (mr *MockExerciseHubMockRecorder) GetUserStats(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserStats", reflect.TypeOf((*MockExerciseHub)(nil).GetUserStats), ctx, userID)
}