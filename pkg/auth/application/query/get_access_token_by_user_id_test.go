package query_test

import (
	"context"
	"testing"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	mockjwt "github.com/namhq1989/vocab-booster-server-app/internal/mock/jwt"
	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
	"github.com/namhq1989/vocab-booster-server-app/pkg/auth/application/query"
	"github.com/namhq1989/vocab-booster-server-app/pkg/auth/dto"
	"github.com/namhq1989/vocab-booster-server-app/pkg/auth/infrastructure"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type getAccessTokenByUserIDTestSuite struct {
	suite.Suite
	handler  query.GetAccessTokenByUserIDHandler
	mockCtrl *gomock.Controller
	mockJwt  *mockjwt.MockOperations
}

func (s *getAccessTokenByUserIDTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *getAccessTokenByUserIDTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockJwt = mockjwt.NewMockOperations(s.mockCtrl)

	jwtRepository := infrastructure.NewJwtRepository(s.mockJwt)
	s.handler = query.NewGetAccessTokenByUserIDHandler(jwtRepository)
}

func (s *getAccessTokenByUserIDTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *getAccessTokenByUserIDTestSuite) Test_1_Success() {
	// mock data
	s.mockJwt.EXPECT().
		GenerateAccessToken(gomock.Any(), gomock.Any()).
		Return("access_token", nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.GetAccessTokenByUserID(ctx, dto.GetAccessTokenByUserIDRequest{UserID: database.NewStringID()})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), "access_token", resp.AccessToken)
}

func (s *getAccessTokenByUserIDTestSuite) Test_2_Fail_InvalidID() {
	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.GetAccessTokenByUserID(ctx, dto.GetAccessTokenByUserIDRequest{UserID: "invalid id"})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.User.InvalidUserID, err)
}

//
// END OF CASES
//

func TestGetAccessTokenByUserIDTestSuite(t *testing.T) {
	suite.Run(t, new(getAccessTokenByUserIDTestSuite))
}
