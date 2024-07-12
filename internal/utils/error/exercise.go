package apperrors

import "errors"

var Exercise = struct {
	InvalidExerciseID error
}{
	InvalidExerciseID: errors.New("exercise_invalid_id"),
}
