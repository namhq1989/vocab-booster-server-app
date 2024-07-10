package query_test

import (
	"context"
	"testing"

	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/application/query"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type getSubscriptionPlansSuite struct {
	suite.Suite
	handler query.GetSubscriptionPlansHandler
}

func (s *getSubscriptionPlansSuite) SetupSuite() {
	s.setupApplication()
}

func (s *getSubscriptionPlansSuite) setupApplication() {
	s.handler = query.NewGetSubscriptionPlansHandler()
}

func (*getSubscriptionPlansSuite) TearDownTest() {
	// do nothing
}

//
// CASES
//

func (s *getSubscriptionPlansSuite) Test_1_Success() {
	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.GetSubscriptionPlans(ctx, "", dto.GetSubscriptionPlansRequest{})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), 4, len(resp.Plans))
}

//
// END OF CASES
//

func TestGetSubscriptionPlansSuite(t *testing.T) {
	suite.Run(t, new(getSubscriptionPlansSuite))
}
