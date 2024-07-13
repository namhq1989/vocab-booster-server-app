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

type addAnswerExercisePointTestSuite struct {
	suite.Suite
	handler                 worker.AddAnswerExercisePointHandler
	mockCtrl                *gomock.Controller
	mockPointRepository     *mockgamification.MockPointRepository
	mockUserPointRepository *mockgamification.MockUserPointRepository
	mockService             *mockgamification.MockService
}

func (s *addAnswerExercisePointTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *addAnswerExercisePointTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockPointRepository = mockgamification.NewMockPointRepository(s.mockCtrl)
	s.mockUserPointRepository = mockgamification.NewMockUserPointRepository(s.mockCtrl)
	s.mockService = mockgamification.NewMockService(s.mockCtrl)

	s.handler = worker.NewAddAnswerExercisePointHandler(s.mockPointRepository, s.mockUserPointRepository, s.mockService)
}

func (s *addAnswerExercisePointTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *addAnswerExercisePointTestSuite) Test_1_Success() {
	// mock data
	s.mockService.EXPECT().
		AddPoint(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewGRPC(context.Background())
	err := s.handler.AddAnswerExercisePoint(ctx, domain.QueueAddAnswerExercisePoint{
		UserID:     database.NewStringID(),
		ExerciseID: database.NewStringID(),
		Point:      30,
	})

	assert.Nil(s.T(), err)
}

func (s *addAnswerExercisePointTestSuite) Test_1_Fail_InvalidData() {
	// call
	ctx := appcontext.NewGRPC(context.Background())
	err := s.handler.AddAnswerExercisePoint(ctx, domain.QueueAddAnswerExercisePoint{
		UserID:     database.NewStringID(),
		ExerciseID: "",
		Point:      30,
	})

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), apperrors.Gamification.InvalidPointData, err)
}

func (s *addAnswerExercisePointTestSuite) Test_1_Fail_InvalidExercise() {
	// call
	ctx := appcontext.NewGRPC(context.Background())
	err := s.handler.AddAnswerExercisePoint(ctx, domain.QueueAddAnswerExercisePoint{
		UserID:     database.NewStringID(),
		ExerciseID: "invalid id",
		Point:      30,
	})

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), apperrors.Exercise.InvalidExerciseID, err)
}

func (s *addAnswerExercisePointTestSuite) Test_1_Fail_InvalidPoint() {
	// call
	ctx := appcontext.NewGRPC(context.Background())
	err := s.handler.AddAnswerExercisePoint(ctx, domain.QueueAddAnswerExercisePoint{
		UserID:     database.NewStringID(),
		ExerciseID: database.NewStringID(),
		Point:      1000,
	})

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), apperrors.Gamification.InvalidPoint, err)
}

func (s *addAnswerExercisePointTestSuite) Test_1_Fail_InvalidUserID() {
	// call
	ctx := appcontext.NewGRPC(context.Background())
	err := s.handler.AddAnswerExercisePoint(ctx, domain.QueueAddAnswerExercisePoint{
		UserID:     "invalid id",
		ExerciseID: database.NewStringID(),
		Point:      30,
	})

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), apperrors.User.InvalidUserID, err)
}

//
// END OF CASES
//

func TestAddAnswerExercisePointTestSuite(t *testing.T) {
	suite.Run(t, new(addAnswerExercisePointTestSuite))
}
