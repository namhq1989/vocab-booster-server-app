package query

import (
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type GetJourneysHandler struct {
	journeyRepository domain.JourneyRepository
}

func NewGetJourneysHandler(journeyRepository domain.JourneyRepository) GetJourneysHandler {
	return GetJourneysHandler{
		journeyRepository: journeyRepository,
	}
}

func (h GetJourneysHandler) GetJourneys(ctx *appcontext.AppContext, performerID string, _ dto.GetJourneysRequest) (*dto.GetJourneysResponse, error) {
	ctx.Logger().Info("[query] new get journeys request", appcontext.Fields{"performerID": performerID})

	ctx.Logger().Text("find user's journeys in db")
	journeys, err := h.journeyRepository.FindJourneysByUserID(ctx, performerID)
	if err != nil {
		ctx.Logger().Error("failed to find journey in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("convert to response")
	result := make([]dto.Journey, 0)
	for _, j := range journeys {
		result = append(result, dto.Journey{}.FromDomain(j))
	}

	ctx.Logger().Text("done get journeys request")
	return &dto.GetJourneysResponse{
		Journeys: result,
	}, nil
}
