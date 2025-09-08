package b64

import "encoding/base64"

// EncodeBase64 encodes a string into its Base64 representation.
// It takes a string as input and returns the Base64-encoded string.
func EncodeBase64(input string) string {
	// Convert the input string to a byte slice.
	data := []byte(input)

	// Encode the byte slice using the standard Base64 encoding.
	encodedString := base64.StdEncoding.EncodeToString(data)

	// Return the encoded string.
	return encodedString
}

// DecodeBase64 decodes a Base64-encoded string back to its original string representation.
// It takes a Base64-encoded string as input and returns the decoded string and an error,
// which will be non-nil if the input is not a valid Base64 string.
func DecodeBase64(input string) (string, error) {
	// Decode the Base64 string. The result is a byte slice.
	decodedBytes, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		// Return the error if decoding fails.
		return "", err
	}

	// Convert the decoded byte slice back to a string and return it.
	return string(decodedBytes), nil
}
