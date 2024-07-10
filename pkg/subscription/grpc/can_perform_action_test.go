package grpc_test

import (
	"context"
	"testing"
	"time"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/subscriptionpb"
	mocksubscription "github.com/namhq1989/vocab-booster-server-app/internal/mock/subscription"
	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/grpc"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type canPerformActionTestSuite struct {
	suite.Suite
	handler     grpc.CanPerformActionHandler
	mockCtrl    *gomock.Controller
	mockService *mocksubscription.MockService
}

func (s *canPerformActionTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *canPerformActionTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockService = mocksubscription.NewMockService(s.mockCtrl)

	s.handler = grpc.NewCanPerformActionHandler(s.mockService)
}

func (s *canPerformActionTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *canPerformActionTestSuite) Test_1_Success() {
	// mock data
	var userID = database.NewStringID()

	s.mockService.EXPECT().
		GetUserSubscription(gomock.Any(), gomock.Any()).
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

func (s *canPerformActionTestSuite) Test_1_SuccessButExceedLimit() {
	// mock data
	var (
		userID = database.NewStringID()
	)

	s.mockService.EXPECT().
		GetUserSubscription(gomock.Any(), gomock.Any()).
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
		TotalPerformedToday: 50,
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), false, resp.GetCan())
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

func TestCanPerformActionTestSuite(t *testing.T) {
	suite.Run(t, new(canPerformActionTestSuite))
}
