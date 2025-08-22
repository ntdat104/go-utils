package errors

import (
	"fmt"
	"testing"
)

func TestNewError_KnownCodes(t *testing.T) {
	tests := []struct {
		code     ErrorCode
		expected string
	}{
		{INVALID_REQUEST_PAYLOAD, "Invalid request payload"},
		{MISSING_REQUIRED_PARAM, "Missing required parameter"},
		{INVALID_AUTH_TOKEN, "Invalid or expired authentication token"},
		{ACCESS_DENIED, "You do not have permission to perform this action"},
		{RESOURCE_NOT_FOUND, "The requested resource was not found"},
		{VALIDATION_FAILED, "Validation failed for one or more fields"},
		{RATE_LIMIT_EXCEEDED, "Too many requests, please try again later"},
		{INTERNAL_SERVER_ERROR, "Internal server error"},
		{DATABASE_ERROR, "Database operation failed"},
		{CACHE_ERROR, "Cache service error"},
		{SERVICE_UNAVAILABLE, "Service is temporarily unavailable"},
		{TIMEOUT_ERROR, "Request timed out"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("code_%d", tt.code), func(t *testing.T) {
			errResp := NewError(tt.code)
			if errResp.ErrorCode != tt.code {
				t.Errorf("expected code %d, got %d", tt.code, errResp.ErrorCode)
			}
			if errResp.ErrorMessage != tt.expected {
				t.Errorf("expected message %q, got %q", tt.expected, errResp.ErrorMessage)
			}
		})
	}
}

func TestNewError_UnknownCode(t *testing.T) {
	unknownCode := ErrorCode(123456)
	errResp := NewError(unknownCode)

	if errResp.ErrorCode != unknownCode {
		t.Errorf("expected code %d, got %d", unknownCode, errResp.ErrorCode)
	}

	expectedMessage := fmt.Sprintf("Unknown error (%d)", unknownCode)
	if errResp.ErrorMessage != expectedMessage {
		t.Errorf("expected message %q, got %q", expectedMessage, errResp.ErrorMessage)
	}
}
