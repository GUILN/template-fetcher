package models

type BoilerplateError struct {
	Message    string
	InnerError error
}

func (be *BoilerplateError) Error() string {
	return be.Message
}
