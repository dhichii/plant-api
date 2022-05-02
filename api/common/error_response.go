package common

import "net/http"

type response struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func BadRequestResponse() response {
	return response{
		http.StatusBadRequest,
		"Bad request",
	}
}

func NotFoundResponse() response {
	return response{
		http.StatusNotFound,
		"Not found",
	}
}

func InternalServerErrorResponse() response {
	return response{
		http.StatusInternalServerError,
		"Internal server error",
	}
}