// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/gamification/domain/service.go
//
// Generated by this command:
//
//	mockgen -source=pkg/gamification/domain/service.go -destination=internal/mock/gamification/service.go -package=mockgamification
//

// Package mockgamification is a generated GoMock package.
package mockgamification

import (
	reflect "reflect"

	domain "github.com/namhq1989/vocab-booster-server-app/pkg/gamification/domain"
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

// ExerciseAnswered mocks base method.
func (m *MockService) ExerciseAnswered(ctx *appcontext.AppContext, point domain.Point, completionTime domain.CompletionTime) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExerciseAnswered", ctx, point, completionTime)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExerciseAnswered indicates an expected call of ExerciseAnswered.
func (mr *MockServiceMockRecorder) ExerciseAnswered(ctx, point, completionTime any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExerciseAnswered", reflect.TypeOf((*MockService)(nil).ExerciseAnswered), ctx, point, completionTime)
}