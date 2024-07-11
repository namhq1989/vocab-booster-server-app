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

type getReadyForReviewExercisesTest struct {
	suite.Suite
	handler         query.GetReadyForReviewExercisesHandler
	mockCtrl        *gomock.Controller
	mockExerciseHub *mockexercise.MockExerciseHub
}

func (s *getReadyForReviewExercisesTest) SetupSuite() {
	s.setupApplication()
}

func (s *getReadyForReviewExercisesTest) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockExerciseHub = mockexercise.NewMockExerciseHub(s.mockCtrl)

	s.handler = query.NewGetReadyForReviewExercisesHandler(s.mockExerciseHub)
}

func (s *getReadyForReviewExercisesTest) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *getReadyForReviewExercisesTest) Test_1_Success() {
	// mock data
	s.mockExerciseHub.EXPECT().
		GetReadyForReviewExercises(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(make([]domain.Exercise, 0), nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.GetReadyForReviewExercises(ctx, database.NewStringID(), language.Vietnamese, dto.GetReadyForReviewExercisesRequest{})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), 0, len(resp.Exercises))
}

//
// END OF CASES
//

func TestGetReadyForReviewExercisesTest(t *testing.T) {
	suite.Run(t, new(getReadyForReviewExercisesTest))
}
