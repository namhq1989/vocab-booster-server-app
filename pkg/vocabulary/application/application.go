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
			SearchVocabularyHandler:              query.NewSearchVocabularyHandler(vocabularyHub),
			GetUserBookmarkedVocabulariesHandler: query.NewGetUserBookmarkedVocabulariesHandler(vocabularyHub),
		},
	}
}
