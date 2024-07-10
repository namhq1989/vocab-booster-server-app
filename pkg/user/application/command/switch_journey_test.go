package command_test

import (
	"context"
	"testing"
	"time"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	mockuser "github.com/namhq1989/vocab-booster-server-app/internal/mock/user"
	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/application/command"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type switchJourneyTestSuite struct {
	suite.Suite
	handler               command.SwitchJourneyHandler
	mockCtrl              *gomock.Controller
	mockJourneyRepository *mockuser.MockJourneyRepository
}

func (s *switchJourneyTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *switchJourneyTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockJourneyRepository = mockuser.NewMockJourneyRepository(s.mockCtrl)

	s.handler = command.NewSwitchJourneyHandler(s.mockJourneyRepository)
}

func (s *switchJourneyTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *switchJourneyTestSuite) Test_1_SuccessWithIsLearning() {
	// mock data
	s.mockJourneyRepository.EXPECT().
		FindJourneysByUserID(gomock.Any(), gomock.Any()).
		Return([]domain.Journey{
			{
				ID:         database.NewStringID(),
				UserID:     database.NewStringID(),
				Lang:       language.Vietnamese,
				IsLearning: true,
				Stats:      domain.JourneyStats{},
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			},
		}, nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.SwitchJourney(ctx, database.NewStringID(), dto.SwitchJourneyRequest{
		Lang: language.Vietnamese.String(),
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
}

func (s *switchJourneyTestSuite) Test_1_SuccessWithIsNotLearning() {
	// mock data
	s.mockJourneyRepository.EXPECT().
		FindJourneysByUserID(gomock.Any(), gomock.Any()).
		Return([]domain.Journey{
			{
				ID:         database.NewStringID(),
				UserID:     database.NewStringID(),
				Lang:       language.Vietnamese,
				IsLearning: false,
				Stats:      domain.JourneyStats{},
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			},
			{
				ID:         database.NewStringID(),
				UserID:     database.NewStringID(),
				Lang:       language.English,
				IsLearning: true,
				Stats:      domain.JourneyStats{},
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			},
		}, nil)

	s.mockJourneyRepository.EXPECT().
		UpdateJourney(gomock.Any(), gomock.Any()).
		Times(2).
		Return(nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.SwitchJourney(ctx, database.NewStringID(), dto.SwitchJourneyRequest{
		Lang: language.Vietnamese.String(),
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
}

func (s *switchJourneyTestSuite) Test_2_Fail_LangIsEnglish() {
	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.SwitchJourney(ctx, database.NewStringID(), dto.SwitchJourneyRequest{
		Lang: language.English.String(),
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.InvalidLanguage, err)
}

func (s *switchJourneyTestSuite) Test_2_Fail_JourneyNotFound() {
	// mock data
	s.mockJourneyRepository.EXPECT().
		FindJourneysByUserID(gomock.Any(), gomock.Any()).
		Return(make([]domain.Journey, 0), nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.SwitchJourney(ctx, database.NewStringID(), dto.SwitchJourneyRequest{
		Lang: language.Vietnamese.String(),
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.User.JourneyNotFound, err)
}

//
// END OF CASES
//

func TestSwitchJourneyTestSuite(t *testing.T) {
	suite.Run(t, new(switchJourneyTestSuite))
}
