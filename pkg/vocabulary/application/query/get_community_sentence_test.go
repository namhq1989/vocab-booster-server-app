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

type getCommunitySentenceTestSuite struct {
	suite.Suite
	handler           query.GetCommunitySentenceHandler
	mockCtrl          *gomock.Controller
	mockVocabularyHub *mockvocabulary.MockVocabularyHub
}

func (s *getCommunitySentenceTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *getCommunitySentenceTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockVocabularyHub = mockvocabulary.NewMockVocabularyHub(s.mockCtrl)

	s.handler = query.NewGetCommunitySentenceHandler(s.mockVocabularyHub)
}

func (s *getCommunitySentenceTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *getCommunitySentenceTestSuite) Test_1_Success() {
	// mock data
	s.mockVocabularyHub.EXPECT().
		GetCommunitySentence(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&domain.CommunitySentence{
			ID:           database.NewStringID(),
			VocabularyID: database.NewStringID(),
		}, nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.GetCommunitySentence(ctx, database.NewStringID(), database.NewStringID(), language.English, dto.GetCommunitySentenceRequest{})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.NotNil(s.T(), resp.Sentence)
}

//
// END OF CASES
//

func TestGetCommunitySentenceTestSuite(t *testing.T) {
	suite.Run(t, new(getCommunitySentenceTestSuite))
}
