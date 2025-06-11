package utils

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrorHandler(err error, defaultCode codes.Code) error {
	if err == nil {
		return nil
	}
	return status.Errorf(defaultCode, "%v", err)
}
