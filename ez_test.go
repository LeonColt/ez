package ez_test

import (
	"reflect"
	"testing"

	"github.com/LeonColt/ez"
)

func getBuilder() interface{} {
	return &ez.Error{
		Code:    ez.ErrorCodeOk,
		Message: "OK",
	}
}

func TestCheckType(t *testing.T) {
	err := getBuilder()

	res, ok := err.(ez.ErrorInterface)
	if !ok {
		t.Fatalf("Expected type ez.ErrorInterface, got %T", err)
	}
	if reflect.TypeOf(res) != reflect.TypeOf(&ez.Error{}) {
		t.Fatalf("Expected type ez.Error, got %T", res)
	}
}

func TestGetCode(t *testing.T) {
	err := getBuilder()

	res, ok := err.(ez.ErrorInterface)
	if !ok {
		t.Fatalf("Expected type ez.ErrorInterface, got %T", err)
	}
	if reflect.TypeOf(res) != reflect.TypeOf(&ez.Error{}) {
		t.Errorf("Expected type ez.Error, got %T", res)
	}
	if ez.ErrorCode(ez.ErrorCodeOk) != res.GetCode() {
		t.Errorf("Expected error code %d, got %d", ez.ErrorCodeOk, res.GetCode())
	}
}

func TestGetError(t *testing.T) {
	err := getBuilder()

	res, ok := err.(ez.ErrorInterface)
	if !ok {
		t.Fatalf("Expected type ez.ErrorInterface, got %T", err)
	}
	if reflect.TypeOf(res) != reflect.TypeOf(&ez.Error{}) {
		t.Errorf("Expected type ez.Error, got %T", res)
	}

	if res.Error() != "OK" {
		t.Errorf("Expected error message 'OK', got '%s'", res.Error())
	}
}

func TestNew(t *testing.T) {
	const operation = "TestNew"
	err := ez.New(ez.ErrorCodeConflict, "An error message", ez.WithOperation(operation))

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}
	if ez.ErrorCode(ez.ErrorCodeConflict) != err.Code {
		t.Errorf("Expected error code %d, got %d", ez.ErrorCodeConflict, err.Code)
	}
	if err.Message != "An error message" {
		t.Errorf("Expected error message 'An error message', got '%s'", err.Message)
	}
	if err.Operation != operation {
		t.Errorf("Expected operation '%s', got '%s'", operation, err.Operation)
	}
	if err.Err != nil {
		t.Errorf("Expected no wrapped error, got '%v'", err.Err)
	}
}

func TestWrap(t *testing.T) {
	const operation = "TestNew"
	wrappedErr := ez.New(ez.ErrorCodeConflict, "An error message", ez.WithOperation(operation))

	err := ez.Wrap(wrappedErr)

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}
	if ez.ErrorCode(ez.ErrorCodeConflict) != err.Code {
		t.Errorf("Expected error code %d, got %d", ez.ErrorCodeConflict, err.Code)
	}
	if err.Message != "An error message" {
		t.Errorf("Expected error message 'An error message', got '%s'", err.Message)
	}
	if err.Operation != operation {
		t.Errorf("Expected operation '%s', got '%s'", operation, err.Operation)
	}
	if err.Err != wrappedErr {
		t.Errorf("Expected wrapped error '%v', got '%v'", wrappedErr, err.Err)
	}
}

func TestWrapWithOperation(t *testing.T) {
	const operation = "TestNew"
	wrappedErr := ez.New(ez.ErrorCodeConflict, "An error message", ez.WithOperation(operation))

	const newOperation = "TestWrap"
	err := ez.WrapWithOperation(newOperation, wrappedErr)

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}
	if ez.ErrorCode(ez.ErrorCodeConflict) != err.Code {
		t.Errorf("Expected error code %d, got %d", ez.ErrorCodeConflict, err.Code)
	}
	if err.Message != "An error message" {
		t.Errorf("Expected error message 'An error message', got '%s'", err.Message)
	}
	if err.Operation != newOperation {
		t.Errorf("Expected operation '%s', got '%s'", newOperation, err.Operation)
	}
	if err.Err != wrappedErr {
		t.Errorf("Expected wrapped error '%v', got '%v'", wrappedErr, err.Err)
	}
}

func TestError(t *testing.T) {
	const operation = "TestError"
	err := ez.New(ez.ErrorCodeConflict, "An internal error", ez.WithOperation(operation))

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}
	if err.Error() != "6: TestError: An internal error" {
		t.Errorf("Expected error message '6: TestError: An internal error', got '%s'", err.Error())
	}
}

func TestErrorCode(t *testing.T) {
	const operation = "TestErrorCode"
	err := ez.New(ez.ErrorCodeInvalidArgument, "An invalid error", ez.WithOperation(operation))

	code := ez.ErrorCodeFromError(err)

	if ez.ErrorCode(ez.ErrorCodeInvalidArgument) != code {
		t.Errorf("Expected error code %d, got %d", ez.ErrorCodeInvalidArgument, code)
	}
}

func TestErrorMessage(t *testing.T) {
	const op = "TestErrorMessage"
	err := ez.New(ez.ErrorCodeNotFound, "A not found error", ez.WithOperation(op))

	msg := ez.ErrorMessageFromError(err)

	if msg != "A not found error" {
		t.Errorf("Expected error message 'A not found error', got '%s'", msg)
	}
}
