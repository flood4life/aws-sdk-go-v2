// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
)

// AWS doesn't currently have enough available capacity to fulfill your request.
// Try your request later.
type InsufficientCapacityException struct {
	Message *string
}

func (e *InsufficientCapacityException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InsufficientCapacityException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InsufficientCapacityException) ErrorCode() string             { return "InsufficientCapacityException" }
func (e *InsufficientCapacityException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Your request is valid, but Network Firewall couldn’t perform the operation
// because of a system problem. Retry your request.
type InternalServerError struct {
	Message *string
}

func (e *InternalServerError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServerError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServerError) ErrorCode() string             { return "InternalServerError" }
func (e *InternalServerError) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The operation failed because it's not valid. For example, you might have tried
// to delete a rule group or firewall policy that's in use.
type InvalidOperationException struct {
	Message *string
}

func (e *InvalidOperationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidOperationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidOperationException) ErrorCode() string             { return "InvalidOperationException" }
func (e *InvalidOperationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The operation failed because of a problem with your request. Examples
// include:
//
// * You specified an unsupported parameter name or value.
//
// * You tried
// to update a property with a value that isn't among the available types.
//
// * Your
// request references an ARN that is malformed, or corresponds to a resource that
// isn't valid in the context of the request.
type InvalidRequestException struct {
	Message *string
}

func (e *InvalidRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRequestException) ErrorCode() string             { return "InvalidRequestException" }
func (e *InvalidRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

//
type InvalidResourcePolicyException struct {
	Message *string
}

func (e *InvalidResourcePolicyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidResourcePolicyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidResourcePolicyException) ErrorCode() string             { return "InvalidResourcePolicyException" }
func (e *InvalidResourcePolicyException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The token you provided is stale or isn't valid for the operation.
type InvalidTokenException struct {
	Message *string
}

func (e *InvalidTokenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidTokenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidTokenException) ErrorCode() string             { return "InvalidTokenException" }
func (e *InvalidTokenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Unable to perform the operation because doing so would violate a limit setting.
type LimitExceededException struct {
	Message *string
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string             { return "LimitExceededException" }
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Unable to send logs to a configured logging destination.
type LogDestinationPermissionException struct {
	Message *string
}

func (e *LogDestinationPermissionException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LogDestinationPermissionException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LogDestinationPermissionException) ErrorCode() string {
	return "LogDestinationPermissionException"
}
func (e *LogDestinationPermissionException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Unable to locate a resource using the parameters that you provided.
type ResourceNotFoundException struct {
	Message *string
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string             { return "ResourceNotFoundException" }
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

//
type ResourceOwnerCheckException struct {
	Message *string
}

func (e *ResourceOwnerCheckException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceOwnerCheckException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceOwnerCheckException) ErrorCode() string             { return "ResourceOwnerCheckException" }
func (e *ResourceOwnerCheckException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Unable to process the request due to throttling limitations.
type ThrottlingException struct {
	Message *string
}

func (e *ThrottlingException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ThrottlingException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ThrottlingException) ErrorCode() string             { return "ThrottlingException" }
func (e *ThrottlingException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The operation you requested isn't supported by Network Firewall.
type UnsupportedOperationException struct {
	Message *string
}

func (e *UnsupportedOperationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedOperationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedOperationException) ErrorCode() string             { return "UnsupportedOperationException" }
func (e *UnsupportedOperationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }