package query

import (
	"sync"

	"github.com/namhq1989/vocab-booster-server-app/pkg/user/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type GetStatsHandler struct {
	gamificationHub domain.GamificationHub
	exerciseHub     domain.ExerciseHub
}

func NewGetStatsHandler(gamificationHub domain.GamificationHub, exerciseHub domain.ExerciseHub) GetStatsHandler {
	return GetStatsHandler{
		gamificationHub: gamificationHub,
		exerciseHub:     exerciseHub,
	}
}

func (h GetStatsHandler) GetStats(ctx *appcontext.AppContext, performerID string, _ dto.GetStatsRequest) (*dto.GetStatsResponse, error) {
	ctx.Logger().Info("[query] new get stats request", appcontext.Fields{"userID": performerID})

	var (
		result = &dto.GetStatsResponse{}
		wg     sync.WaitGroup
	)

	wg.Add(2)

	go func() {
		defer wg.Done()
		ctx.Logger().Text("get gamification stats")
		stats, err := h.gamificationHub.GetUserStats(ctx, performerID)
		if err != nil {
			ctx.Logger().Error("failed to get gamification stats", err, appcontext.Fields{})
			return
		}

		result.Point = stats.Point
		result.CompletionTime = stats.CompletionTime
	}()

	go func() {
		defer wg.Done()
		ctx.Logger().Text("get exercise stats")
		stats, err := h.exerciseHub.GetUserStats(ctx, performerID)
		if err != nil {
			ctx.Logger().Error("failed to get gamification stats", err, appcontext.Fields{})
			return
		}

		result.MasteredExercises = stats.Mastered
		result.WaitingForReviewExercises = stats.WaitingForReview
	}()

	wg.Wait()

	ctx.Logger().Text("done get stats request")
	return result, nil
}
