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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type getUserBookmarkedVocabulariesTestSuite struct {
	suite.Suite
	handler           query.GetUserBookmarkedVocabulariesHandler
	mockCtrl          *gomock.Controller
	mockVocabularyHub *mockvocabulary.MockVocabularyHub
}

func (s *getUserBookmarkedVocabulariesTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *getUserBookmarkedVocabulariesTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockVocabularyHub = mockvocabulary.NewMockVocabularyHub(s.mockCtrl)

	s.handler = query.NewGetUserBookmarkedVocabulariesHandler(s.mockVocabularyHub)
}

func (s *getUserBookmarkedVocabulariesTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *getUserBookmarkedVocabulariesTestSuite) Test_1_Success() {
	// mock data
	s.mockVocabularyHub.EXPECT().
		GetUserBookmarkedVocabularies(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(make([]domain.VocabularyBrief, 0), "", nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.GetUserBookmarkedVocabularies(ctx, database.NewStringID(), dto.GetUserBookmarkedVocabulariesRequest{
		PageToken: "",
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.NotNil(s.T(), 0, len(resp.Vocabularies))
}

//
// END OF CASES
//

func TestGetUserBookmarkedVocabulariesTestSuite(t *testing.T) {
	suite.Run(t, new(getUserBookmarkedVocabulariesTestSuite))
}
