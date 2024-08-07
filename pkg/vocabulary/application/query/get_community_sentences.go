package query

import (
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type GetCommunitySentencesHandler struct {
	vocabularyHub domain.VocabularyHub
}

func NewGetCommunitySentencesHandler(vocabularyHub domain.VocabularyHub) GetCommunitySentencesHandler {
	return GetCommunitySentencesHandler{
		vocabularyHub: vocabularyHub,
	}
}

func (h GetCommunitySentencesHandler) GetCommunitySentences(ctx *appcontext.AppContext, performerID, vocabularyID string, lang language.Language, req dto.GetCommunitySentencesRequest) (*dto.GetCommunitySentencesResponse, error) {
	ctx.Logger().Info("[query] new get community sentences request", appcontext.Fields{"performerID": performerID, "vocabularyID": vocabularyID, "lang": lang.String(), "pageToken": req.PageToken})

	ctx.Logger().Text("find sentence via grpc")
	sentences, nextPageToken, err := h.vocabularyHub.GetCommunitySentences(ctx, performerID, vocabularyID, lang.String(), req.PageToken)
	if err != nil {
		ctx.Logger().Error("failed to find sentence via grpc", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("convert response data")
	var result = &dto.GetCommunitySentencesResponse{
		Sentences:     make([]dto.CommunitySentenceBrief, 0),
		NextPageToken: nextPageToken,
	}
	for _, s := range sentences {
		result.Sentences = append(result.Sentences, dto.CommunitySentenceBrief{}.FromDomain(s, lang.String()))
	}

	ctx.Logger().Text("done get community sentences request")
	return result, nil
}
