package common

import "net/http"

type response struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func BadRequestResponse() response {
	return response{
		http.StatusBadRequest,
		"Bad Request",
	}
}

func NotFoundResponse() response {
	return response{
		http.StatusNotFound,
		"Not Found",
	}
}

func InternalServerErrorResponse() response {
	return response{
		http.StatusInternalServerError,
		"Internal Server Error",
	}
}

func ConflictResponse(msg string) response {
	return response{
		http.StatusConflict,
		msg,
	}
}

func UnauthorizedResponse(msg string) response {
	return response{
		http.StatusUnauthorized,
		msg,
	}
}

func ForbiddenResponse() response {
	return response{
		http.StatusForbidden,
		"Forbidden",
	}
}