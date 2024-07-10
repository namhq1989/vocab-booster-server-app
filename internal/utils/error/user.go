package apperrors

import "errors"

var User = struct {
	InvalidUserID     error
	UserNotFound      error
	InvalidVisibility error
}{
	InvalidUserID:     errors.New("user_invalid_id"),
	UserNotFound:      errors.New("user_not_found"),
	InvalidVisibility: errors.New("user_invalid_visibility"),
}
