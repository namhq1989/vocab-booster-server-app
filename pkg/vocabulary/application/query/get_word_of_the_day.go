package query

import (
	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type GetWordOfTheDayHandler struct {
	vocabularyHub domain.VocabularyHub
}

func NewGetWordOfTheDayHandler(vocabularyHub domain.VocabularyHub) GetWordOfTheDayHandler {
	return GetWordOfTheDayHandler{
		vocabularyHub: vocabularyHub,
	}
}

func (h GetWordOfTheDayHandler) GetWordOfTheDay(ctx *appcontext.AppContext, performerID string, lang language.Language, _ dto.GetWordOfTheDayRequest) (*dto.GetWordOfTheDayResponse, error) {
	ctx.Logger().Info("[query] new get word of the day request", appcontext.Fields{"performerID": performerID, "lang": lang.String()})

	ctx.Logger().Text("get word of the day via grpc")
	wotd, err := h.vocabularyHub.GetWordOfTheDay(ctx, lang.String())
	if err != nil {
		ctx.Logger().Error("failed to get word of the day via grpc", err, appcontext.Fields{})
		return nil, err
	}
	if wotd == nil {
		ctx.Logger().ErrorText("word of the day not found")
		return nil, apperrors.Vocabulary.VocabularyNotFound
	}

	ctx.Logger().Text("convert response data")
	result := &dto.GetWordOfTheDayResponse{
		Vocabulary:  dto.VocabularyBrief{}.FromDomain(wotd.Vocabulary),
		Information: wotd.Information.GetLocalized(lang.String()),
	}

	ctx.Logger().Text("done get word of the day request")
	return result, nil
}
