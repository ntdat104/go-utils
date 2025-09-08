package sha256

import "testing"

// TestCalculateSHA256 tests the CalculateSHA256 function with a set of known inputs and expected outputs.
func TestCalculateSHA256(t *testing.T) {
	// Define a slice of structs to hold our test cases.
	// Each struct contains an input string and the expected SHA-256 hash.
	var tests = []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Test with a simple string",
			input:    "hello world",
			expected: "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9",
		},
		{
			name:     "Test with an empty string",
			input:    "",
			expected: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		},
		{
			name:     "Test with a string containing numbers and symbols",
			input:    "test123!@#$%^&*()_+",
			expected: "c2d33beeb2559d476f4601ae9a4ff74282e8daccd58f623e3010e284e7d6ab19",
		},
		{
			name:     "Test with a different simple string",
			input:    "golang",
			expected: "d754ed9f64ac293b10268157f283ee23256fb32a4f8dedb25c8446ca5bcb0bb3",
		},
	}

	// Loop through each test case defined in the slice.
	for _, tt := range tests {
		// Use a subtest for each test case to provide clearer output in case of failure.
		t.Run(tt.name, func(t *testing.T) {
			// Call the function being tested.
			got := CalculateSHA256(tt.input)

			// Compare the result with the expected hash.
			if got != tt.expected {
				// If they don't match, report an error.
				t.Errorf("CalculateSHA256(%q) = %q; want %q", tt.input, got, tt.expected)
			}
		})
	}
}
