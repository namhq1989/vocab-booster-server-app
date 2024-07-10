package command

import (
	"slices"

	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type CreateJourneyHandler struct {
	journeyRepository domain.JourneyRepository
}

func NewCreateJourneyHandler(journeyRepository domain.JourneyRepository) CreateJourneyHandler {
	return CreateJourneyHandler{
		journeyRepository: journeyRepository,
	}
}

func (h CreateJourneyHandler) CreateJourney(ctx *appcontext.AppContext, performerID string, req dto.CreateJourneyRequest) (*dto.CreateJourneyResponse, error) {
	ctx.Logger().Info("[command] new create journey request", appcontext.Fields{"performerID": performerID, "lang": req.Lang})

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
	var index = slices.IndexFunc(journeys, func(journey domain.Journey) bool { return journey.Lang.String() == req.Lang })
	if index >= 0 {
		journey = &journeys[index]
	}

	if journey != nil {
		ctx.Logger().Info("journey already existed", appcontext.Fields{"journeyID": journey.ID})

		if !journey.IsLearning {
			ctx.Logger().Text("this journey is not learning, set isLearning to true")
			_ = journey.SetIsLearning(true)

			ctx.Logger().Text("update journey in db")
			if err = h.journeyRepository.UpdateJourney(ctx, *journey); err != nil {
				ctx.Logger().Error("failed to update journey in db", err, appcontext.Fields{})
				return nil, err
			}
		}

		return &dto.CreateJourneyResponse{ID: journey.ID}, nil
	}

	ctx.Logger().Text("create new journey model")
	journey, err = domain.NewJourney(performerID, req.Lang)
	if err != nil {
		ctx.Logger().Error("failed to create new journey model", err, appcontext.Fields{})
		return nil, err
	}

	if len(journeys) == 0 {
		ctx.Logger().Text("this is the first journey, set isLearning to true")
		_ = journey.SetIsLearning(true)
	}

	ctx.Logger().Text("persist journey in db")
	if err = h.journeyRepository.CreateJourney(ctx, *journey); err != nil {
		ctx.Logger().Error("failed to persist journey in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Info("done create journey request", appcontext.Fields{"journeyID": journey.ID})
	return &dto.CreateJourneyResponse{ID: journey.ID}, nil
}
