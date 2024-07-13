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

type getUserSubscriptionTestSuite struct {
	suite.Suite
	handler                 grpc.GetUserSubscriptionHandler
	mockCtrl                *gomock.Controller
	mockUserSubscriptionHub *mocksubscription.MockUserSubscriptionHub
}

func (s *getUserSubscriptionTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *getUserSubscriptionTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockUserSubscriptionHub = mocksubscription.NewMockUserSubscriptionHub(s.mockCtrl)

	s.handler = grpc.NewGetUserSubscriptionHandler(s.mockUserSubscriptionHub)
}

func (s *getUserSubscriptionTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *getUserSubscriptionTestSuite) Test_1_Success() {
	// mock data
	var (
		id     = database.NewStringID()
		userID = database.NewStringID()
	)

	s.mockUserSubscriptionHub.EXPECT().
		FindUserSubscriptionByUserID(gomock.Any(), gomock.Any()).
		Return(&domain.UserSubscription{
			ID:      id,
			UserID:  userID,
			Plan:    domain.PlanFree,
			StartAt: time.Time{},
			EndAt:   time.Time{},
		}, nil)

	// call
	ctx := appcontext.NewGRPC(context.Background())
	resp, err := s.handler.GetUserSubscription(ctx, &subscriptionpb.GetUserSubscriptionRequest{
		UserId: userID,
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), id, resp.GetPlan().GetId())
}

func (s *getUserSubscriptionTestSuite) Test_2_Fail_NotFound() {
	// mock data
	s.mockUserSubscriptionHub.EXPECT().
		FindUserSubscriptionByUserID(gomock.Any(), gomock.Any()).
		Return(nil, nil)

	// call
	ctx := appcontext.NewGRPC(context.Background())
	resp, err := s.handler.GetUserSubscription(ctx, &subscriptionpb.GetUserSubscriptionRequest{
		UserId: database.NewStringID(),
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Subscription.SubscriptionNotFound, err)
}

//
// END OF CASES
//

func TestGetUserSubscriptionTestSuite(t *testing.T) {
	suite.Run(t, new(getUserSubscriptionTestSuite))
}
