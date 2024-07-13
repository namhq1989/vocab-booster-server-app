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

type addContributeVocabularySentencePointTestSuite struct {
	suite.Suite
	handler                 worker.AddContributeVocabularySentencePointHandler
	mockCtrl                *gomock.Controller
	mockPointRepository     *mockgamification.MockPointRepository
	mockUserPointRepository *mockgamification.MockUserPointRepository
	mockService             *mockgamification.MockService
}

func (s *addContributeVocabularySentencePointTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *addContributeVocabularySentencePointTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockPointRepository = mockgamification.NewMockPointRepository(s.mockCtrl)
	s.mockUserPointRepository = mockgamification.NewMockUserPointRepository(s.mockCtrl)
	s.mockService = mockgamification.NewMockService(s.mockCtrl)

	s.handler = worker.NewAddContributeVocabularySentencePointHandler(s.mockPointRepository, s.mockUserPointRepository, s.mockService)
}

func (s *addContributeVocabularySentencePointTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *addContributeVocabularySentencePointTestSuite) Test_1_Success() {
	// mock data
	s.mockService.EXPECT().
		AddPoint(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewGRPC(context.Background())
	err := s.handler.AddContributeVocabularySentencePoint(ctx, domain.QueueAddContributeVocabularySentencePoint{
		UserID:       database.NewStringID(),
		VocabularyID: database.NewStringID(),
		Point:        30,
	})

	assert.Nil(s.T(), err)
}

func (s *addContributeVocabularySentencePointTestSuite) Test_1_Fail_InvalidData() {
	// call
	ctx := appcontext.NewGRPC(context.Background())
	err := s.handler.AddContributeVocabularySentencePoint(ctx, domain.QueueAddContributeVocabularySentencePoint{
		UserID:       database.NewStringID(),
		VocabularyID: "",
		Point:        30,
	})

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), apperrors.Gamification.InvalidPointData, err)
}

func (s *addContributeVocabularySentencePointTestSuite) Test_1_Fail_InvalidVocabulary() {
	// call
	ctx := appcontext.NewGRPC(context.Background())
	err := s.handler.AddContributeVocabularySentencePoint(ctx, domain.QueueAddContributeVocabularySentencePoint{
		UserID:       database.NewStringID(),
		VocabularyID: "invalid id",
		Point:        30,
	})

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), apperrors.Vocabulary.InvalidVocabularyID, err)
}

func (s *addContributeVocabularySentencePointTestSuite) Test_1_Fail_InvalidPoint() {
	// call
	ctx := appcontext.NewGRPC(context.Background())
	err := s.handler.AddContributeVocabularySentencePoint(ctx, domain.QueueAddContributeVocabularySentencePoint{
		UserID:       database.NewStringID(),
		VocabularyID: database.NewStringID(),
		Point:        1000,
	})

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), apperrors.Gamification.InvalidPoint, err)
}

func (s *addContributeVocabularySentencePointTestSuite) Test_1_Fail_InvalidUserID() {
	// call
	ctx := appcontext.NewGRPC(context.Background())
	err := s.handler.AddContributeVocabularySentencePoint(ctx, domain.QueueAddContributeVocabularySentencePoint{
		UserID:       "invalid id",
		VocabularyID: database.NewStringID(),
		Point:        30,
	})

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), apperrors.User.InvalidUserID, err)
}

//
// END OF CASES
//

func TestAddContributeVocabularySentencePointTestSuite(t *testing.T) {
	suite.Run(t, new(addContributeVocabularySentencePointTestSuite))
}
