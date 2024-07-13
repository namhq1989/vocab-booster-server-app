package shared

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	"github.com/namhq1989/vocab-booster-server-app/pkg/gamification/domain"
)

type Service struct {
	db                       database.Operations
	pointRepository          domain.PointRepository
	completionTimeRepository domain.CompletionTimeRepository
	userStatsRepository      domain.UserStatsRepository
}

func NewService(
	db database.Operations,
	pointRepository domain.PointRepository,
	completionTimeRepository domain.CompletionTimeRepository,
	userStatsRepository domain.UserStatsRepository,
) Service {
	return Service{
		db:                       db,
		pointRepository:          pointRepository,
		completionTimeRepository: completionTimeRepository,
		userStatsRepository:      userStatsRepository,
	}
}
