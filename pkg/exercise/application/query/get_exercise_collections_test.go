package query_test

import (
	"context"
	"testing"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	mockexercise "github.com/namhq1989/vocab-booster-server-app/internal/mock/exercise"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/application/query"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type getExerciseCollectionTestSuite struct {
	suite.Suite
	handler         query.GetExerciseCollectionsHandler
	mockCtrl        *gomock.Controller
	mockExerciseHub *mockexercise.MockExerciseHub
}

func (s *getExerciseCollectionTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *getExerciseCollectionTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockExerciseHub = mockexercise.NewMockExerciseHub(s.mockCtrl)

	s.handler = query.NewGetExerciseCollectionsHandler(s.mockExerciseHub)
}

func (s *getExerciseCollectionTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *getExerciseCollectionTestSuite) Test_1_Success() {
	// mock data
	s.mockExerciseHub.EXPECT().
		GetExerciseCollections(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(make([]domain.ExerciseCollection, 0), nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.GetExerciseCollections(ctx, database.NewStringID(), language.Vietnamese, dto.GetExerciseCollectionsRequest{})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), 0, len(resp.Collections))
}

//
// END OF CASES
//

func TestGetExerciseCollectionTestSuite(t *testing.T) {
	suite.Run(t, new(getExerciseCollectionTestSuite))
}
