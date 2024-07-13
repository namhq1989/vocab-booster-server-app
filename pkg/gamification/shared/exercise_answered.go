package shared

import (
	"github.com/namhq1989/vocab-booster-server-app/pkg/gamification/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s Service) ExerciseAnswered(ctx *appcontext.AppContext, point domain.Point, completionTime domain.CompletionTime) error {
	return s.db.Transaction(ctx, func(ssCtx mongo.SessionContext) (interface{}, error) {
		var err error

		if point.Point == 0 {
			ctx.Logger().Text("point is 0, skip creating point document")
		} else {
			ctx.Logger().Text("add point in db")
			if err = s.pointRepository.CreatePoint(ctx, point); err != nil {
				ctx.Logger().Error("failed to add point in db", err, appcontext.Fields{})
				return nil, err
			}
		}

		ctx.Logger().Text("add completion time in db")
		if err = s.completionTimeRepository.CreateCompletionTime(ctx, completionTime); err != nil {
			ctx.Logger().Error("failed to add point in db", err, appcontext.Fields{})
			return nil, err
		}

		ctx.Logger().Text("increase user stats in db")
		if err = s.userStatsRepository.IncreaseUserStats(ctx, point.UserID, point.Point, completionTime.Seconds); err != nil {
			ctx.Logger().Error("failed to add user point in db", err, appcontext.Fields{})
			return nil, err
		}

		return nil, nil
	})
}
