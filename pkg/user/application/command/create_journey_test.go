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

type createJourneyTestSuite struct {
	suite.Suite
	handler               command.CreateJourneyHandler
	mockCtrl              *gomock.Controller
	mockJourneyRepository *mockuser.MockJourneyRepository
}

func (s *createJourneyTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *createJourneyTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockJourneyRepository = mockuser.NewMockJourneyRepository(s.mockCtrl)

	s.handler = command.NewCreateJourneyHandler(s.mockJourneyRepository)
}

func (s *createJourneyTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *createJourneyTestSuite) Test_1_SuccessWithNotExistedJourney() {
	// mock data
	s.mockJourneyRepository.EXPECT().
		FindJourneysByUserID(gomock.Any(), gomock.Any()).
		Return(make([]domain.Journey, 0), nil)

	s.mockJourneyRepository.EXPECT().
		CreateJourney(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.CreateJourney(ctx, database.NewStringID(), dto.CreateJourneyRequest{
		Lang: language.Vietnamese.String(),
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
}

func (s *createJourneyTestSuite) Test_1_SuccessWithExistedJourneyAndIsLearning() {
	// mock data
	var journeyID = database.NewStringID()

	s.mockJourneyRepository.EXPECT().
		FindJourneysByUserID(gomock.Any(), gomock.Any()).
		Return([]domain.Journey{
			{
				ID:         journeyID,
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
	resp, err := s.handler.CreateJourney(ctx, database.NewStringID(), dto.CreateJourneyRequest{
		Lang: language.Vietnamese.String(),
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), journeyID, resp.ID)
}

func (s *createJourneyTestSuite) Test_1_SuccessWithExistedJourneyAndIsNotLearning() {
	// mock data
	var journeyID = database.NewStringID()

	s.mockJourneyRepository.EXPECT().
		FindJourneysByUserID(gomock.Any(), gomock.Any()).
		Return([]domain.Journey{
			{
				ID:         journeyID,
				UserID:     database.NewStringID(),
				Lang:       language.Vietnamese,
				IsLearning: false,
				Stats:      domain.JourneyStats{},
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			},
		}, nil)

	s.mockJourneyRepository.EXPECT().
		UpdateJourney(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.CreateJourney(ctx, database.NewStringID(), dto.CreateJourneyRequest{
		Lang: language.Vietnamese.String(),
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), journeyID, resp.ID)
}

func (s *createJourneyTestSuite) Test_2_Fail_LangIsEnglish() {
	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.CreateJourney(ctx, database.NewStringID(), dto.CreateJourneyRequest{
		Lang: language.English.String(),
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.InvalidLanguage, err)
}

func (s *createJourneyTestSuite) Test_2_Fail_InvalidUserID() {
	// mock data
	s.mockJourneyRepository.EXPECT().
		FindJourneysByUserID(gomock.Any(), gomock.Any()).
		Return(make([]domain.Journey, 0), nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.CreateJourney(ctx, "invalid id", dto.CreateJourneyRequest{
		Lang: language.Vietnamese.String(),
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.User.InvalidUserID, err)
}

func (s *createJourneyTestSuite) Test_2_Fail_InvalidLang() {
	// mock data
	s.mockJourneyRepository.EXPECT().
		FindJourneysByUserID(gomock.Any(), gomock.Any()).
		Return(make([]domain.Journey, 0), nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.CreateJourney(ctx, database.NewStringID(), dto.CreateJourneyRequest{
		Lang: "invalid lang",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.InvalidLanguage, err)
}

//
// END OF CASES
//

func TestCreateJourneyTestSuite(t *testing.T) {
	suite.Run(t, new(createJourneyTestSuite))
}
