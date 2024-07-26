package apperrors

import "errors"

var Vocabulary = struct {
	InvalidVocabularyID error
	InvalidTerm         error
}{
	InvalidVocabularyID: errors.New("vocabulary_invalid_id"),
	InvalidTerm:         errors.New("vocabulary_invalid_term"),
}
