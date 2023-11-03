package ez

import (
	"bytes"
	"fmt"
)

type ErrorCode int

// Application error codes
const (
	ErrorCodeOk                 = 0  //Not an error; returned on success.
	ErrorCodeCancelled          = 1  //The operation was cancelled, typically by the caller.
	ErrorCodeUnknown            = 2  //Unknown error. For example, this error may be returned when a Status value received from another address space belongs to an error space that is not known in this address space. Also errors raised by APIs that do not return enough error information may be converted to this error.
	ErrorCodeInvalidArgument    = 3  // validation failed
	ErrorCodeDeadlineExceeded   = 4  // deadline exceeded
	ErrorCodeNotFound           = 5  // entity does not exist
	ErrorCodeConflict           = 6  // action cannot be performed
	ErrorCodeNotAuthorized      = 7  // requester does not have permissions to perform action
	ErrorCodeResourceExhausted  = 8  // the resource has been exhausted
	ErrorCodeFailedPrecondition = 9  // operation was rejected because the system is not in a state required for the operation's execution
	ErrorCodeAborted            = 10 // operation was aborted
	ErrorCodeOutOfRange         = 11 // operation was attempted past the valid range
	ErrorCodeUnimplemented      = 12 // the operation has not been implemented
	ErrorCodeInternal           = 13 // internal error
	ErrorCodeUnavailable        = 14 // the system or operation is not available
	ErrorCodeDataLoss           = 15 // unrecoverable data loss or corruption
	ErrorCodeUnauthenticated    = 16 // requester is not authenticated
)

func (code ErrorCode) String() string {
	switch code {
	case ErrorCodeOk:
		return "ok"
	case ErrorCodeCancelled:
		return "cancelled"
	case ErrorCodeUnknown:
		return "unknown"
	case ErrorCodeInvalidArgument:
		return "invalid_argument"
	case ErrorCodeDeadlineExceeded:
		return "deadline_exceeded"
	case ErrorCodeNotFound:
		return "not_found"
	case ErrorCodeConflict:
		return "conflict"
	case ErrorCodeNotAuthorized:
		return "not_authorized"
	case ErrorCodeResourceExhausted:
		return "resource_exhausted"
	case ErrorCodeFailedPrecondition:
		return "failed_pre_condition"
	case ErrorCodeAborted:
		return "aborted"
	case ErrorCodeOutOfRange:
		return "out_of_range"
	case ErrorCodeUnimplemented:
		return "unimplemented"
	case ErrorCodeInternal:
		return "internal"
	case ErrorCodeUnavailable:
		return "unavailable"
	case ErrorCodeDataLoss:
		return "data_loss"
	case ErrorCodeUnauthenticated:
		return "not_authenticated"
	}
	return "unspecified"
}

type ErrorInterface interface {
	GetCode() ErrorCode
	Error() string
}

type Error struct {
	Code      ErrorCode `json:"code"`
	Message   string    `json:"message"`
	Operation string    `json:"operation"`
	Err       error     `json:"err"`
}

// New creates and returns a new error
func New(code ErrorCode, message, operation string, err error) *Error {
	return &Error{Code: code, Message: message, Operation: operation, Err: err}
}

// Wrap returns a new error that contains the passed error but with a different operation, useful for creating stacktraces
func Wrap(operation string, err error) *Error {
	return &Error{Code: ErrorCodeFromError(err), Message: ErrorMessageFromError(err), Operation: operation, Err: err}
}

func (ptr *Error) GetCode() ErrorCode { return ptr.Code }

// Error returns the string representation of the error message.
func (e *Error) Error() string {
	var buf bytes.Buffer

	// Print the current operation in our stack, if any.
	if e.Operation != "" {
		fmt.Fprintf(&buf, "%s: ", e.Operation)
	}

	// If wrapping an error, print its Error() message.
	// Otherwise print the error code & message.
	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	} else {
		if e.Code != ErrorCodeOk {
			fmt.Fprintf(&buf, "<%s> ", e.Code)
		}
		buf.WriteString(e.Message)
	}
	return buf.String()
}

// String returns a simplified string representation of the error message
func (e *Error) String() string {
	return fmt.Sprintf(`%s <%s> "%s"`, e.Operation, e.Code, e.Message)
}

// ErrorCode returns the code of the root error, if available.
// Otherwise returns ErrorCodeInternal.
func ErrorCodeFromError(err error) ErrorCode {
	if err == nil {
		return ErrorCodeOk
	} else if e, ok := err.(*Error); ok && e.Code != ErrorCodeOk {
		return e.Code
	} else if ok && e.Err != nil {
		return ErrorCodeFromError(e.Err)
	}
	return ErrorCodeInternal
}

// ErrorMessage returns the human-readable message of the error, if available.
// Otherwise returns a generic error message.
func ErrorMessageFromError(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.Message != "" {
		return e.Message
	} else if ok && e.Err != nil {
		return ErrorMessageFromError(e.Err)
	}
	return "An internal error has occurred. Please contact technical support."
}
