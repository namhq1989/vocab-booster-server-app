package command_test

import (
	"context"
	"testing"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	mockvocabulary "github.com/namhq1989/vocab-booster-server-app/internal/mock/vocabulary"
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/application/command"
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type bookmarkVocabularyTestSuite struct {
	suite.Suite
	handler           command.BookmarkVocabularyHandler
	mockCtrl          *gomock.Controller
	mockVocabularyHub *mockvocabulary.MockVocabularyHub
}

func (s *bookmarkVocabularyTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *bookmarkVocabularyTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockVocabularyHub = mockvocabulary.NewMockVocabularyHub(s.mockCtrl)

	s.handler = command.NewBookmarkVocabularyHandler(s.mockVocabularyHub)
}

func (s *bookmarkVocabularyTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *bookmarkVocabularyTestSuite) Test_1_Success() {
	// mock data
	s.mockVocabularyHub.EXPECT().
		BookmarkVocabulary(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(true, nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.BookmarkVocabulary(ctx, database.NewStringID(), database.NewStringID(), dto.BookmarkVocabularyRequest{})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), true, resp.IsBookmarked)
}

//
// END OF CASES
//

func TestBookmarkVocabularyTestSuite(t *testing.T) {
	suite.Run(t, new(bookmarkVocabularyTestSuite))
}
