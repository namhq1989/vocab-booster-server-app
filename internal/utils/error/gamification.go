package apperrors

import "errors"

var Gamification = struct {
	InvalidPoint     error
	InvalidLevel     error
	InvalidPointData error
}{
	InvalidPoint:     errors.New("gamification_invalid_point"),
	InvalidLevel:     errors.New("gamification_invalid_level"),
	InvalidPointData: errors.New("gamification_invalid_point_data"),
}
