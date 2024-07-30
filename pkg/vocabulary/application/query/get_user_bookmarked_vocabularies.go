package query

import (
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type GetUserBookmarkedVocabulariesHandler struct {
	vocabularyHub domain.VocabularyHub
}

func NewGetUserBookmarkedVocabulariesHandler(vocabularyHub domain.VocabularyHub) GetUserBookmarkedVocabulariesHandler {
	return GetUserBookmarkedVocabulariesHandler{
		vocabularyHub: vocabularyHub,
	}
}

func (h GetUserBookmarkedVocabulariesHandler) GetUserBookmarkedVocabularies(ctx *appcontext.AppContext, performerID string, req dto.GetUserBookmarkedVocabulariesRequest) (*dto.GetUserBookmarkedVocabulariesResponse, error) {
	ctx.Logger().Info("[query] new get user bookmarked vocabularies request", appcontext.Fields{"performerID": performerID, "pageToken": req.PageToken})

	ctx.Logger().Text(" get bookmarked vocabularies via grpc")
	vocabularies, nextPageToken, err := h.vocabularyHub.GetUserBookmarkedVocabularies(ctx, performerID, req.PageToken)
	if err != nil {
		ctx.Logger().Error("failed to get bookmarked vocabularies via grpc", err, appcontext.Fields{})
		return nil, err
	}

	totalVocabularies := len(vocabularies)
	if totalVocabularies == 0 {
		ctx.Logger().Text("no vocabularies found")
		return &dto.GetUserBookmarkedVocabulariesResponse{
			Vocabularies:  make([]dto.VocabularyBrief, 0),
			NextPageToken: "",
		}, nil
	}

	ctx.Logger().Text("convert response data")
	var result = make([]dto.VocabularyBrief, 0)
	for _, v := range vocabularies {
		result = append(result, dto.VocabularyBrief{}.FromDomain(v))
	}

	ctx.Logger().Text("done get user bookmarked vocabularies request")
	return &dto.GetUserBookmarkedVocabulariesResponse{
		Vocabularies:  result,
		NextPageToken: nextPageToken,
	}, nil
}
