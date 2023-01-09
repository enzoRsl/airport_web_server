package errorUtils

type ResponseError struct {
	ErrorMessage string
	Code         int
}

func NewResponseError(ErrorMessage string, code int) *ResponseError {
	reponse := ResponseError{ErrorMessage: ErrorMessage, Code: code}
	return &reponse
}
