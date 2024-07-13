package grpc

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/gamificationpb"
	"github.com/namhq1989/vocab-booster-server-app/pkg/gamification/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type GetUserStatsHandler struct {
	userStatsRepository domain.UserStatsRepository
}

func NewGetUserStatsHandler(userStatsRepository domain.UserStatsRepository) GetUserStatsHandler {
	return GetUserStatsHandler{
		userStatsRepository: userStatsRepository,
	}
}

func (h GetUserStatsHandler) GetUserStats(ctx *appcontext.AppContext, req *gamificationpb.GetUserStatsRequest) (*gamificationpb.GetUserStatsResponse, error) {
	ctx.Logger().Info("[hub] new get user stats request", appcontext.Fields{"userID": req.GetUserId()})

	ctx.Logger().Text("find in db")
	up, err := h.userStatsRepository.FindUserStats(ctx, req.GetUserId())
	if err != nil {
		ctx.Logger().Error("failed to find in db", err, appcontext.Fields{})
		return nil, err
	}

	var (
		point          int64 = 0
		completionTime       = 0
	)
	if up != nil {
		point = up.Point
		completionTime = up.CompletionTime
	}

	ctx.Logger().Text("done get user point request")
	return &gamificationpb.GetUserStatsResponse{
		Point:          point,
		CompletionTime: int32(completionTime),
	}, nil
}
