package apperrors

import "errors"

var Payment = struct {
	InvalidPaymentID error
	PaymentNotFound  error
}{
	InvalidPaymentID: errors.New("payment_invalid_id"),
	PaymentNotFound:  errors.New("payment_not_found"),
}
