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

type searchVocabularyTestSuite struct {
	suite.Suite
	handler           query.SearchVocabularyHandler
	mockCtrl          *gomock.Controller
	mockVocabularyHub *mockvocabulary.MockVocabularyHub
}

func (s *searchVocabularyTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *searchVocabularyTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockVocabularyHub = mockvocabulary.NewMockVocabularyHub(s.mockCtrl)

	s.handler = query.NewSearchVocabularyHandler(s.mockVocabularyHub)
}

func (s *searchVocabularyTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *searchVocabularyTestSuite) Test_1_Success() {
	// mock data
	var term = "random"

	s.mockVocabularyHub.EXPECT().
		SearchVocabulary(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&domain.Vocabulary{
			ID:       database.NewStringID(),
			AuthorID: database.NewStringID(),
			Term:     term,
		}, make([]string, 0), nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.SearchVocabulary(ctx, database.NewStringID(), language.Vietnamese, dto.SearchVocabularyRequest{
		Term: term,
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.NotNil(s.T(), resp.Vocabulary)
	assert.Equal(s.T(), term, resp.Vocabulary.Term)
}

//
// END OF CASES
//

func TestSearchVocabularyTestSuite(t *testing.T) {
	suite.Run(t, new(searchVocabularyTestSuite))
}
