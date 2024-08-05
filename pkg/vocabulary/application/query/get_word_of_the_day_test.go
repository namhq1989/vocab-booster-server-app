package query_test

import (
	"context"
	"testing"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	mockvocabulary "github.com/namhq1989/vocab-booster-server-app/internal/mock/vocabulary"
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/application/query"
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type getWordOfTheDayTestSuite struct {
	suite.Suite
	handler           query.GetWordOfTheDayHandler
	mockCtrl          *gomock.Controller
	mockVocabularyHub *mockvocabulary.MockVocabularyHub
}

func (s *getWordOfTheDayTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *getWordOfTheDayTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockVocabularyHub = mockvocabulary.NewMockVocabularyHub(s.mockCtrl)

	s.handler = query.NewGetWordOfTheDayHandler(s.mockVocabularyHub)
}

func (s *getWordOfTheDayTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *getWordOfTheDayTestSuite) Test_1_Success() {
	// mock data
	s.mockVocabularyHub.EXPECT().
		GetWordOfTheDay(gomock.Any(), gomock.Any()).
		Return(&domain.WordOfTheDay{
			Vocabulary: domain.VocabularyBrief{
				ID:            database.NewStringID(),
				Term:          "hello",
				PartsOfSpeech: make([]string, 0),
				Ipa:           "",
				Audio:         "",
			},
		}, nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.GetWordOfTheDay(ctx, database.NewStringID(), language.Vietnamese, dto.GetWordOfTheDayRequest{})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.NotNil(s.T(), resp.Vocabulary)
}

//
// END OF CASES
//

func TestGetWordOfTheDayTestSuite(t *testing.T) {
	suite.Run(t, new(getWordOfTheDayTestSuite))
}
