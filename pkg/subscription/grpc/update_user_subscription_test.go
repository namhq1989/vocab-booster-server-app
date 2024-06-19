package grpc_test

import (
	"context"
	"testing"

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

type updateUserSubscriptionTestSuite struct {
	suite.Suite
	handler                               grpc.UpdateUserSubscriptionHandler
	mockCtrl                              *gomock.Controller
	mockUserSubscriptionRepository        *mocksubscription.MockUserSubscriptionRepository
	mockUserSubscriptionHistoryRepository *mocksubscription.MockUserSubscriptionHistoryRepository
}

func (s *updateUserSubscriptionTestSuite) SetupSuite() {
	s.setupApplication()
}

func (*updateUserSubscriptionTestSuite) AfterTest(_, _ string) {
	// do nothing
}

func (s *updateUserSubscriptionTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockUserSubscriptionRepository = mocksubscription.NewMockUserSubscriptionRepository(s.mockCtrl)
	s.mockUserSubscriptionHistoryRepository = mocksubscription.NewMockUserSubscriptionHistoryRepository(s.mockCtrl)

	s.handler = grpc.NewUpdateUserSubscriptionHandler(s.mockUserSubscriptionRepository, s.mockUserSubscriptionHistoryRepository)
}

func (s *updateUserSubscriptionTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *updateUserSubscriptionTestSuite) Test_1_Success() {
	// mock data
	s.mockUserSubscriptionRepository.EXPECT().
		FindUserSubscriptionByUserID(gomock.Any(), gomock.Any()).
		Return(nil, nil)

	s.mockUserSubscriptionRepository.EXPECT().
		UpsertUserSubscription(gomock.Any(), gomock.Any()).
		Return(nil).
		AnyTimes()

	s.mockUserSubscriptionHistoryRepository.EXPECT().
		CreateUserSubscriptionHistory(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewGRPC(context.Background())
	resp, err := s.handler.UpdateUserSubscription(ctx, &subscriptionpb.UpdateUserSubscriptionRequest{
		UserId:    database.NewStringID(),
		PaymentId: database.NewStringID(),
		Plan:      domain.PlanPremiumMonthly.String(),
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
}

func (s *updateUserSubscriptionTestSuite) Test_2_Fail_InvalidPlan() {
	// call
	ctx := appcontext.NewGRPC(context.Background())
	resp, err := s.handler.UpdateUserSubscription(ctx, &subscriptionpb.UpdateUserSubscriptionRequest{
		UserId:    database.NewStringID(),
		PaymentId: database.NewStringID(),
		Plan:      "invalid plan",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Subscription.InvalidPlan, err)
}

func (s *updateUserSubscriptionTestSuite) Test_2_Fail_FreePlan() {
	// call
	ctx := appcontext.NewGRPC(context.Background())
	resp, err := s.handler.UpdateUserSubscription(ctx, &subscriptionpb.UpdateUserSubscriptionRequest{
		UserId:    database.NewStringID(),
		PaymentId: database.NewStringID(),
		Plan:      domain.PlanFree.String(),
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Subscription.InvalidPlan, err)
}

func (s *updateUserSubscriptionTestSuite) Test_2_Fail_InvalidUserID() {
	// mock data
	s.mockUserSubscriptionRepository.EXPECT().
		FindUserSubscriptionByUserID(gomock.Any(), gomock.Any()).
		Return(nil, apperrors.User.InvalidUserID)

	// call
	ctx := appcontext.NewGRPC(context.Background())
	resp, err := s.handler.UpdateUserSubscription(ctx, &subscriptionpb.UpdateUserSubscriptionRequest{
		UserId:    "invalid id",
		PaymentId: database.NewStringID(),
		Plan:      domain.PlanPremiumMonthly.String(),
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.User.InvalidUserID, err)
}

//
// END OF CASES
//

func TestUpdateUserSubscriptionTestSuite(t *testing.T) {
	suite.Run(t, new(updateUserSubscriptionTestSuite))
}
