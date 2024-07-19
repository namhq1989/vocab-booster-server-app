package command_test

import (
	"context"
	"testing"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	mockexercise "github.com/namhq1989/vocab-booster-server-app/internal/mock/exercise"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/application/command"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type changeExerciseFavoriteTestSuite struct {
	suite.Suite
	handler         command.ChangeExerciseFavoriteHandler
	mockCtrl        *gomock.Controller
	mockExerciseHub *mockexercise.MockExerciseHub
}

func (s *changeExerciseFavoriteTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *changeExerciseFavoriteTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockExerciseHub = mockexercise.NewMockExerciseHub(s.mockCtrl)

	s.handler = command.NewChangeExerciseFavoriteHandler(s.mockExerciseHub)
}

func (s *changeExerciseFavoriteTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *changeExerciseFavoriteTestSuite) Test_1_Success() {
	// mock data
	s.mockExerciseHub.EXPECT().
		ChangeExerciseFavorite(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(true, nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.ChangeExerciseFavorite(ctx, database.NewStringID(), database.NewStringID(), dto.ChangeExerciseFavoriteRequest{
		IsFavorite: true,
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), true, resp.IsFavorite)
}

//
// END OF CASES
//

func TestChangeExerciseFavoriteTestSuite(t *testing.T) {
	suite.Run(t, new(changeExerciseFavoriteTestSuite))
}
