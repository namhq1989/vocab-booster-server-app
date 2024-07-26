package application

import (
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/application/query"
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type (
	Queries interface {
		SearchVocabulary(ctx *appcontext.AppContext, performerID string, lang language.Language, req dto.SearchVocabularyRequest) (*dto.SearchVocabularyResponse, error)
	}
	Instance interface {
		Queries
	}

	appQueryHandler struct {
		query.SearchVocabularyHandler
	}
	Application struct {
		appQueryHandler
	}
)

var _ Instance = (*Application)(nil)

func New(
	vocabularyHub domain.VocabularyHub,
) *Application {
	return &Application{
		appQueryHandler: appQueryHandler{
			SearchVocabularyHandler: query.NewSearchVocabularyHandler(vocabularyHub),
		},
	}
}
