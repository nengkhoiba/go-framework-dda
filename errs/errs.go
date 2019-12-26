package errs

import "errors"

var (
	InternalServerError = errors.New("Opps, something went wrong. Please try again.")
	BadRequest          = errors.New("Bad request.")
	Unauthorized        = errors.New("You are not authorized to perform this action.")
	NoHeader            = errors.New("Authorization header required.")
	NoBearer            = errors.New("Authorization requires Bearer scheme.")
	UnintendedExecution = errors.New("This should not be executing.")
)
