package apierrors

type DataNotFoundErrorWrapper struct {
	ReturnedStatusCode int
	OriginalError      error
}

// Error the original error message remains as it is for logging reasons etc.
// and the wrapper error message is empty because we don't want the client to see anything
func (err DataNotFoundErrorWrapper) Error() string {
	return ""
}

func (err DataNotFoundErrorWrapper) Unwrap() error {
	return err.OriginalError
}

type UserValidationError struct {
	ReturnedStatusCode int
	OriginalError      error
}

func (err UserValidationError) Error() string {
	return "error validating user"
}

func (err UserValidationError) Unwrap() error {
	return err.OriginalError
}
