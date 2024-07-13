package worker

import (
	"github.com/namhq1989/vocab-booster-server-app/pkg/gamification/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type VocabularySentenceContributedHandler struct {
	pointRepository          domain.PointRepository
	completionTimeRepository domain.CompletionTimeRepository
	userStatsRepository      domain.UserStatsRepository
	service                  domain.Service
}

func NewVocabularySentenceContributedHandler(
	pointRepository domain.PointRepository,
	completionTimeRepository domain.CompletionTimeRepository,
	userStatsRepository domain.UserStatsRepository,
	service domain.Service,
) VocabularySentenceContributedHandler {
	return VocabularySentenceContributedHandler{
		pointRepository:          pointRepository,
		completionTimeRepository: completionTimeRepository,
		userStatsRepository:      userStatsRepository,
		service:                  service,
	}
}

func (w VocabularySentenceContributedHandler) VocabularySentenceContributed(ctx *appcontext.AppContext, payload domain.QueueVocabularySentenceContributedPoint) error {
	ctx.Logger().Text("new point model")
	point, err := domain.NewPoint(payload.UserID, "", payload.VocabularyID, domain.ActionContributeVocabularySentence.String(), payload.Point)
	if err != nil {
		ctx.Logger().Error("failed to create new point model", err, appcontext.Fields{})
		return err
	}

	ctx.Logger().Text("new completion time model")
	completionTime, err := domain.NewCompletionTime(payload.UserID, domain.ActionContributeVocabularySentence.String(), payload.CompletionTime)
	if err != nil {
		ctx.Logger().Error("failed to create new completion time model", err, appcontext.Fields{})
		return err
	}

	ctx.Logger().Text("persist in db with transaction")
	if err = w.service.ExerciseAnswered(ctx, *point, *completionTime); err != nil {
		ctx.Logger().Error("failed to persist in db with transaction", err, appcontext.Fields{})
		return err
	}

	return nil
}
