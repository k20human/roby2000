package handler

type successResponse struct {
	Message string `json:"message"`
}

type errorResponse struct {
	Error   error  `json:"error"`
	Message string `json:"message"`
}
