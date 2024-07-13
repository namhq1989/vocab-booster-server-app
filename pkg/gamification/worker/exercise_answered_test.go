package worker_test

import (
	"context"
	"testing"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	mockgamification "github.com/namhq1989/vocab-booster-server-app/internal/mock/gamification"
	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
	"github.com/namhq1989/vocab-booster-server-app/pkg/gamification/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/gamification/worker"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type exerciseAnsweredTestSuite struct {
	suite.Suite
	handler                      worker.ExerciseAnsweredHandler
	mockCtrl                     *gomock.Controller
	mockPointRepository          *mockgamification.MockPointRepository
	mockCompletionTimeRepository *mockgamification.MockCompletionTimeRepository
	mockUserStatsRepository      *mockgamification.MockUserStatsRepository
	mockService                  *mockgamification.MockService
}

func (s *exerciseAnsweredTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *exerciseAnsweredTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockPointRepository = mockgamification.NewMockPointRepository(s.mockCtrl)
	s.mockCompletionTimeRepository = mockgamification.NewMockCompletionTimeRepository(s.mockCtrl)
	s.mockUserStatsRepository = mockgamification.NewMockUserStatsRepository(s.mockCtrl)
	s.mockService = mockgamification.NewMockService(s.mockCtrl)

	s.handler = worker.NewExerciseAnsweredHandler(s.mockPointRepository, s.mockCompletionTimeRepository, s.mockUserStatsRepository, s.mockService)
}

func (s *exerciseAnsweredTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *exerciseAnsweredTestSuite) Test_1_Success() {
	// mock data
	s.mockService.EXPECT().
		ExerciseAnswered(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewGRPC(context.Background())
	err := s.handler.ExerciseAnswered(ctx, domain.QueueExerciseAnsweredPoint{
		UserID:         database.NewStringID(),
		ExerciseID:     database.NewStringID(),
		Point:          30,
		CompletionTime: 5,
	})

	assert.Nil(s.T(), err)
}

func (s *exerciseAnsweredTestSuite) Test_1_Fail_InvalidData() {
	// call
	ctx := appcontext.NewGRPC(context.Background())
	err := s.handler.ExerciseAnswered(ctx, domain.QueueExerciseAnsweredPoint{
		UserID:         database.NewStringID(),
		ExerciseID:     "",
		Point:          30,
		CompletionTime: 5,
	})

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), apperrors.Gamification.InvalidPointData, err)
}

func (s *exerciseAnsweredTestSuite) Test_1_Fail_InvalidExercise() {
	// call
	ctx := appcontext.NewGRPC(context.Background())
	err := s.handler.ExerciseAnswered(ctx, domain.QueueExerciseAnsweredPoint{
		UserID:         database.NewStringID(),
		ExerciseID:     "invalid id",
		Point:          30,
		CompletionTime: 5,
	})

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), apperrors.Exercise.InvalidExerciseID, err)
}

func (s *exerciseAnsweredTestSuite) Test_1_Fail_InvalidPoint() {
	// call
	ctx := appcontext.NewGRPC(context.Background())
	err := s.handler.ExerciseAnswered(ctx, domain.QueueExerciseAnsweredPoint{
		UserID:         database.NewStringID(),
		ExerciseID:     database.NewStringID(),
		Point:          1000,
		CompletionTime: 5,
	})

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), apperrors.Gamification.InvalidPoint, err)
}

func (s *exerciseAnsweredTestSuite) Test_1_Fail_InvalidCompletionTime() {
	// call
	ctx := appcontext.NewGRPC(context.Background())
	err := s.handler.ExerciseAnswered(ctx, domain.QueueExerciseAnsweredPoint{
		UserID:         database.NewStringID(),
		ExerciseID:     database.NewStringID(),
		Point:          30,
		CompletionTime: -1,
	})

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), apperrors.Gamification.InvalidCompletionTime, err)
}

func (s *exerciseAnsweredTestSuite) Test_1_Fail_InvalidUserID() {
	// call
	ctx := appcontext.NewGRPC(context.Background())
	err := s.handler.ExerciseAnswered(ctx, domain.QueueExerciseAnsweredPoint{
		UserID:         "invalid id",
		ExerciseID:     database.NewStringID(),
		Point:          30,
		CompletionTime: 5,
	})

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), apperrors.User.InvalidUserID, err)
}

//
// END OF CASES
//

func TestExerciseAnsweredTestSuite(t *testing.T) {
	suite.Run(t, new(exerciseAnsweredTestSuite))
}
