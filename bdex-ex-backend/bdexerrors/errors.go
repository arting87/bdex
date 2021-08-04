package bdexerrors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	_ = 10000 + iota // 10000 + 0
	CodeInvalidArg
	CodeBadRequest
	CodeInvalidName
	CodeUnauthorized
	CodeNotFound
	CodeInChain
	CodeNilErr
)

var (
	ErrDataLoss     = statusError(codes.DataLoss, "unrecoverable data loss or corruption")
	ErrInvalidArg   = statusError(CodeInvalidArg, "invalid argument")
	ErrBadRequest   = statusError(CodeBadRequest, "bad request")
	ErrInvalidName  = statusError(CodeInvalidName, "invalid account name")
	ErrUnauthorized = statusError(CodeUnauthorized, "Unauthorized account")
	ErrNotFound     = statusError(CodeNotFound, "resource not found")
	ErrInChain      = statusError(CodeInChain, "send to chain err")
	ErrNil          = statusError(CodeNilErr, "nil Error")
)

func statusError(code codes.Code, msg string) error {
	return status.New(code, msg).Err()
}
