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

type downgradeUserSubscriptionTestSuite struct {
	suite.Suite
	handler                        worker.DowngradeUserSubscriptionHandler
	mockCtrl                       *gomock.Controller
	mockUserSubscriptionRepository *mocksubscription.MockUserSubscriptionRepository
}

func (s *downgradeUserSubscriptionTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *downgradeUserSubscriptionTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockUserSubscriptionRepository = mocksubscription.NewMockUserSubscriptionRepository(s.mockCtrl)

	s.handler = worker.NewDowngradeUserSubscriptionHandler(s.mockUserSubscriptionRepository)
}

func (s *downgradeUserSubscriptionTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *downgradeUserSubscriptionTestSuite) Test_1_Success() {
	// mock data
	s.mockUserSubscriptionRepository.EXPECT().
		UpsertUserSubscription(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewGRPC(context.Background())
	err := s.handler.DowngradeUserSubscription(ctx, domain.QueueDowngradeUserSubscriptionPayload{
		Subscription: domain.UserSubscription{
			ID:      database.NewStringID(),
			UserID:  database.NewStringID(),
			Plan:    domain.PlanPremiumMonthly,
			StartAt: manipulation.NowUTC(),
			EndAt:   manipulation.NowUTC(),
		},
	})

	assert.Nil(s.T(), err)
}

//
// END OF CASES
//

func TestDowngradeUserSubscriptionTestSuite(t *testing.T) {
	suite.Run(t, new(downgradeUserSubscriptionTestSuite))
}
