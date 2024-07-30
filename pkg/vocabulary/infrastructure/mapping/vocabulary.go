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
		Examples:      make([]domain.VocabularyExample, 0),
	}

	for _, def := range vocab.GetDefinitions() {
		result.Definitions = append(result.Definitions, domain.VocabularyDefinition{
			Pos:        def.GetPos(),
			Definition: dto.ConvertGrpcDataToMultilingual(def.GetDefinition()),
		})
	}

	for _, example := range vocab.GetExamples() {
		result.Examples = append(result.Examples, domain.VocabularyExample{
			ID:      example.GetId(),
			Audio:   example.GetAudio(),
			Content: dto.ConvertGrpcDataToMultilingual(example.GetContent()),
			MainWord: domain.VocabularyMainWord{
				Word: example.GetMainWord().GetWord(),
				Base: example.GetMainWord().GetBase(),
				Pos:  example.GetMainWord().GetPos(),
			},
		})
	}

	return &result, nil
}

func (VocabularyMapper) FromGrpcToDomainBrief(vocab *vocabularypb.VocabularyBrief) (*domain.VocabularyBrief, error) {
	result := domain.VocabularyBrief{
		ID:            vocab.GetId(),
		Term:          vocab.GetTerm(),
		PartsOfSpeech: vocab.GetPartsOfSpeech(),
		Ipa:           vocab.GetIpa(),
		Audio:         vocab.GetAudio(),
	}

	return &result, nil
}
