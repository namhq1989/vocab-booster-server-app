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

type getUserRecentPointsChartTestSuite struct {
	suite.Suite
	handler             grpc.GetUserRecentPointsChartHandler
	mockCtrl            *gomock.Controller
	mockPointRepository *mockgamification.MockPointRepository
}

func (s *getUserRecentPointsChartTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *getUserRecentPointsChartTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockPointRepository = mockgamification.NewMockPointRepository(s.mockCtrl)

	s.handler = grpc.NewGetUserRecentPointsChartHandler(s.mockPointRepository)
}

func (s *getUserRecentPointsChartTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *getUserRecentPointsChartTestSuite) Test_1_Success() {
	// mock data
	s.mockPointRepository.EXPECT().
		AggregateUserPointsInTimeRange(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(make([]domain.UserAggregatedPoint, 0), nil)

	// call
	ctx := appcontext.NewGRPC(context.Background())
	resp, err := s.handler.GetUserRecentPointsChart(ctx, &gamificationpb.GetUserRecentPointsChartRequest{
		UserId: database.NewStringID(),
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), 0, len(resp.GetPoints()))
}

func (s *getUserRecentPointsChartTestSuite) Test_2_Fail_InvalidUserID() {
	// mock data
	s.mockPointRepository.EXPECT().
		AggregateUserPointsInTimeRange(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil, apperrors.User.InvalidUserID)

	// call
	ctx := appcontext.NewGRPC(context.Background())
	resp, err := s.handler.GetUserRecentPointsChart(ctx, &gamificationpb.GetUserRecentPointsChartRequest{
		UserId: "invalid id",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.User.InvalidUserID, err)
}

//
// END OF CASES
//

func TestGetUserRecentPointsChartTestSuite(t *testing.T) {
	suite.Run(t, new(getUserRecentPointsChartTestSuite))
}
