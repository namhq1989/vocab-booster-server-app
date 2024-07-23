package infrastructure

import (
	"time"

	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/gamificationpb"
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/manipulation"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/infrastructure/mapping"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type GamificationHub struct {
	client gamificationpb.GamificationServiceClient
}

func NewGamificationHub(client gamificationpb.GamificationServiceClient) GamificationHub {
	return GamificationHub{
		client: client,
	}
}

func (r GamificationHub) GetUserRecentPointsChart(ctx *appcontext.AppContext, userID, timezone string, from, to time.Time) ([]domain.UserAggregatedPoint, error) {
	resp, err := r.client.GetUserRecentPointsChart(ctx.Context(), &gamificationpb.GetUserRecentPointsChartRequest{
		UserId:   userID,
		Timezone: timezone,
		From:     manipulation.ConvertToProtoTimestamp(from),
		To:       manipulation.ConvertToProtoTimestamp(to),
	})
	if err != nil {
		return nil, err
	}

	var (
		result = make([]domain.UserAggregatedPoint, 0)
		mapper = mapping.UserAggregatedPointMapper{}
	)

	for _, p := range resp.GetPoints() {
		point, _ := mapper.FromGrpcToDomain(p)
		if point != nil {
			result = append(result, *point)

		}
	}

	return result, nil
}
