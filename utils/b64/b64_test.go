package b64

import "testing"

// TestEncodeBase64 tests the EncodeBase64 function with a variety of inputs.
func TestEncodeBase64(t *testing.T) {
	// Define a slice of structs to hold our test cases.
	var tests = []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Simple string",
			input:    "hello world",
			expected: "aGVsbG8gd29ybGQ=",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "String with special characters",
			input:    "test!@#$%",
			expected: "dGVzdCFAIyQl",
		},
		{
			name:     "String with numbers",
			input:    "1234567890",
			expected: "MTIzNDU2Nzg5MA==",
		},
	}

	// Iterate over each test case.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EncodeBase64(tt.input)
			if got != tt.expected {
				t.Errorf("EncodeBase64(%q) = %q; want %q", tt.input, got, tt.expected)
			}
		})
	}
}

// TestDecodeBase64 tests the DecodeBase64 function, including error handling for invalid input.
func TestDecodeBase64(t *testing.T) {
	// Define a slice of structs for decode test cases.
	// `input` is the Base64 string, `expected` is the decoded string, and `expectError` indicates if an error is expected.
	var tests = []struct {
		name        string
		input       string
		expected    string
		expectError bool
	}{
		{
			name:        "Valid Base64 string",
			input:       "aGVsbG8gd29ybGQ=",
			expected:    "hello world",
			expectError: false,
		},
		{
			name:        "Empty string",
			input:       "",
			expected:    "",
			expectError: false,
		},
		{
			name:        "Invalid Base64 string (contains an invalid character)",
			input:       "aGVsbG8gd29ybGQ!@#",
			expected:    "",
			expectError: true,
		},
	}

	// Iterate over each test case.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodeBase64(tt.input)

			// Check if we expect an error but didn't get one.
			if tt.expectError && err == nil {
				t.Errorf("DecodeBase64(%q) did not return an error, but one was expected", tt.input)
				return
			}

			// Check if we didn't expect an error but got one.
			if !tt.expectError && err != nil {
				t.Errorf("DecodeBase64(%q) returned an unexpected error: %v", tt.input, err)
				return
			}

			// If no error was expected, compare the returned value with the expected value.
			if !tt.expectError && got != tt.expected {
				t.Errorf("DecodeBase64(%q) = %q; want %q", tt.input, got, tt.expected)
			}
		})
	}
}
