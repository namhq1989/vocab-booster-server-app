package grpc_test

import (
	"context"
	"testing"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/gamificationpb"
	mockgamification "github.com/namhq1989/vocab-booster-server-app/internal/mock/gamification"
	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
	"github.com/namhq1989/vocab-booster-server-app/pkg/gamification/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/gamification/grpc"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type getUserStatsTestSuite struct {
	suite.Suite
	handler                 grpc.GetUserStatsHandler
	mockCtrl                *gomock.Controller
	mockUserStatsRepository *mockgamification.MockUserStatsRepository
}

func (s *getUserStatsTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *getUserStatsTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockUserStatsRepository = mockgamification.NewMockUserStatsRepository(s.mockCtrl)

	s.handler = grpc.NewGetUserStatsHandler(s.mockUserStatsRepository)
}

func (s *getUserStatsTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *getUserStatsTestSuite) Test_1_Success() {
	// mock data
	var (
		userID               = database.NewStringID()
		point          int64 = 500
		completionTime       = 2451
	)

	s.mockUserStatsRepository.EXPECT().
		FindUserStats(gomock.Any(), gomock.Any()).
		Return(&domain.UserStats{
			ID:             database.NewStringID(),
			UserID:         userID,
			Point:          point,
			CompletionTime: completionTime,
		}, nil)

	// call
	ctx := appcontext.NewGRPC(context.Background())
	resp, err := s.handler.GetUserStats(ctx, &gamificationpb.GetUserStatsRequest{
		UserId: userID,
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), point, resp.GetPoint())
	assert.Equal(s.T(), completionTime, int(resp.GetCompletionTime()))
}

func (s *getUserStatsTestSuite) Test_2_Fail_InvalidUserID() {
	// mock data
	s.mockUserStatsRepository.EXPECT().
		FindUserStats(gomock.Any(), gomock.Any()).
		Return(nil, apperrors.User.InvalidUserID)

	// call
	ctx := appcontext.NewGRPC(context.Background())
	resp, err := s.handler.GetUserStats(ctx, &gamificationpb.GetUserStatsRequest{
		UserId: "invalid id",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.User.InvalidUserID, err)
}

//
// END OF CASES
//

func TestGetUserStatsTestSuite(t *testing.T) {
	suite.Run(t, new(getUserStatsTestSuite))
}
