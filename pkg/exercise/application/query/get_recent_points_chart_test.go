package query_test

import (
	"context"
	"testing"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	mockexercise "github.com/namhq1989/vocab-booster-server-app/internal/mock/exercise"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/application/query"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/timezone"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type getRecentPointsChartTestSuite struct {
	suite.Suite
	handler             query.GetRecentPointsChartHandler
	mockCtrl            *gomock.Controller
	mockGamificationHub *mockexercise.MockGamificationHub
}

func (s *getRecentPointsChartTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *getRecentPointsChartTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockGamificationHub = mockexercise.NewMockGamificationHub(s.mockCtrl)

	s.handler = query.NewGetRecentPointsChartHandler(s.mockGamificationHub)
}

func (s *getRecentPointsChartTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *getRecentPointsChartTestSuite) Test_1_Success() {
	// mock data
	s.mockGamificationHub.EXPECT().
		GetUserRecentPointsChart(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(make([]domain.UserAggregatedPoint, 0), nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.GetRecentPointsChart(ctx, database.NewStringID(), *timezone.UTC, dto.GetRecentPointsChartRequest{})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), 0, len(resp.Points))
}

//
// END OF CASES
//

func TestGetRecentPointsChartTestSuite(t *testing.T) {
	suite.Run(t, new(getRecentPointsChartTestSuite))
}
