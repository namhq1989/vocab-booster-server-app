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

type vocabularySentenceContributedTestSuite struct {
	suite.Suite
	handler                      worker.VocabularySentenceContributedHandler
	mockCtrl                     *gomock.Controller
	mockPointRepository          *mockgamification.MockPointRepository
	mockCompletionTimeRepository *mockgamification.MockCompletionTimeRepository
	mockUserStatsRepository      *mockgamification.MockUserStatsRepository
	mockService                  *mockgamification.MockService
}

func (s *vocabularySentenceContributedTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *vocabularySentenceContributedTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockPointRepository = mockgamification.NewMockPointRepository(s.mockCtrl)
	s.mockCompletionTimeRepository = mockgamification.NewMockCompletionTimeRepository(s.mockCtrl)
	s.mockUserStatsRepository = mockgamification.NewMockUserStatsRepository(s.mockCtrl)
	s.mockService = mockgamification.NewMockService(s.mockCtrl)

	s.handler = worker.NewVocabularySentenceContributedHandler(s.mockPointRepository, s.mockCompletionTimeRepository, s.mockUserStatsRepository, s.mockService)
}

func (s *vocabularySentenceContributedTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *vocabularySentenceContributedTestSuite) Test_1_Success() {
	// mock data
	s.mockService.EXPECT().
		ExerciseAnswered(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewGRPC(context.Background())
	err := s.handler.VocabularySentenceContributed(ctx, domain.QueueVocabularySentenceContributedPoint{
		UserID:         database.NewStringID(),
		VocabularyID:   database.NewStringID(),
		Point:          30,
		CompletionTime: 5,
	})

	assert.Nil(s.T(), err)
}

func (s *vocabularySentenceContributedTestSuite) Test_1_Fail_InvalidData() {
	// call
	ctx := appcontext.NewGRPC(context.Background())
	err := s.handler.VocabularySentenceContributed(ctx, domain.QueueVocabularySentenceContributedPoint{
		UserID:         database.NewStringID(),
		VocabularyID:   "",
		Point:          30,
		CompletionTime: 5,
	})

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), apperrors.Gamification.InvalidPointData, err)
}

func (s *vocabularySentenceContributedTestSuite) Test_1_Fail_InvalidVocabulary() {
	// call
	ctx := appcontext.NewGRPC(context.Background())
	err := s.handler.VocabularySentenceContributed(ctx, domain.QueueVocabularySentenceContributedPoint{
		UserID:         database.NewStringID(),
		VocabularyID:   "invalid id",
		Point:          30,
		CompletionTime: 5,
	})

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), apperrors.Vocabulary.InvalidVocabularyID, err)
}

func (s *vocabularySentenceContributedTestSuite) Test_1_Fail_InvalidPoint() {
	// call
	ctx := appcontext.NewGRPC(context.Background())
	err := s.handler.VocabularySentenceContributed(ctx, domain.QueueVocabularySentenceContributedPoint{
		UserID:         database.NewStringID(),
		VocabularyID:   database.NewStringID(),
		Point:          1000,
		CompletionTime: 5,
	})

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), apperrors.Gamification.InvalidPoint, err)
}

func (s *vocabularySentenceContributedTestSuite) Test_1_Fail_InvalidCompletionTime() {
	// call
	ctx := appcontext.NewGRPC(context.Background())
	err := s.handler.VocabularySentenceContributed(ctx, domain.QueueVocabularySentenceContributedPoint{
		UserID:         database.NewStringID(),
		VocabularyID:   database.NewStringID(),
		Point:          30,
		CompletionTime: -1,
	})

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), apperrors.Gamification.InvalidCompletionTime, err)
}

func (s *vocabularySentenceContributedTestSuite) Test_1_Fail_InvalidUserID() {
	// call
	ctx := appcontext.NewGRPC(context.Background())
	err := s.handler.VocabularySentenceContributed(ctx, domain.QueueVocabularySentenceContributedPoint{
		UserID:         "invalid id",
		VocabularyID:   database.NewStringID(),
		Point:          30,
		CompletionTime: 5,
	})

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), apperrors.User.InvalidUserID, err)
}

//
// END OF CASES
//

func TestVocabularySentenceContributedTestSuite(t *testing.T) {
	suite.Run(t, new(vocabularySentenceContributedTestSuite))
}
