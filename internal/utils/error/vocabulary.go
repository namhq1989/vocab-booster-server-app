package apperrors

import "errors"

var Vocabulary = struct {
	InvalidVocabularyID error
}{
	InvalidVocabularyID: errors.New("vocabulary_invalid_id"),
}
