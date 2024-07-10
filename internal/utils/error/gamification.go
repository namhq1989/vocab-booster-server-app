package apperrors

import "errors"

var Gamification = struct {
	InvalidPoint error
	InvalidLevel error
}{
	InvalidPoint: errors.New("gamification_invalid_point"),
	InvalidLevel: errors.New("gamification_invalid_level"),
}
