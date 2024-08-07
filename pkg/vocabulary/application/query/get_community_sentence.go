package query

import (
	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type GetCommunitySentenceHandler struct {
	vocabularyHub domain.VocabularyHub
}

func NewGetCommunitySentenceHandler(vocabularyHub domain.VocabularyHub) GetCommunitySentenceHandler {
	return GetCommunitySentenceHandler{
		vocabularyHub: vocabularyHub,
	}
}

func (h GetCommunitySentenceHandler) GetCommunitySentence(ctx *appcontext.AppContext, userID, sentenceID string, lang language.Language, _ dto.GetCommunitySentenceRequest) (*dto.GetCommunitySentenceResponse, error) {
	ctx.Logger().Info("[query] new get community sentence request", appcontext.Fields{"userID": userID, "sentenceID": sentenceID, "lang": lang.String()})

	ctx.Logger().Text("find sentence via grpc")
	sentence, err := h.vocabularyHub.GetCommunitySentence(ctx, userID, sentenceID)
	if err != nil {
		ctx.Logger().Error("failed to find sentence via grpc", err, appcontext.Fields{})
		return nil, err
	}
	if sentence == nil {
		ctx.Logger().ErrorText("community sentence not found")
		return nil, apperrors.Vocabulary.CommunitySentenceNotFound
	}

	ctx.Logger().Text("convert response data")
	result := &dto.GetCommunitySentenceResponse{
		Sentence: dto.CommunitySentence{}.FromDomain(*sentence, lang.String()),
	}

	ctx.Logger().Text("done get community sentence request")
	return result, nil
}
