package apperrors

import "errors"

var Vocabulary = struct {
	VocabularyNotFound  error
	InvalidVocabularyID error
	InvalidTerm         error
}{
	VocabularyNotFound:  errors.New("vocabulary_not_found"),
	InvalidVocabularyID: errors.New("vocabulary_invalid_id"),
	InvalidTerm:         errors.New("vocabulary_invalid_term"),
}
