package apperrors

import "errors"

var User = struct {
	InvalidUserID     error
	UserNotFound      error
	InvalidVisibility error
	InvalidJourneyID  error
	JourneyNotFound   error
}{
	InvalidUserID:     errors.New("user_invalid_id"),
	UserNotFound:      errors.New("user_not_found"),
	InvalidVisibility: errors.New("user_invalid_visibility"),
	InvalidJourneyID:  errors.New("user_invalid_journey_id"),
	JourneyNotFound:   errors.New("user_journey_not_found"),
}
