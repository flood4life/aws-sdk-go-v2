// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mobileanalytics

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

const opPutEvents = "PutEvents"

// PutEventsRequest generates a "aws.Request" representing the
// client's request for the PutEvents operation. The "output" return
// value will be populated with the request's response once the request complets
// successfuly.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See PutEvents for more information on using the PutEvents
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the PutEventsRequest method.
//    req, resp := client.PutEventsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MobileAnalytics) PutEventsRequest(input *PutEventsInput) (req *aws.Request, output *PutEventsOutput) {
	op := &aws.Operation{
		Name:       opPutEvents,
		HTTPMethod: "POST",
		HTTPPath:   "/2014-06-05/events",
	}

	if input == nil {
		input = &PutEventsInput{}
	}

	output = &PutEventsOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return
}

// PutEvents API operation for Amazon Mobile Analytics.
//
// The PutEvents operation records one or more events. You can have up to 1,500
// unique custom events per app, any combination of up to 40 attributes and
// metrics per custom event, and any number of attribute or metric values.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Mobile Analytics's
// API operation PutEvents for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeBadRequestException "BadRequestException"
//   An exception object returned when a request fails.
//
func (c *MobileAnalytics) PutEvents(input *PutEventsInput) (*PutEventsOutput, error) {
	req, out := c.PutEventsRequest(input)
	return out, req.Send()
}

// PutEventsWithContext is the same as PutEvents with the addition of
// the ability to pass a context and additional request options.
//
// See PutEvents for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MobileAnalytics) PutEventsWithContext(ctx aws.Context, input *PutEventsInput, opts ...aws.Option) (*PutEventsOutput, error) {
	req, out := c.PutEventsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// A JSON object representing a batch of unique event occurrences in your app.
type Event struct {
	_ struct{} `type:"structure"`

	// A collection of key-value pairs that give additional context to the event.
	// The key-value pairs are specified by the developer.
	//
	// This collection can be empty or the attribute object can be omitted.
	Attributes map[string]*string `locationName:"attributes" type:"map"`

	// A name signifying an event that occurred in your app. This is used for grouping
	// and aggregating like events together for reporting purposes.
	//
	// EventType is a required field
	EventType *string `locationName:"eventType" min:"1" type:"string" required:"true"`

	// A collection of key-value pairs that gives additional, measurable context
	// to the event. The key-value pairs are specified by the developer.
	//
	// This collection can be empty or the attribute object can be omitted.
	Metrics map[string]*float64 `locationName:"metrics" type:"map"`

	// The session the event occured within.
	Session *Session `locationName:"session" type:"structure"`

	// The time the event occurred in ISO 8601 standard date time format. For example,
	// 2014-06-30T19:07:47.885Z
	//
	// Timestamp is a required field
	Timestamp *string `locationName:"timestamp" type:"string" required:"true"`

	// The version of the event.
	Version *string `locationName:"version" min:"1" type:"string"`
}

// String returns the string representation
func (s Event) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Event) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Event) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Event"}

	if s.EventType == nil {
		invalidParams.Add(aws.NewErrParamRequired("EventType"))
	}
	if s.EventType != nil && len(*s.EventType) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EventType", 1))
	}

	if s.Timestamp == nil {
		invalidParams.Add(aws.NewErrParamRequired("Timestamp"))
	}
	if s.Version != nil && len(*s.Version) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Version", 1))
	}
	if s.Session != nil {
		if err := s.Session.Validate(); err != nil {
			invalidParams.AddNested("Session", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAttributes sets the Attributes field's value.
func (s *Event) SetAttributes(v map[string]*string) *Event {
	s.Attributes = v
	return s
}

// SetEventType sets the EventType field's value.
func (s *Event) SetEventType(v string) *Event {
	s.EventType = &v
	return s
}

// SetMetrics sets the Metrics field's value.
func (s *Event) SetMetrics(v map[string]*float64) *Event {
	s.Metrics = v
	return s
}

// SetSession sets the Session field's value.
func (s *Event) SetSession(v *Session) *Event {
	s.Session = v
	return s
}

// SetTimestamp sets the Timestamp field's value.
func (s *Event) SetTimestamp(v string) *Event {
	s.Timestamp = &v
	return s
}

// SetVersion sets the Version field's value.
func (s *Event) SetVersion(v string) *Event {
	s.Version = &v
	return s
}

// A container for the data needed for a PutEvent operation
type PutEventsInput struct {
	_ struct{} `type:"structure"`

	// The client context including the client ID, app title, app version and package
	// name.
	//
	// ClientContext is a required field
	ClientContext *string `location:"header" locationName:"x-amz-Client-Context" type:"string" required:"true"`

	// The encoding used for the client context.
	ClientContextEncoding *string `location:"header" locationName:"x-amz-Client-Context-Encoding" type:"string"`

	// An array of Event JSON objects
	//
	// Events is a required field
	Events []*Event `locationName:"events" type:"list" required:"true"`
}

// String returns the string representation
func (s PutEventsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PutEventsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutEventsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutEventsInput"}

	if s.ClientContext == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientContext"))
	}

	if s.Events == nil {
		invalidParams.Add(aws.NewErrParamRequired("Events"))
	}
	if s.Events != nil {
		for i, v := range s.Events {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Events", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientContext sets the ClientContext field's value.
func (s *PutEventsInput) SetClientContext(v string) *PutEventsInput {
	s.ClientContext = &v
	return s
}

// SetClientContextEncoding sets the ClientContextEncoding field's value.
func (s *PutEventsInput) SetClientContextEncoding(v string) *PutEventsInput {
	s.ClientContextEncoding = &v
	return s
}

// SetEvents sets the Events field's value.
func (s *PutEventsInput) SetEvents(v []*Event) *PutEventsInput {
	s.Events = v
	return s
}

type PutEventsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutEventsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PutEventsOutput) GoString() string {
	return s.String()
}

// Describes the session. Session information is required on ALL events.
type Session struct {
	_ struct{} `type:"structure"`

	// The duration of the session.
	Duration *int64 `locationName:"duration" type:"long"`

	// A unique identifier for the session
	Id *string `locationName:"id" min:"1" type:"string"`

	// The time the event started in ISO 8601 standard date time format. For example,
	// 2014-06-30T19:07:47.885Z
	StartTimestamp *string `locationName:"startTimestamp" type:"string"`

	// The time the event terminated in ISO 8601 standard date time format. For
	// example, 2014-06-30T19:07:47.885Z
	StopTimestamp *string `locationName:"stopTimestamp" type:"string"`
}

// String returns the string representation
func (s Session) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Session) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Session) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Session"}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDuration sets the Duration field's value.
func (s *Session) SetDuration(v int64) *Session {
	s.Duration = &v
	return s
}

// SetId sets the Id field's value.
func (s *Session) SetId(v string) *Session {
	s.Id = &v
	return s
}

// SetStartTimestamp sets the StartTimestamp field's value.
func (s *Session) SetStartTimestamp(v string) *Session {
	s.StartTimestamp = &v
	return s
}

// SetStopTimestamp sets the StopTimestamp field's value.
func (s *Session) SetStopTimestamp(v string) *Session {
	s.StopTimestamp = &v
	return s
}
