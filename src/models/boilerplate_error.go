package models

type BoilerplateError struct {
	Message    string
	InnerError error
}

func CreateBoilerplateErrorFromError(err error, message string) *BoilerplateError {
	return &BoilerplateError{Message: message, InnerError: err}
}

func (be *BoilerplateError) Error() string {
	return be.Message
}
