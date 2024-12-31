package utils

type HttpError struct {
	StatusCode int
	Message    string
	Reason     error
}

func NewHttpError(statusCode int, message string, reason error) *HttpError {
	return &HttpError{statusCode, message, reason}
}

func (e *HttpError) Error() string {
	return e.Reason.Error()
}

type UniqueFieldError struct {
	*HttpError
	Field string
}

func NewUniqueFieldError(field string, message string, reason error) *UniqueFieldError {
	return &UniqueFieldError{
		HttpError: NewHttpError(409, message, reason),
		Field:     field,
	}
}

func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
