package grpc

import (
	"context"
	"testing"

	"github.com/namhq1989/vocab-booster-server-app/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-server-app/core/error"
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/userpb"
	mockmongo "github.com/namhq1989/vocab-booster-server-app/internal/mock/mongo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type createUserTestSuite struct {
	suite.Suite
	handler     CreateUserHandler
	mockCtrl    *gomock.Controller
	mockUserHub *mockmongo.MockUserHub
}

func (s *createUserTestSuite) SetupSuite() {
	s.setupApplication()
}

func (*createUserTestSuite) AfterTest(_, _ string) {
	// do nothing
}

func (s *createUserTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockUserHub = mockmongo.NewMockUserHub(s.mockCtrl)

	s.handler = NewCreateUserHandler(s.mockUserHub)
}

func (s *createUserTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *createUserTestSuite) Test_1_Success() {
	// mock data
	s.mockUserHub.EXPECT().
		CreateUser(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewGRPC(context.Background())
	resp, err := s.handler.CreateUser(ctx, &userpb.CreateUserRequest{
		Name:  "Test user",
		Email: "test@gmail.com",
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
}

func (s *createUserTestSuite) Test_2_Fail_InvalidEmail() {
	// call
	ctx := appcontext.NewGRPC(context.Background())
	resp, err := s.handler.CreateUser(ctx, &userpb.CreateUserRequest{
		Name:  "Test user",
		Email: "invalid email",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.InvalidEmail, err)
}

//
// END OF CASES
//

func TestCreateUserTestSuite(t *testing.T) {
	suite.Run(t, new(createUserTestSuite))
}
