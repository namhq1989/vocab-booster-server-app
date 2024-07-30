package command

import (
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type BookmarkVocabularyHandler struct {
	vocabularyHub domain.VocabularyHub
}

func NewBookmarkVocabularyHandler(vocabularyHub domain.VocabularyHub) BookmarkVocabularyHandler {
	return BookmarkVocabularyHandler{
		vocabularyHub: vocabularyHub,
	}
}

func (h BookmarkVocabularyHandler) BookmarkVocabulary(ctx *appcontext.AppContext, performerID, vocabularyID string, _ dto.BookmarkVocabularyRequest) (*dto.BookmarkVocabularyResponse, error) {
	ctx.Logger().Info("[command] new bookmark vocabulary request", appcontext.Fields{"performerID": performerID, "vocabularyID": vocabularyID})

	ctx.Logger().Text("bookmark vocabulary via grpc")
	isBookmarked, err := h.vocabularyHub.BookmarkVocabulary(ctx, performerID, vocabularyID)
	if err != nil {
		ctx.Logger().Error("failed to bookmark vocabulary via grpc", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("done bookmark vocabulary request")
	return &dto.BookmarkVocabularyResponse{
		IsBookmarked: isBookmarked,
	}, nil
}
