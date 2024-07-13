package query_test

import (
	"context"
	"testing"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	mockuser "github.com/namhq1989/vocab-booster-server-app/internal/mock/user"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/application/query"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type getStatsTestSuite struct {
	suite.Suite
	handler             query.GetStatsHandler
	mockCtrl            *gomock.Controller
	mockGamificationHub *mockuser.MockGamificationHub
}

func (s *getStatsTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *getStatsTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockGamificationHub = mockuser.NewMockGamificationHub(s.mockCtrl)

	s.handler = query.NewGetStatsHandler(s.mockGamificationHub)
}

func (s *getStatsTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *getStatsTestSuite) Test_1_Success() {
	// mock data
	var (
		point          int64 = 1000
		completionTime       = 5000
	)

	s.mockGamificationHub.EXPECT().
		GetUserStats(gomock.Any(), gomock.Any()).
		Return(&domain.GamificationUserStats{
			Point:          point,
			CompletionTime: completionTime,
		}, nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.GetStats(ctx, database.NewStringID(), dto.GetStatsRequest{})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), point, resp.Point)
	assert.Equal(s.T(), completionTime, resp.CompletionTime)
}

//
// END OF CASES
//

func TestGetStatsTestSuite(t *testing.T) {
	suite.Run(t, new(getStatsTestSuite))
}
