package errors

import "fmt"

type ErrorCode int

const (
	// --- 4xx: Client Errors ---
	INVALID_REQUEST_PAYLOAD ErrorCode = 400001
	MISSING_REQUIRED_PARAM  ErrorCode = 400002
	INVALID_AUTH_TOKEN      ErrorCode = 400003
	ACCESS_DENIED           ErrorCode = 400004
	RESOURCE_NOT_FOUND      ErrorCode = 400005
	VALIDATION_FAILED       ErrorCode = 400006
	RATE_LIMIT_EXCEEDED     ErrorCode = 400007

	// --- 5xx: Server Errors ---
	INTERNAL_SERVER_ERROR ErrorCode = 500001
	DATABASE_ERROR        ErrorCode = 500002
	CACHE_ERROR           ErrorCode = 500003
	SERVICE_UNAVAILABLE   ErrorCode = 500004
	TIMEOUT_ERROR         ErrorCode = 500005
)

var errorMessages = map[ErrorCode]string{
	// 4xx
	INVALID_REQUEST_PAYLOAD: "Invalid request payload",
	MISSING_REQUIRED_PARAM:  "Missing required parameter",
	INVALID_AUTH_TOKEN:      "Invalid or expired authentication token",
	ACCESS_DENIED:           "You do not have permission to perform this action",
	RESOURCE_NOT_FOUND:      "The requested resource was not found",
	VALIDATION_FAILED:       "Validation failed for one or more fields",
	RATE_LIMIT_EXCEEDED:     "Too many requests, please try again later",

	// 5xx
	INTERNAL_SERVER_ERROR: "Internal server error",
	DATABASE_ERROR:        "Database operation failed",
	CACHE_ERROR:           "Cache service error",
	SERVICE_UNAVAILABLE:   "Service is temporarily unavailable",
	TIMEOUT_ERROR:         "Request timed out",
}

type ErrorResponse struct {
	ErrorCode    ErrorCode `json:"error_code"`
	ErrorMessage string    `json:"error_message"`
}

func NewError(code ErrorCode) ErrorResponse {
	msg, ok := errorMessages[code]
	if !ok {
		msg = fmt.Sprintf("Unknown error (%d)", code)
	}

	return ErrorResponse{
		ErrorCode:    code,
		ErrorMessage: msg,
	}
}
