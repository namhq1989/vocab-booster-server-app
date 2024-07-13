package apperrors

import "errors"

var Exercise = struct {
	InvalidExerciseID error
	InvalidPoint      error
}{
	InvalidExerciseID: errors.New("exercise_invalid_id"),
	InvalidPoint:      errors.New("exercise_invalid_point"),
}
