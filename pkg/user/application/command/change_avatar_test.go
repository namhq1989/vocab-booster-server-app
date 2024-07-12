package command_test

import (
	"context"
	"testing"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	mockuser "github.com/namhq1989/vocab-booster-server-app/internal/mock/user"
	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/application/command"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type changeAvatarTestSuite struct {
	suite.Suite
	handler            command.ChangeAvatarHandler
	mockCtrl           *gomock.Controller
	mockUserRepository *mockuser.MockUserRepository
}

func (s *changeAvatarTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *changeAvatarTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockUserRepository = mockuser.NewMockUserRepository(s.mockCtrl)

	s.handler = command.NewChangeAvatarHandler(s.mockUserRepository)
}

func (s *changeAvatarTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *changeAvatarTestSuite) Test_1_Success() {
	// mock data
	var id = database.NewStringID()

	s.mockUserRepository.EXPECT().
		FindUserByID(gomock.Any(), gomock.Any()).
		Return(&domain.User{
			ID:   id,
			Name: "Test user",
		}, nil)

	s.mockUserRepository.EXPECT().
		UpdateUser(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.ChangeAvatar(ctx, id, dto.ChangeAvatarRequest{
		Avatar: "1",
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
}

func (s *changeAvatarTestSuite) Test_2_Fail_InvalidID() {
	s.mockUserRepository.EXPECT().
		FindUserByID(gomock.Any(), gomock.Any()).
		Return(nil, apperrors.User.InvalidUserID)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.ChangeAvatar(ctx, "invalid id", dto.ChangeAvatarRequest{
		Avatar: "1",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.User.InvalidUserID, err)
}

//
// END OF CASES
//

func TestChangeAvatarTestSuite(t *testing.T) {
	suite.Run(t, new(changeAvatarTestSuite))
}
