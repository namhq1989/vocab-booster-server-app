package worker

import (
	"github.com/namhq1989/vocab-booster-server-app/pkg/gamification/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type AddContributeVocabularySentencePointHandler struct {
	pointRepository     domain.PointRepository
	userPointRepository domain.UserPointRepository
	service             domain.Service
}

func NewAddContributeVocabularySentencePointHandler(
	pointRepository domain.PointRepository,
	userPointRepository domain.UserPointRepository,
	service domain.Service,
) AddContributeVocabularySentencePointHandler {
	return AddContributeVocabularySentencePointHandler{
		pointRepository:     pointRepository,
		userPointRepository: userPointRepository,
		service:             service,
	}
}

func (w AddContributeVocabularySentencePointHandler) AddContributeVocabularySentencePoint(ctx *appcontext.AppContext, payload domain.QueueAddContributeVocabularySentencePoint) error {
	ctx.Logger().Text("new point model")
	point, err := domain.NewPoint(payload.UserID, "", payload.VocabularyID, domain.ActionContributeVocabularySentence.String(), payload.Point)
	if err != nil {
		ctx.Logger().Error("failed to create new point model", err, appcontext.Fields{})
		return err
	}

	ctx.Logger().Text("persist in db with transaction")
	if err = w.service.AddPoint(ctx, *point); err != nil {
		ctx.Logger().Error("failed to persist in db with transaction", err, appcontext.Fields{})
		return err
	}

	return nil
}
