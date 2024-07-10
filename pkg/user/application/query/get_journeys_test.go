package query_test

import (
	"context"
	"testing"
	"time"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	mockuser "github.com/namhq1989/vocab-booster-server-app/internal/mock/user"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/application/query"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type getJourneysTest struct {
	suite.Suite
	handler               query.GetJourneysHandler
	mockCtrl              *gomock.Controller
	mockJourneyRepository *mockuser.MockJourneyRepository
}

func (s *getJourneysTest) SetupSuite() {
	s.setupApplication()
}

func (s *getJourneysTest) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockJourneyRepository = mockuser.NewMockJourneyRepository(s.mockCtrl)

	s.handler = query.NewGetJourneysHandler(s.mockJourneyRepository)
}

func (s *getJourneysTest) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *getJourneysTest) Test_1_Success() {
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
	resp, err := s.handler.GetJourneys(ctx, database.NewStringID(), dto.GetJourneysRequest{})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), 1, len(resp.Journeys))
}

//
// END OF CASES
//

func TestGetJourneysTest(t *testing.T) {
	suite.Run(t, new(getJourneysTest))
}
