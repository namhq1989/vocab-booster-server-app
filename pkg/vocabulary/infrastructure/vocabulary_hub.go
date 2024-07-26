package infrastructure

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/vocabularypb"
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/infrastructure/mapping"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type VocabularyHub struct {
	client vocabularypb.VocabularyServiceClient
}

func NewVocabularyHub(client vocabularypb.VocabularyServiceClient) VocabularyHub {
	return VocabularyHub{
		client: client,
	}
}

func (r VocabularyHub) SearchVocabulary(ctx *appcontext.AppContext, performerID, term string) (*domain.Vocabulary, []string, error) {
	suggestions := make([]string, 0)

	resp, err := r.client.SearchVocabulary(ctx.Context(), &vocabularypb.SearchVocabularyRequest{
		PerformerId: performerID,
		Term:        term,
	})
	if err != nil {
		return nil, suggestions, err
	}

	if !resp.GetFound() {
		suggestions = resp.GetSuggestions()
		return nil, suggestions, nil
	}

	var (
		mapper = mapping.VocabularyMapper{}
	)

	result, err := mapper.FromGrpcToDomain(resp.GetVocabulary())
	if err != nil {
		return nil, suggestions, err
	}

	return result, suggestions, nil
}
