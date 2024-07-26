package mapping

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/vocabularypb"
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/dto"
)

type VocabularyMapper struct{}

func (VocabularyMapper) FromGrpcToDomain(vocab *vocabularypb.Vocabulary) (*domain.Vocabulary, error) {
	result := domain.Vocabulary{
		ID:            vocab.GetId(),
		AuthorID:      vocab.GetAuthorId(),
		Term:          vocab.GetTerm(),
		Definitions:   make([]domain.VocabularyDefinition, 0),
		PartsOfSpeech: vocab.GetPartsOfSpeech(),
		Ipa:           vocab.GetIpa(),
		Audio:         vocab.GetAudio(),
		Synonyms:      vocab.GetSynonyms(),
		Antonyms:      vocab.GetAntonyms(),
	}

	for _, def := range vocab.GetDefinitions() {
		result.Definitions = append(result.Definitions, domain.VocabularyDefinition{
			Pos:        def.GetPos(),
			Definition: dto.ConvertGrpcDataToMultilingual(def.GetDefinition()),
		})
	}

	return &result, nil
}
