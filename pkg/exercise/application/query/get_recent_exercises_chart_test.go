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

type getRecentExercisesChartTestSuite struct {
	suite.Suite
	handler         query.GetRecentExercisesChartHandler
	mockCtrl        *gomock.Controller
	mockExerciseHub *mockexercise.MockExerciseHub
}

func (s *getRecentExercisesChartTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *getRecentExercisesChartTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockExerciseHub = mockexercise.NewMockExerciseHub(s.mockCtrl)

	s.handler = query.NewGetRecentExercisesChartHandler(s.mockExerciseHub)
}

func (s *getRecentExercisesChartTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *getRecentExercisesChartTestSuite) Test_1_Success() {
	// mock data
	s.mockExerciseHub.EXPECT().
		AggregateUserExercisesInTimeRange(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(make([]domain.UserAggregatedExercise, 0), nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.GetRecentExercisesChart(ctx, database.NewStringID(), *timezone.UTC, dto.GetRecentExercisesChartRequest{})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), 0, len(resp.Exercises))
}

//
// END OF CASES
//

func TestGetRecentExercisesChartTestSuite(t *testing.T) {
	suite.Run(t, new(getRecentExercisesChartTestSuite))
}
