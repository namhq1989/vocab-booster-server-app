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

type getCommunitySentencesTestSuite struct {
	suite.Suite
	handler           query.GetCommunitySentencesHandler
	mockCtrl          *gomock.Controller
	mockVocabularyHub *mockvocabulary.MockVocabularyHub
}

func (s *getCommunitySentencesTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *getCommunitySentencesTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockVocabularyHub = mockvocabulary.NewMockVocabularyHub(s.mockCtrl)

	s.handler = query.NewGetCommunitySentencesHandler(s.mockVocabularyHub)
}

func (s *getCommunitySentencesTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *getCommunitySentencesTestSuite) Test_1_Success() {
	// mock data
	s.mockVocabularyHub.EXPECT().
		GetCommunitySentences(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(make([]domain.CommunitySentenceBrief, 0), "", nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.GetCommunitySentences(ctx, database.NewStringID(), database.NewStringID(), language.English, dto.GetCommunitySentencesRequest{
		PageToken: "",
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), 0, len(resp.Sentences))
}

//
// END OF CASES
//

func TestGetCommunitySentencesTestSuite(t *testing.T) {
	suite.Run(t, new(getCommunitySentencesTestSuite))
}
