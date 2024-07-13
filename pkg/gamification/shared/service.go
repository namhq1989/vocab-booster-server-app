package shared

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	"github.com/namhq1989/vocab-booster-server-app/pkg/gamification/domain"
)

type Service struct {
	db                  database.Operations
	pointRepository     domain.PointRepository
	userPointRepository domain.UserPointRepository
}

func NewService(
	db database.Operations,
	pointRepository domain.PointRepository,
	userPointRepository domain.UserPointRepository,
) Service {
	return Service{
		db:                  db,
		pointRepository:     pointRepository,
		userPointRepository: userPointRepository,
	}
}
