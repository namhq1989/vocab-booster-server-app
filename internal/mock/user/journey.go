// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/user/domain/journey.go
//
// Generated by this command:
//
//	mockgen -source=pkg/user/domain/journey.go -destination=internal/mock/user/journey.go -package=mockuser
//

// Package mockuser is a generated GoMock package.
package mockuser

import (
	reflect "reflect"

	domain "github.com/namhq1989/vocab-booster-server-app/pkg/user/domain"
	appcontext "github.com/namhq1989/vocab-booster-utilities/appcontext"
	gomock "go.uber.org/mock/gomock"
)

// MockJourneyRepository is a mock of JourneyRepository interface.
type MockJourneyRepository struct {
	ctrl     *gomock.Controller
	recorder *MockJourneyRepositoryMockRecorder
}

// MockJourneyRepositoryMockRecorder is the mock recorder for MockJourneyRepository.
type MockJourneyRepositoryMockRecorder struct {
	mock *MockJourneyRepository
}

// NewMockJourneyRepository creates a new mock instance.
func NewMockJourneyRepository(ctrl *gomock.Controller) *MockJourneyRepository {
	mock := &MockJourneyRepository{ctrl: ctrl}
	mock.recorder = &MockJourneyRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJourneyRepository) EXPECT() *MockJourneyRepositoryMockRecorder {
	return m.recorder
}

// CreateJourney mocks base method.
func (m *MockJourneyRepository) CreateJourney(ctx *appcontext.AppContext, journey domain.Journey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateJourney", ctx, journey)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateJourney indicates an expected call of CreateJourney.
func (mr *MockJourneyRepositoryMockRecorder) CreateJourney(ctx, journey any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJourney", reflect.TypeOf((*MockJourneyRepository)(nil).CreateJourney), ctx, journey)
}

// FindJourneysByUserID mocks base method.
func (m *MockJourneyRepository) FindJourneysByUserID(ctx *appcontext.AppContext, userID string) ([]domain.Journey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindJourneysByUserID", ctx, userID)
	ret0, _ := ret[0].([]domain.Journey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindJourneysByUserID indicates an expected call of FindJourneysByUserID.
func (mr *MockJourneyRepositoryMockRecorder) FindJourneysByUserID(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindJourneysByUserID", reflect.TypeOf((*MockJourneyRepository)(nil).FindJourneysByUserID), ctx, userID)
}

// FindUserCurrentJourney mocks base method.
func (m *MockJourneyRepository) FindUserCurrentJourney(ctx *appcontext.AppContext, userID string) (*domain.Journey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserCurrentJourney", ctx, userID)
	ret0, _ := ret[0].(*domain.Journey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserCurrentJourney indicates an expected call of FindUserCurrentJourney.
func (mr *MockJourneyRepositoryMockRecorder) FindUserCurrentJourney(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserCurrentJourney", reflect.TypeOf((*MockJourneyRepository)(nil).FindUserCurrentJourney), ctx, userID)
}

// UpdateJourney mocks base method.
func (m *MockJourneyRepository) UpdateJourney(ctx *appcontext.AppContext, journey domain.Journey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateJourney", ctx, journey)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateJourney indicates an expected call of UpdateJourney.
func (mr *MockJourneyRepositoryMockRecorder) UpdateJourney(ctx, journey any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateJourney", reflect.TypeOf((*MockJourneyRepository)(nil).UpdateJourney), ctx, journey)
}