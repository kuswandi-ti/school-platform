package errors

type Code string

const (
	CodeUnauthorized          Code = "UNAUTHORIZED"
	CodeForbidden             Code = "FORBIDDEN"
	CodeValidationError       Code = "VALIDATION_ERROR"
	CodeNotFound              Code = "NOT_FOUND"
	CodeConflict              Code = "CONFLICT"
	CodeBusinessRuleViolation Code = "BUSINESS_RULE_VIOLATION"
	CodeApprovalRequired      Code = "APPROVAL_REQUIRED"
	CodeResourceLocked        Code = "RESOURCE_LOCKED"
	CodeRateLimited           Code = "RATE_LIMITED"
	CodeInternalError         Code = "INTERNAL_ERROR"
	CodeServiceUnavailable    Code = "SERVICE_UNAVAILABLE"
)

type AppError struct {
	Code    Code
	Message string
	Details any
}

func (e AppError) Error() string {
	return e.Message
}
