package errors

import "testing"

func TestAppErrorReturnsMessage(t *testing.T) {
	err := AppError{
		Code:    CodeValidationError,
		Message: "validation failed",
	}

	if err.Error() != "validation failed" {
		t.Fatalf("expected error message, got %q", err.Error())
	}
}
