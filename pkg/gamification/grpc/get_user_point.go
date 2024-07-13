package grpc

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/gamificationpb"
	"github.com/namhq1989/vocab-booster-server-app/pkg/gamification/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type GetUserPointHandler struct {
	userPointRepository domain.UserPointRepository
}

func NewGetUserPointHandler(userPointRepository domain.UserPointRepository) GetUserPointHandler {
	return GetUserPointHandler{
		userPointRepository: userPointRepository,
	}
}

func (h GetUserPointHandler) GetUserPoint(ctx *appcontext.AppContext, req *gamificationpb.GetUserPointRequest) (*gamificationpb.GetUserPointResponse, error) {
	ctx.Logger().Info("[hub] new get user point request", appcontext.Fields{"userID": req.GetUserId()})

	ctx.Logger().Text("find in db")
	up, err := h.userPointRepository.FindUserPoint(ctx, req.GetUserId())
	if err != nil {
		ctx.Logger().Error("failed to find in db", err, appcontext.Fields{})
		return nil, err
	}

	var point int64 = 0
	if up != nil {
		point = up.Point
	}

	ctx.Logger().Text("done get user point request")
	return &gamificationpb.GetUserPointResponse{
		Point: point,
	}, nil
}
