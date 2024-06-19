package grpc_test

import (
	"context"
	"testing"
	"time"

	"github.com/namhq1989/vocab-booster-server-app/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-server-app/core/error"
	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/subscriptionpb"
	mockmongo "github.com/namhq1989/vocab-booster-server-app/internal/mock/mongo"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/grpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type createUserTestSuite struct {
	suite.Suite
	handler                 grpc.CreateUserSubscriptionHandler
	mockCtrl                *gomock.Controller
	mockUserSubscriptionHub *mockmongo.MockUserSubscriptionHub
}

func (s *createUserTestSuite) SetupSuite() {
	s.setupApplication()
}

func (*createUserTestSuite) AfterTest(_, _ string) {
	// do nothing
}

func (s *createUserTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockUserSubscriptionHub = mockmongo.NewMockUserSubscriptionHub(s.mockCtrl)

	s.handler = grpc.NewCreateUserSubscriptionHandler(s.mockUserSubscriptionHub)
}

func (s *createUserTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *createUserTestSuite) Test_1_Success() {
	// mock data
	var userID = database.NewStringID()

	s.mockUserSubscriptionHub.EXPECT().
		FindUserSubscriptionByUserID(gomock.Any(), gomock.Any()).
		Return(nil, nil)

	s.mockUserSubscriptionHub.EXPECT().
		CreateUserSubscription(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewGRPC(context.Background())
	resp, err := s.handler.CreateUserSubscription(ctx, &subscriptionpb.CreateUserSubscriptionRequest{
		UserId: userID,
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
}

func (s *createUserTestSuite) Test_2_Fail_InvalidUserID() {
	// mock data
	s.mockUserSubscriptionHub.EXPECT().
		FindUserSubscriptionByUserID(gomock.Any(), gomock.Any()).
		Return(nil, apperrors.User.InvalidUserID)

	// call
	ctx := appcontext.NewGRPC(context.Background())
	resp, err := s.handler.CreateUserSubscription(ctx, &subscriptionpb.CreateUserSubscriptionRequest{
		UserId: "invalid id",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.User.InvalidUserID, err)
}

func (s *createUserTestSuite) Test_2_Fail_UserSubscriptionExisted() {
	// mock data
	var userID = database.NewStringID()

	s.mockUserSubscriptionHub.EXPECT().
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
	resp, err := s.handler.CreateUserSubscription(ctx, &subscriptionpb.CreateUserSubscriptionRequest{
		UserId: userID,
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.AlreadyExisted, err)
}

//
// END OF CASES
//

func TestCreateUserTestSuite(t *testing.T) {
	suite.Run(t, new(createUserTestSuite))
}
