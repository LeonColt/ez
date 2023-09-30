package ez

import (
	"fmt"
)

type HttpException interface {
	GetCode() int
}

type HttpExceptionBuilder struct {
	Code    int
	Message string
}

func (ptr *HttpExceptionBuilder) GetCode() int { return ptr.Code }

func (ptr *HttpExceptionBuilder) Error() string { return ptr.Message }

type HttpExceptionBuilderWithError struct {
	HttpExceptionBuilder
	Err error
}

func (ptr *HttpExceptionBuilderWithError) GetCode() int { return ptr.Code }

func (ptr *HttpExceptionBuilderWithError) Error() string {
	return fmt.Sprintf("%s: %#v", ptr.Message, ptr.Err)
}
