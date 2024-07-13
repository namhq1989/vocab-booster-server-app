package infrastructure

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/gamificationpb"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/domain"
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

func (r GamificationHub) GetUserStats(ctx *appcontext.AppContext, userID string) (*domain.GamificationUserStats, error) {
	resp, err := r.client.GetUserStats(ctx.Context(), &gamificationpb.GetUserStatsRequest{
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}

	return &domain.GamificationUserStats{
		Point:          resp.GetPoint(),
		CompletionTime: int(resp.GetCompletionTime()),
	}, nil
}
