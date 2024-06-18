package query_test

import (
	"context"
	"testing"

	"github.com/namhq1989/vocab-booster-server-app/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-server-app/core/error"
	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	mockmongo "github.com/namhq1989/vocab-booster-server-app/internal/mock/mongo"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/application/query"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type getMeTestSuite struct {
	suite.Suite
	handler            query.GetMeHandler
	mockCtrl           *gomock.Controller
	mockUserRepository *mockmongo.MockUserRepository
}

func (s *getMeTestSuite) SetupSuite() {
	s.setupApplication()
}

func (*getMeTestSuite) AfterTest(_, _ string) {
	// do nothing
}

func (s *getMeTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockUserRepository = mockmongo.NewMockUserRepository(s.mockCtrl)

	s.handler = query.NewGetMeHandler(s.mockUserRepository)
}

func (s *getMeTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *getMeTestSuite) Test_1_Success() {
	// mock data
	var id = database.NewStringID()

	s.mockUserRepository.EXPECT().
		FindUserByID(gomock.Any(), gomock.Any()).
		Return(&domain.User{
			ID:   id,
			Name: "Test user",
		}, nil)

	// call
	ctx := appcontext.New(context.Background())
	resp, err := s.handler.GetMe(ctx, id, dto.GetMeRequest{})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), id, resp.User.ID)
}

func (s *getMeTestSuite) Test_2_Fail_InvalidID() {
	// mock data
	s.mockUserRepository.EXPECT().
		FindUserByID(gomock.Any(), gomock.Any()).
		Return(nil, apperrors.User.InvalidUserID)

	// call
	ctx := appcontext.New(context.Background())
	resp, err := s.handler.GetMe(ctx, "invalid id", dto.GetMeRequest{})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.User.InvalidUserID, err)
}

//
// END OF CASES
//

func TestGetMeTestSuite(t *testing.T) {
	suite.Run(t, new(getMeTestSuite))
}
