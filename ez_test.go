package ez_test

import (
	"testing"

	"github.com/LeonColt/ez"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	require.True(t, ok)
	require.IsType(t, &ez.Error{}, res)
}

func TestGetCode(t *testing.T) {
	err := getBuilder()

	res, ok := err.(ez.ErrorInterface)
	require.True(t, ok)
	require.IsType(t, &ez.Error{}, res)

	require.Equal(t, ez.ErrorCode(ez.ErrorCodeOk), ez.ErrorCode(ez.ErrorCodeOk))
}

func TestGetError(t *testing.T) {
	err := getBuilder()

	res, ok := err.(ez.ErrorInterface)
	require.True(t, ok)
	require.IsType(t, &ez.Error{}, res)

	require.Equal(t, "OK", res.Error())
}

func TestNew(t *testing.T) {
	const operation = "TestNew"
	err := ez.New(ez.ErrorCodeConflict, "An error message", operation, nil)

	assert.NotNil(t, err)
	assert.Equal(t, ez.ErrorCode(ez.ErrorCodeConflict), err.Code)
	assert.Equal(t, err.Message, "An error message")
	assert.Equal(t, err.Operation, operation)
	assert.Equal(t, err.Err, nil)
}

func TestWrap(t *testing.T) {
	const operation = "TestNew"
	wrappedErr := ez.New(ez.ErrorCodeConflict, "An error message", operation, nil)

	const newOperation = "TestWrap"
	err := ez.Wrap(newOperation, wrappedErr)

	assert.NotNil(t, err)
	assert.Equal(t, ez.ErrorCode(ez.ErrorCodeConflict), err.Code)
	assert.Equal(t, err.Message, "An error message")
	assert.Equal(t, err.Operation, newOperation)
	assert.Equal(t, err.Err, wrappedErr)
}

func TestError(t *testing.T) {
	const operation = "TestError"
	err := ez.New(ez.ErrorCodeConflict, "An internal error", operation, nil)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "6: TestError: An internal error")
}

func TestErrorCode(t *testing.T) {
	const operation = "TestErrorCode"
	err := ez.New(ez.ErrorCodeInvalidArgument, "An invalid error", operation, nil)

	code := ez.ErrorCodeFromError(err)

	assert.NotNil(t, err)
	assert.Equal(t, ez.ErrorCode(ez.ErrorCodeInvalidArgument), code)
}

func TestErrorMessage(t *testing.T) {
	const op = "TestErrorMessage"
	err := ez.New(ez.ErrorCodeNotFound, "A not found error", op, nil)

	msg := ez.ErrorMessageFromError(err)

	assert.NotNil(t, err)
	assert.Equal(t, msg, "A not found error")
}
