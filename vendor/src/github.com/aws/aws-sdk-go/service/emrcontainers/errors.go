// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package emrcontainers

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeEKSRequestThrottledException for service response error code
	// "EKSRequestThrottledException".
	//
	// The request exceeded the Amazon EKS API operation limits.
	ErrCodeEKSRequestThrottledException = "EKSRequestThrottledException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// This is an internal server exception.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeRequestThrottledException for service response error code
	// "RequestThrottledException".
	//
	// The request throttled.
	ErrCodeRequestThrottledException = "RequestThrottledException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The specified resource was not found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// There are invalid parameters in the client request.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"EKSRequestThrottledException": newErrorEKSRequestThrottledException,
	"InternalServerException":      newErrorInternalServerException,
	"RequestThrottledException":    newErrorRequestThrottledException,
	"ResourceNotFoundException":    newErrorResourceNotFoundException,
	"ValidationException":          newErrorValidationException,
}
