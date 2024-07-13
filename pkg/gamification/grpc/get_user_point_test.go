package grpc_test

import (
	"context"
	"testing"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/gamificationpb"
	mockgamification "github.com/namhq1989/vocab-booster-server-app/internal/mock/gamification"
	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
	"github.com/namhq1989/vocab-booster-server-app/pkg/gamification/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/gamification/grpc"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type getUserPointTestSuite struct {
	suite.Suite
	handler                 grpc.GetUserPointHandler
	mockCtrl                *gomock.Controller
	mockUserPointRepository *mockgamification.MockUserPointRepository
}

func (s *getUserPointTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *getUserPointTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockUserPointRepository = mockgamification.NewMockUserPointRepository(s.mockCtrl)

	s.handler = grpc.NewGetUserPointHandler(s.mockUserPointRepository)
}

func (s *getUserPointTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *getUserPointTestSuite) Test_1_Success() {
	// mock data
	var (
		userID       = database.NewStringID()
		point  int64 = 500
	)

	s.mockUserPointRepository.EXPECT().
		FindUserPoint(gomock.Any(), gomock.Any()).
		Return(&domain.UserPoint{
			ID:     database.NewStringID(),
			UserID: userID,
			Point:  point,
		}, nil)

	// call
	ctx := appcontext.NewGRPC(context.Background())
	resp, err := s.handler.GetUserPoint(ctx, &gamificationpb.GetUserPointRequest{
		UserId: userID,
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), point, resp.GetPoint())
}

func (s *getUserPointTestSuite) Test_2_Fail_InvalidUserID() {
	// mock data
	s.mockUserPointRepository.EXPECT().
		FindUserPoint(gomock.Any(), gomock.Any()).
		Return(nil, apperrors.User.InvalidUserID)

	// call
	ctx := appcontext.NewGRPC(context.Background())
	resp, err := s.handler.GetUserPoint(ctx, &gamificationpb.GetUserPointRequest{
		UserId: "invalid id",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.User.InvalidUserID, err)
}

//
// END OF CASES
//

func TestGetUserPointTestSuite(t *testing.T) {
	suite.Run(t, new(getUserPointTestSuite))
}
