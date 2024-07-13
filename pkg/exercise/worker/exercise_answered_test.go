package worker_test

import (
	"context"
	"testing"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	mockexercise "github.com/namhq1989/vocab-booster-server-app/internal/mock/exercise"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/worker"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type exerciseAnsweredTestSuite struct {
	suite.Suite
	handler             worker.ExerciseAnsweredHandler
	mockCtrl            *gomock.Controller
	mockQueueRepository *mockexercise.MockQueueRepository
}

func (s *exerciseAnsweredTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *exerciseAnsweredTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockQueueRepository = mockexercise.NewMockQueueRepository(s.mockCtrl)

	s.handler = worker.NewExerciseAnsweredHandler(s.mockQueueRepository)
}

func (s *exerciseAnsweredTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *exerciseAnsweredTestSuite) Test_1_Success() {
	// mock data
	s.mockQueueRepository.EXPECT().
		GamificationExerciseAnswered(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewGRPC(context.Background())
	err := s.handler.ExerciseAnswered(ctx, domain.QueueExerciseAnsweredPayload{
		UserID:         database.NewStringID(),
		ExerciseID:     database.NewStringID(),
		Point:          30,
		CompletionTime: 5,
	})

	assert.Nil(s.T(), err)
}

//
// END OF CASES
//

func TestExerciseAnsweredTestSuite(t *testing.T) {
	suite.Run(t, new(exerciseAnsweredTestSuite))
}
