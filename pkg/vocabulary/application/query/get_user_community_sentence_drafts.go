package query

import (
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type GetUserCommunitySentenceDraftsHandler struct {
	vocabularyHub domain.VocabularyHub
}

func NewGetUserCommunitySentenceDraftsHandler(vocabularyHub domain.VocabularyHub) GetUserCommunitySentenceDraftsHandler {
	return GetUserCommunitySentenceDraftsHandler{
		vocabularyHub: vocabularyHub,
	}
}

func (h GetUserCommunitySentenceDraftsHandler) GetUserCommunitySentenceDrafts(ctx *appcontext.AppContext, performerID string, lang language.Language, req dto.GetUserCommunitySentenceDraftsRequest) (*dto.GetUserCommunitySentenceDraftsResponse, error) {
	ctx.Logger().Info("[query] new get user community sentence drafts request", appcontext.Fields{"performerID": performerID, "vocabularyID": req.VocabularyID, "lang": lang.String(), "pageToken": req.PageToken})

	ctx.Logger().Text("get sentences via grpc")
	sentences, nextPageToken, err := h.vocabularyHub.GetUserCommunitySentencesDraft(ctx, performerID, req.VocabularyID, req.PageToken)
	if err != nil {
		ctx.Logger().Error("failed to get sentences via grpc", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("convert response data")
	var result = &dto.GetUserCommunitySentenceDraftsResponse{
		Sentences:     make([]dto.CommunitySentenceDraft, 0),
		NextPageToken: nextPageToken,
	}
	for _, s := range sentences {
		result.Sentences = append(result.Sentences, dto.CommunitySentenceDraft{}.FromDomain(s, lang.String()))
	}

	ctx.Logger().Text("done get community sentence drafts request")
	return result, nil
}
