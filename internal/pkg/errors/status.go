package errors

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

type StatusCode struct {
	http int
	grpc codes.Code
}

var (
	StatusInternalServer      = &StatusCode{http: http.StatusInternalServerError, grpc: codes.Internal}
	StatusBadRequest          = &StatusCode{http: http.StatusBadRequest, grpc: codes.InvalidArgument}
	StatusUnauthorized        = &StatusCode{http: http.StatusUnauthorized, grpc: codes.Unauthenticated}
	StatusForbidden           = &StatusCode{http: http.StatusForbidden, grpc: codes.PermissionDenied}
	StatusAccepted            = &StatusCode{http: http.StatusAccepted, grpc: codes.OK}
	StatusOK                  = &StatusCode{http: http.StatusOK, grpc: codes.OK}
	StatusCreated             = &StatusCode{http: http.StatusCreated, grpc: codes.OK}
	StatusNoContent           = &StatusCode{http: http.StatusNoContent, grpc: codes.OK}
	StatusAlreadyExists       = &StatusCode{http: http.StatusBadRequest, grpc: codes.AlreadyExists}
	StatusNotFound            = &StatusCode{http: http.StatusNotFound, grpc: codes.NotFound}
	StatusServiceUnavailable  = &StatusCode{http: http.StatusServiceUnavailable, grpc: codes.Unavailable}
	StatusInsufficientStorage = &StatusCode{http: http.StatusInsufficientStorage, grpc: codes.ResourceExhausted}
	StatusNotImplemented      = &StatusCode{http: http.StatusNotImplemented, grpc: codes.Unimplemented}
)
