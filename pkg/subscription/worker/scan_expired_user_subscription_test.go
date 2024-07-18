package worker_test

import (
	"context"
	"testing"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	mocksubscription "github.com/namhq1989/vocab-booster-server-app/internal/mock/subscription"
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/manipulation"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/worker"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type scanExpiredUserSubscriptionTestSuite struct {
	suite.Suite
	handler                        worker.ScanExpiredUserSubscriptionHandler
	mockCtrl                       *gomock.Controller
	mockUserSubscriptionRepository *mocksubscription.MockUserSubscriptionRepository
	mockQueueRepository            *mocksubscription.MockQueueRepository
}

func (s *scanExpiredUserSubscriptionTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *scanExpiredUserSubscriptionTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockUserSubscriptionRepository = mocksubscription.NewMockUserSubscriptionRepository(s.mockCtrl)
	s.mockQueueRepository = mocksubscription.NewMockQueueRepository(s.mockCtrl)

	s.handler = worker.NewScanExpiredUserSubscriptionHandler(s.mockUserSubscriptionRepository, s.mockQueueRepository)
}

func (s *scanExpiredUserSubscriptionTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *scanExpiredUserSubscriptionTestSuite) Test_1_Success() {
	// mock data
	s.mockUserSubscriptionRepository.EXPECT().
		FindExpiredUserSubscriptionsByDate(gomock.Any(), gomock.Any()).
		Return([]domain.UserSubscription{
			{
				ID:      database.NewStringID(),
				UserID:  database.NewStringID(),
				Plan:    domain.PlanPremiumMonthly,
				StartAt: manipulation.Now(),
				EndAt:   manipulation.Now(),
			},
		}, nil)

	s.mockQueueRepository.EXPECT().
		DowngradeUserSubscription(gomock.Any(), gomock.Any()).
		AnyTimes().
		Return(nil)

	// call
	ctx := appcontext.NewGRPC(context.Background())
	err := s.handler.ScanExpiredUserSubscription(ctx, domain.QueueScanExpiredUserSubscription{})

	assert.Nil(s.T(), err)
}

//
// END OF CASES
//

func TestScanExpiredUserSubscriptionTestSuite(t *testing.T) {
	suite.Run(t, new(scanExpiredUserSubscriptionTestSuite))
}
