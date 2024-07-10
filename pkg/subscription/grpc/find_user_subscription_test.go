package grpc_test

import (
	"context"
	"testing"
	"time"

	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/subscriptionpb"
	mocksubscription "github.com/namhq1989/vocab-booster-server-app/internal/mock/subscription"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/grpc"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type findUserSubscriptionTestSuite struct {
	suite.Suite
	handler                 grpc.FindUserSubscriptionHandler
	mockCtrl                *gomock.Controller
	mockUserSubscriptionHub *mocksubscription.MockUserSubscriptionHub
}

func (s *findUserSubscriptionTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *findUserSubscriptionTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockUserSubscriptionHub = mocksubscription.NewMockUserSubscriptionHub(s.mockCtrl)

	s.handler = grpc.NewFindUserSubscriptionHandler(s.mockUserSubscriptionHub)
}

func (s *findUserSubscriptionTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *findUserSubscriptionTestSuite) Test_1_Success() {
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
	resp, err := s.handler.FindUserSubscription(ctx, &subscriptionpb.FindUserSubscriptionRequest{
		UserId: userID,
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), id, resp.GetPlan().GetId())
}

func (s *findUserSubscriptionTestSuite) Test_2_Fail_NotFound() {
	// mock data
	s.mockUserSubscriptionHub.EXPECT().
		FindUserSubscriptionByUserID(gomock.Any(), gomock.Any()).
		Return(nil, nil)

	// call
	ctx := appcontext.NewGRPC(context.Background())
	resp, err := s.handler.FindUserSubscription(ctx, &subscriptionpb.FindUserSubscriptionRequest{
		UserId: database.NewStringID(),
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Subscription.SubscriptionNotFound, err)
}

//
// END OF CASES
//

func TestFindUserSubscriptionTestSuite(t *testing.T) {
	suite.Run(t, new(findUserSubscriptionTestSuite))
}
