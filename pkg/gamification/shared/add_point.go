package shared

import (
	"github.com/namhq1989/vocab-booster-server-app/pkg/gamification/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s Service) AddPoint(ctx *appcontext.AppContext, point domain.Point) error {
	return s.db.Transaction(ctx, func(ssCtx mongo.SessionContext) (interface{}, error) {
		ctx.Logger().Text("add point in db")
		err := s.pointRepository.CreatePoint(ctx, point)
		if err != nil {
			ctx.Logger().Error("failed to add point in db", err, appcontext.Fields{})
			return nil, err
		}

		ctx.Logger().Text("increase user point in db")
		err = s.userPointRepository.IncreasePoint(ctx, point.UserID, point.Point)
		if err != nil {
			ctx.Logger().Error("failed to add user point in db", err, appcontext.Fields{})
			return nil, err
		}
		return nil, nil
	})
}
