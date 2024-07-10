package command

import (
	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type SwitchJourneyHandler struct {
	journeyRepository domain.JourneyRepository
}

func NewSwitchJourneyHandler(journeyRepository domain.JourneyRepository) SwitchJourneyHandler {
	return SwitchJourneyHandler{
		journeyRepository: journeyRepository,
	}
}

func (h SwitchJourneyHandler) SwitchJourney(ctx *appcontext.AppContext, performerID string, req dto.SwitchJourneyRequest) (*dto.SwitchJourneyResponse, error) {
	ctx.Logger().Info("[command] new switch journey request", appcontext.Fields{"performerID": performerID, "lang": req.Lang})

	if req.Lang == language.English.String() {
		ctx.Logger().ErrorText("cannot create a journey with lang English")
		return nil, apperrors.Common.InvalidLanguage
	}

	ctx.Logger().Text("find user's journeys in db")
	journeys, err := h.journeyRepository.FindJourneysByUserID(ctx, performerID)
	if err != nil {
		ctx.Logger().Error("failed to find journey in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("find journey with requested lang in user's journeys")
	var journey *domain.Journey
	for _, j := range journeys {
		if j.Lang.String() == req.Lang {
			journey = &j
			break
		}
	}

	if journey == nil {
		ctx.Logger().ErrorText("journey not found, respond")
		return nil, apperrors.User.JourneyNotFound
	}
	if journey.IsLearning {
		ctx.Logger().Text("this journey is learning, respond")
		return &dto.SwitchJourneyResponse{ID: journey.ID}, nil
	}

	ctx.Logger().Text("find current is learning journey")
	var currentLearningJourney *domain.Journey
	for _, j := range journeys {
		if j.IsLearning {
			currentLearningJourney = &j
			break
		}
	}
	if currentLearningJourney != nil {
		ctx.Logger().Text("set current learning journey isLearning to false")
		_ = currentLearningJourney.SetIsLearning(false)

		ctx.Logger().Text("update current learning journey in db")
		if err = h.journeyRepository.UpdateJourney(ctx, *currentLearningJourney); err != nil {
			ctx.Logger().Error("failed to update journey in db", err, appcontext.Fields{})
			return nil, err
		}
	}

	ctx.Logger().Info("set journey isLearning to true", appcontext.Fields{"journeyID": journey.ID})
	_ = journey.SetIsLearning(true)

	ctx.Logger().Text("update journey in db")
	if err = h.journeyRepository.UpdateJourney(ctx, *journey); err != nil {
		ctx.Logger().Error("failed to update journey in db", err, appcontext.Fields{})
		return nil, err
	}

	return &dto.SwitchJourneyResponse{ID: journey.ID}, nil
}
