package infrastructure

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/vocabularypb"
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/dto"
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

func (r VocabularyHub) BookmarkVocabulary(ctx *appcontext.AppContext, userID, vocabularyID string) (bool, error) {
	resp, err := r.client.BookmarkVocabulary(ctx.Context(), &vocabularypb.BookmarkVocabularyRequest{
		UserId:       userID,
		VocabularyId: vocabularyID,
	})
	if err != nil {
		return false, err
	}

	return resp.GetIsBookmarked(), nil
}

func (r VocabularyHub) GetUserBookmarkedVocabularies(ctx *appcontext.AppContext, userID, pageToken string) ([]domain.VocabularyBrief, string, error) {
	resp, err := r.client.GetUserBookmarkedVocabularies(ctx.Context(), &vocabularypb.GetUserBookmarkedVocabulariesRequest{
		UserId:    userID,
		PageToken: pageToken,
	})
	if err != nil {
		return nil, "", err
	}

	var result = make([]domain.VocabularyBrief, 0)
	for _, v := range resp.GetVocabularies() {
		brief, mappingErr := mapping.VocabularyMapper{}.FromGrpcToDomainBrief(v)
		if mappingErr != nil {
			return nil, "", mappingErr
		}
		result = append(result, *brief)
	}

	return result, resp.GetNextPageToken(), nil
}

func (r VocabularyHub) GetWordOfTheDay(ctx *appcontext.AppContext, lang string) (*domain.WordOfTheDay, error) {
	resp, err := r.client.GetWordOfTheDay(ctx.Context(), &vocabularypb.GetWordOfTheDayRequest{
		Lang: lang,
	})
	if err != nil {
		return nil, err
	}

	var (
		vocabularyMapper = mapping.VocabularyMapper{}
		vocabulary, _    = vocabularyMapper.FromGrpcToDomainBrief(resp.GetVocabulary())
	)

	var result = &domain.WordOfTheDay{
		Vocabulary:  *vocabulary,
		Information: dto.ConvertGrpcDataToMultilingual(resp.GetInformation()),
	}

	return result, nil
}

func (r VocabularyHub) GetCommunitySentences(ctx *appcontext.AppContext, userID, vocabularyID, lang, pageToken string) ([]domain.CommunitySentenceBrief, string, error) {
	resp, err := r.client.GetCommunitySentences(ctx.Context(), &vocabularypb.GetCommunitySentencesRequest{
		VocabularyId: vocabularyID,
		UserId:       userID,
		Lang:         lang,
		PageToken:    pageToken,
	})
	if err != nil {
		return nil, "", err
	}

	var (
		result = make([]domain.CommunitySentenceBrief, 0)
		mapper = mapping.CommunitySentenceMapper{}
	)
	for _, s := range resp.GetSentences() {
		sentence, mappingErr := mapper.FromGrpcToDomainBrief(s)
		if mappingErr != nil {
			return nil, "", mappingErr
		}
		result = append(result, *sentence)
	}

	return result, resp.GetNextPageToken(), nil
}

func (r VocabularyHub) GetCommunitySentence(ctx *appcontext.AppContext, userID, sentenceID string) (*domain.CommunitySentence, error) {
	resp, err := r.client.GetCommunitySentence(ctx.Context(), &vocabularypb.GetCommunitySentenceRequest{
		SentenceId: sentenceID,
		UserId:     userID,
	})
	if err != nil {
		return nil, err
	}

	var (
		mapper    = mapping.CommunitySentenceMapper{}
		result, _ = mapper.FromGrpcToDomain(resp.GetSentence())
	)

	return result, nil
}

func (r VocabularyHub) GetUserCommunitySentencesDraft(ctx *appcontext.AppContext, userID, vocabularyID, pageToken string) ([]domain.CommunitySentenceDraft, string, error) {
	resp, err := r.client.GetUserCommunitySentenceDrafts(ctx.Context(), &vocabularypb.GetUserCommunitySentenceDraftsRequest{
		VocabularyId: vocabularyID,
		UserId:       userID,
		PageToken:    pageToken,
	})
	if err != nil {
		return nil, "", err
	}

	var (
		result = make([]domain.CommunitySentenceDraft, 0)
		mapper = mapping.CommunitySentenceDraftMapper{}
	)
	for _, s := range resp.GetSentences() {
		sentence, mappingErr := mapper.FromGrpcToDomain(s)
		if mappingErr != nil {
			return nil, "", mappingErr
		}
		result = append(result, *sentence)
	}

	return result, resp.GetNextPageToken(), nil
}
