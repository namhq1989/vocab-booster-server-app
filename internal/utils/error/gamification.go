package apperrors

import "errors"

var Gamification = struct {
	InvalidPoint          error
	InvalidCompletionTime error
	InvalidLevel          error
	InvalidPointData      error
}{
	InvalidPoint:          errors.New("gamification_invalid_point"),
	InvalidCompletionTime: errors.New("gamification_invalid_completion_time"),
	InvalidLevel:          errors.New("gamification_invalid_level"),
	InvalidPointData:      errors.New("gamification_invalid_point_data"),
}
