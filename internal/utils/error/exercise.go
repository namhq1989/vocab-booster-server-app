package apperrors

import "errors"

var Exercise = struct {
	InvalidExerciseID   error
	InvalidPoint        error
	InvalidCollectionID error
}{
	InvalidExerciseID:   errors.New("exercise_invalid_id"),
	InvalidPoint:        errors.New("exercise_invalid_point"),
	InvalidCollectionID: errors.New("exercise_invalid_collection_id"),
}
