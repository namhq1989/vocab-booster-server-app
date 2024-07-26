package query

import (
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type SearchVocabularyHandler struct {
	vocabularyHub domain.VocabularyHub
}

func NewSearchVocabularyHandler(vocabularyHub domain.VocabularyHub) SearchVocabularyHandler {
	return SearchVocabularyHandler{
		vocabularyHub: vocabularyHub,
	}
}

func (h SearchVocabularyHandler) SearchVocabulary(ctx *appcontext.AppContext, performerID string, lang language.Language, req dto.SearchVocabularyRequest) (*dto.SearchVocabularyResponse, error) {
	ctx.Logger().Info("[query] new search vocabulary request", appcontext.Fields{"performerID": performerID, "lang": lang.String(), "term": req.Term})

	var result = &dto.SearchVocabularyResponse{}

	ctx.Logger().Text("search vocabulary via grpc")
	vocabulary, suggestions, err := h.vocabularyHub.SearchVocabulary(ctx, performerID, req.Term)
	if err != nil {
		ctx.Logger().Error("failed to search vocabulary", err, appcontext.Fields{})
		return nil, err
	}
	if vocabulary == nil {
		ctx.Logger().ErrorText("vocabulary not found, respond")
		result.Suggestions = suggestions
		return result, nil
	}

	ctx.Logger().Text("vocabulary found, convert response data")
	vocabularyData := dto.Vocabulary{}.FromDomain(*vocabulary, lang.String())
	result.Found = true
	result.Vocabulary = &vocabularyData
	result.Suggestions = make([]string, 0)

	ctx.Logger().Text("done search vocabulary request")
	return result, nil
}
