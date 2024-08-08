package application

import (
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/application/command"
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/application/query"
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type (
	Commands interface {
		BookmarkVocabulary(ctx *appcontext.AppContext, performerID, vocabularyID string, _ dto.BookmarkVocabularyRequest) (*dto.BookmarkVocabularyResponse, error)
	}
	Queries interface {
		SearchVocabulary(ctx *appcontext.AppContext, performerID string, lang language.Language, req dto.SearchVocabularyRequest) (*dto.SearchVocabularyResponse, error)
		GetUserBookmarkedVocabularies(ctx *appcontext.AppContext, performerID string, req dto.GetUserBookmarkedVocabulariesRequest) (*dto.GetUserBookmarkedVocabulariesResponse, error)
		GetWordOfTheDay(ctx *appcontext.AppContext, performerID string, lang language.Language, _ dto.GetWordOfTheDayRequest) (*dto.GetWordOfTheDayResponse, error)
		GetCommunitySentences(ctx *appcontext.AppContext, performerID, vocabularyID string, lang language.Language, req dto.GetCommunitySentencesRequest) (*dto.GetCommunitySentencesResponse, error)
		GetCommunitySentence(ctx *appcontext.AppContext, userID, sentenceID string, lang language.Language, _ dto.GetCommunitySentenceRequest) (*dto.GetCommunitySentenceResponse, error)
		GetUserCommunitySentenceDrafts(ctx *appcontext.AppContext, performerID string, lang language.Language, req dto.GetUserCommunitySentenceDraftsRequest) (*dto.GetUserCommunitySentenceDraftsResponse, error)
	}
	Instance interface {
		Commands
		Queries
	}

	appCommandHandler struct {
		command.BookmarkVocabularyHandler
	}
	appQueryHandler struct {
		query.SearchVocabularyHandler
		query.GetUserBookmarkedVocabulariesHandler
		query.GetWordOfTheDayHandler
		query.GetCommunitySentencesHandler
		query.GetCommunitySentenceHandler
		query.GetUserCommunitySentenceDraftsHandler
	}
	Application struct {
		appCommandHandler
		appQueryHandler
	}
)

var _ Instance = (*Application)(nil)

func New(
	vocabularyHub domain.VocabularyHub,
) *Application {
	return &Application{
		appCommandHandler: appCommandHandler{
			BookmarkVocabularyHandler: command.NewBookmarkVocabularyHandler(vocabularyHub),
		},
		appQueryHandler: appQueryHandler{
			SearchVocabularyHandler:               query.NewSearchVocabularyHandler(vocabularyHub),
			GetUserBookmarkedVocabulariesHandler:  query.NewGetUserBookmarkedVocabulariesHandler(vocabularyHub),
			GetWordOfTheDayHandler:                query.NewGetWordOfTheDayHandler(vocabularyHub),
			GetCommunitySentencesHandler:          query.NewGetCommunitySentencesHandler(vocabularyHub),
			GetCommunitySentenceHandler:           query.NewGetCommunitySentenceHandler(vocabularyHub),
			GetUserCommunitySentenceDraftsHandler: query.NewGetUserCommunitySentenceDraftsHandler(vocabularyHub),
		},
	}
}
