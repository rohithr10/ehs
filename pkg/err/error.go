package err

import "net/http"

type MsPayload struct {
}
type MsMessage struct {
	ResponseCode        int    `json:"responseCode,omitempty"`
	Description         string `json:"description,omitempty"`
	DetailedDescription string `json:"detailedDescription,omitempty"`
	ErrorCatch          []byte `json:"error_catch,omitempty"`
}
type AppError struct {
	Messages MsMessage `json:"messages,omitempty"`
	Payload  MsPayload `json:"payload,omitempty"`
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Messages: MsMessage{
			ResponseCode:        e.Messages.ResponseCode,
			Description:         e.Messages.Description,
			DetailedDescription: e.Messages.DetailedDescription,
		},
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Messages: MsMessage{
			ResponseCode:        http.StatusNotFound,
			Description:         "Failed",
			DetailedDescription: message,
		},
	}
}

func MissingDetails(message string) *AppError {
	return &AppError{
		Messages: MsMessage{
			ResponseCode:        406,
			Description:         "Failed",
			DetailedDescription: message,
		},
	}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Messages: MsMessage{
			ResponseCode:        http.StatusInternalServerError,
			Description:         "Failed",
			DetailedDescription: message,
		},
	}
}
func NewUnAuthorizedError(message string) *AppError {
	return &AppError{
		Messages: MsMessage{
			ResponseCode:        http.StatusUnauthorized,
			Description:         "Failed",
			DetailedDescription: message,
		},
	}
}
func NewValidationError(message string) *AppError {
	return &AppError{
		Messages: MsMessage{
			ResponseCode:        http.StatusUnprocessableEntity,
			Description:         "Failed",
			DetailedDescription: message,
		},
	}
}

func NewHttpRequestError(message string) *AppError {
	return &AppError{
		Messages: MsMessage{
			ResponseCode:        http.StatusInternalServerError,
			Description:         "Failed",
			DetailedDescription: message,
		},
	}
}
func NewBadRequest(message string) *AppError {
	return &AppError{
		Messages: MsMessage{
			ResponseCode:        http.StatusBadRequest,
			Description:         "Failed",
			DetailedDescription: message,
		},
	}
}
func NewFound(message string) *AppError {
	return &AppError{
		Messages: MsMessage{
			ResponseCode:        http.StatusOK,
			Description:         "Success",
			DetailedDescription: message,
		},
	}
}

func NewNotFound(message string) *AppError {
	return &AppError{
		Messages: MsMessage{
			ResponseCode:        http.StatusBadRequest,
			Description:         "Failed",
			DetailedDescription: message,
		},
	}
}

func NewCustomErrorCodeFromThirdPartyAPI(httpCode int, message string) *AppError {
	return &AppError{
		Messages: MsMessage{
			ResponseCode:        httpCode,
			Description:         "Failed",
			DetailedDescription: message,
		},
	}
}

func NewCustomErrorCodeFromThirdPartyCatch(err []byte) *AppError {
	return &AppError{
		Messages: MsMessage{
			ErrorCatch: err,
		},
	}
}
