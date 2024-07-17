package mapping

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/gamificationpb"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/domain"
)

type UserAggregatedPointMapper struct{}

func (UserAggregatedPointMapper) FromGrpcToDomain(uap *gamificationpb.UserAggregatedPoint) (*domain.UserAggregatedPoint, error) {
	result := domain.UserAggregatedPoint{
		Date:  uap.Date,
		Point: uap.Point,
	}

	return &result, nil
}
