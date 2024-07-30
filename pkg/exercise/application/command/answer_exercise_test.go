package command_test

import (
	"context"
	"testing"

	"github.com/namhq1989/vocab-booster-utilities/timezone"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	mockexercise "github.com/namhq1989/vocab-booster-server-app/internal/mock/exercise"
	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/httprespond"
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/manipulation"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/application/command"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type answerExerciseTestSuite struct {
	suite.Suite
	handler             command.AnswerExerciseHandler
	mockCtrl            *gomock.Controller
	mockExerciseHub     *mockexercise.MockExerciseHub
	mockQueueRepository *mockexercise.MockQueueRepository
}

func (s *answerExerciseTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *answerExerciseTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockExerciseHub = mockexercise.NewMockExerciseHub(s.mockCtrl)
	s.mockQueueRepository = mockexercise.NewMockQueueRepository(s.mockCtrl)

	s.handler = command.NewAnswerExerciseHandler(s.mockQueueRepository, s.mockExerciseHub)
}

func (s *answerExerciseTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *answerExerciseTestSuite) Test_1_Success() {
	// mock data
	nextReviewAt := manipulation.NowUTC()

	s.mockExerciseHub.EXPECT().
		AnswerExercise(gomock.Any(), gomock.Any()).
		Return(&domain.AnswerExerciseResult{NextReviewAt: nextReviewAt}, nil)

	s.mockQueueRepository.EXPECT().
		ExerciseAnswered(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.AnswerExercise(ctx, database.NewStringID(), database.NewStringID(), *timezone.UTC, dto.AnswerExerciseRequest{
		IsCorrect:      true,
		CompletionTime: 10,
		Point:          10,
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), httprespond.NewTimeResponse(nextReviewAt), resp.NextReviewAt)
}

func (s *answerExerciseTestSuite) Test_2_Fail_InvalidUserID() {
	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.AnswerExercise(ctx, "invalid id", database.NewStringID(), *timezone.UTC, dto.AnswerExerciseRequest{
		IsCorrect:      true,
		CompletionTime: 10,
		Point:          10,
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.User.InvalidUserID, err)
}

func (s *answerExerciseTestSuite) Test_2_Fail_InvalidExerciseID() {
	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.AnswerExercise(ctx, database.NewStringID(), "invalid id", *timezone.UTC, dto.AnswerExerciseRequest{
		IsCorrect:      true,
		CompletionTime: 10,
		Point:          10,
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Exercise.InvalidExerciseID, err)
}

//
// END OF CASES
//

func TestAnswerExerciseTestSuite(t *testing.T) {
	suite.Run(t, new(answerExerciseTestSuite))
}
