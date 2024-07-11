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

type getExercisesTest struct {
	suite.Suite
	handler         query.GetExercisesHandler
	mockCtrl        *gomock.Controller
	mockExerciseHub *mockexercise.MockExerciseHub
}

func (s *getExercisesTest) SetupSuite() {
	s.setupApplication()
}

func (s *getExercisesTest) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockExerciseHub = mockexercise.NewMockExerciseHub(s.mockCtrl)

	s.handler = query.NewGetExercisesHandler(s.mockExerciseHub)
}

func (s *getExercisesTest) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *getExercisesTest) Test_1_Success() {
	// mock data
	s.mockExerciseHub.EXPECT().
		GetExercises(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(make([]domain.Exercise, 0), nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.GetExercises(ctx, database.NewStringID(), language.Vietnamese, dto.GetExercisesRequest{})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), 0, len(resp.Exercises))
}

//
// END OF CASES
//

func TestGetExercisesTest(t *testing.T) {
	suite.Run(t, new(getExercisesTest))
}
