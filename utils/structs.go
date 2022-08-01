package utils

type ErrorResponse struct {
	Reason interface{} `json:"reason"`
}

type CreatedResponse struct {
	ID interface{} `json:"id"`
}
