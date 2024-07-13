package domain

import "github.com/namhq1989/vocab-booster-utilities/appcontext"

type UserPointRepository interface {
	FindUserPoint(ctx *appcontext.AppContext, userID string) (*UserPoint, error)
	IncreasePoint(ctx *appcontext.AppContext, userID string, point int64) error
}

type UserPoint struct {
	ID     string
	UserID string
	Point  int64
}
