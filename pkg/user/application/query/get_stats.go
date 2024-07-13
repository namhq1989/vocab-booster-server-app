package query

import (
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type GetStatsHandler struct {
	gamificationHub domain.GamificationHub
}

func NewGetStatsHandler(gamificationHub domain.GamificationHub) GetStatsHandler {
	return GetStatsHandler{
		gamificationHub: gamificationHub,
	}
}

func (h GetStatsHandler) GetStats(ctx *appcontext.AppContext, performerID string, _ dto.GetStatsRequest) (*dto.GetStatsResponse, error) {
	ctx.Logger().Info("[query] new get stats request", appcontext.Fields{"userID": performerID})

	ctx.Logger().Text("get gamification stats")
	gamificationStats, err := h.gamificationHub.GetUserStats(ctx, performerID)
	if err != nil {
		ctx.Logger().Error("failed to get gamification stats", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("done get stats request")
	return &dto.GetStatsResponse{
		Point:          gamificationStats.Point,
		CompletionTime: gamificationStats.CompletionTime,
	}, nil
}
