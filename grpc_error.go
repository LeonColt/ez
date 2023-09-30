package ez

import (
	"fmt"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type GrpcException interface {
	GetGrpcCode() codes.Code
}

type GrpcErrorBuilder struct {
	GrpcCode codes.Code
	Message  string
}

func (ptr *GrpcErrorBuilder) GetGrpcCode() codes.Code { return ptr.GrpcCode }

func (ptr *GrpcErrorBuilder) Error() string { return ptr.Message }

type GrpcErrorBuilderWithError struct {
	GrpcErrorBuilder
	Err error
}

func (ptr *GrpcErrorBuilderWithError) GetGrpcCode() codes.Code { return ptr.GrpcCode }

func (ptr *GrpcErrorBuilderWithError) Error() string {
	return fmt.Sprintf("%s: %#v", ptr.Message, ptr.Err)
}

func HandleGrpcError(err error) error {
	if err == nil {
		return status.Error(codes.Unknown, "unknown error")
	} else {
		if grpcErr, ok := err.(GrpcException); ok {
			return status.Error(grpcErr.GetGrpcCode(), err.Error())
		} else {
			return status.Errorf(codes.Internal, "unknown error: %#v", err)
		}
	}
}
