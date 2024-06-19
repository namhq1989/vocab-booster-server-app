package grpc_test

import (
	"context"
	"testing"
	"time"

	"github.com/namhq1989/vocab-booster-server-app/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-server-app/core/error"
	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/subscriptionpb"
	mocksubscription "github.com/namhq1989/vocab-booster-server-app/internal/mock/subscription"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/grpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type canPerformActionTestSuite struct {
	suite.Suite
	handler                        grpc.CanPerformActionHandler
	mockCtrl                       *gomock.Controller
	mockUserSubscriptionRepository *mocksubscription.MockUserSubscriptionRepository
	mockCachingRepository          *mocksubscription.MockCachingRepository
}

func (s *canPerformActionTestSuite) SetupSuite() {
	s.setupApplication()
}

func (*canPerformActionTestSuite) AfterTest(_, _ string) {
	// do nothing
}

func (s *canPerformActionTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockUserSubscriptionRepository = mocksubscription.NewMockUserSubscriptionRepository(s.mockCtrl)
	s.mockCachingRepository = mocksubscription.NewMockCachingRepository(s.mockCtrl)

	s.handler = grpc.NewCanPerformActionHandler(s.mockUserSubscriptionRepository, s.mockCachingRepository)
}

func (s *canPerformActionTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *canPerformActionTestSuite) Test_1_SuccessWithDataFromDB() {
	// mock data
	var userID = database.NewStringID()

	s.mockCachingRepository.EXPECT().
		GetUserSubscriptionPlan(gomock.Any(), gomock.Any()).
		Return(nil, nil)

	s.mockUserSubscriptionRepository.EXPECT().
		FindUserSubscriptionByUserID(gomock.Any(), gomock.Any()).
		Return(&domain.UserSubscription{
			ID:      database.NewStringID(),
			UserID:  userID,
			Plan:    domain.PlanFree,
			StartAt: time.Time{},
			EndAt:   time.Time{},
		}, nil)

	// call
	ctx := appcontext.NewGRPC(context.Background())
	resp, err := s.handler.CanPerformAction(ctx, &subscriptionpb.CanPerformActionRequest{
		UserId:              userID,
		Action:              domain.ActionReviewSentence.String(),
		TotalPerformedToday: 0,
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), true, resp.GetCan())
}

func (s *canPerformActionTestSuite) Test_1_SuccessWithDataFromCaching() {
	// mock data
	var (
		userID = database.NewStringID()
		plan   = domain.PlanPremiumMonthly
	)

	s.mockCachingRepository.EXPECT().
		GetUserSubscriptionPlan(gomock.Any(), gomock.Any()).
		Return(&plan, nil)

	// call
	ctx := appcontext.NewGRPC(context.Background())
	resp, err := s.handler.CanPerformAction(ctx, &subscriptionpb.CanPerformActionRequest{
		UserId:              userID,
		Action:              domain.ActionReviewSentence.String(),
		TotalPerformedToday: 0,
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), true, resp.GetCan())
}

func (s *canPerformActionTestSuite) Test_1_SuccessButExceedLimit() {
	// mock data
	var (
		userID = database.NewStringID()
		plan   = domain.PlanFree
	)

	s.mockCachingRepository.EXPECT().
		GetUserSubscriptionPlan(gomock.Any(), gomock.Any()).
		Return(&plan, nil)

	// call
	ctx := appcontext.NewGRPC(context.Background())
	resp, err := s.handler.CanPerformAction(ctx, &subscriptionpb.CanPerformActionRequest{
		UserId:              userID,
		Action:              domain.ActionReviewSentence.String(),
		TotalPerformedToday: 50,
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), false, resp.GetCan())
}

func (s *canPerformActionTestSuite) Test_2_Fail_InvalidUserID() {
	// mock data
	s.mockCachingRepository.EXPECT().
		GetUserSubscriptionPlan(gomock.Any(), gomock.Any()).
		Return(nil, nil)

	s.mockUserSubscriptionRepository.EXPECT().
		FindUserSubscriptionByUserID(gomock.Any(), gomock.Any()).
		Return(nil, apperrors.User.InvalidUserID)

	// call
	ctx := appcontext.NewGRPC(context.Background())
	resp, err := s.handler.CanPerformAction(ctx, &subscriptionpb.CanPerformActionRequest{
		UserId:              "invalid id",
		Action:              domain.ActionReviewSentence.String(),
		TotalPerformedToday: 0,
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.User.InvalidUserID, err)
}

func (s *canPerformActionTestSuite) Test_2_Fail_InvalidAction() {
	// call
	ctx := appcontext.NewGRPC(context.Background())
	resp, err := s.handler.CanPerformAction(ctx, &subscriptionpb.CanPerformActionRequest{
		UserId:              database.NewStringID(),
		Action:              "invalid action",
		TotalPerformedToday: 0,
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.InvalidAction, err)
}

//
// END OF CASES
//

func TestcanPerformActionTestSuite(t *testing.T) {
	suite.Run(t, new(canPerformActionTestSuite))
}
