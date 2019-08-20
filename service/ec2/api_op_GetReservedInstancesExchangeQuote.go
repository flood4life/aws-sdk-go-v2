// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for GetReservedInstanceExchangeQuote.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/GetReservedInstancesExchangeQuoteRequest
type GetReservedInstancesExchangeQuoteInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The IDs of the Convertible Reserved Instances to exchange.
	//
	// ReservedInstanceIds is a required field
	ReservedInstanceIds []string `locationName:"ReservedInstanceId" locationNameList:"ReservedInstanceId" type:"list" required:"true"`

	// The configuration of the target Convertible Reserved Instance to exchange
	// for your current Convertible Reserved Instances.
	TargetConfigurations []TargetConfigurationRequest `locationName:"TargetConfiguration" locationNameList:"TargetConfigurationRequest" type:"list"`
}

// String returns the string representation
func (s GetReservedInstancesExchangeQuoteInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetReservedInstancesExchangeQuoteInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetReservedInstancesExchangeQuoteInput"}

	if s.ReservedInstanceIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReservedInstanceIds"))
	}
	if s.TargetConfigurations != nil {
		for i, v := range s.TargetConfigurations {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "TargetConfigurations", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of GetReservedInstancesExchangeQuote.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/GetReservedInstancesExchangeQuoteResult
type GetReservedInstancesExchangeQuoteOutput struct {
	_ struct{} `type:"structure"`

	// The currency of the transaction.
	CurrencyCode *string `locationName:"currencyCode" type:"string"`

	// If true, the exchange is valid. If false, the exchange cannot be completed.
	IsValidExchange *bool `locationName:"isValidExchange" type:"boolean"`

	// The new end date of the reservation term.
	OutputReservedInstancesWillExpireAt *time.Time `locationName:"outputReservedInstancesWillExpireAt" type:"timestamp"`

	// The total true upfront charge for the exchange.
	PaymentDue *string `locationName:"paymentDue" type:"string"`

	// The cost associated with the Reserved Instance.
	ReservedInstanceValueRollup *ReservationValue `locationName:"reservedInstanceValueRollup" type:"structure"`

	// The configuration of your Convertible Reserved Instances.
	ReservedInstanceValueSet []ReservedInstanceReservationValue `locationName:"reservedInstanceValueSet" locationNameList:"item" type:"list"`

	// The cost associated with the Reserved Instance.
	TargetConfigurationValueRollup *ReservationValue `locationName:"targetConfigurationValueRollup" type:"structure"`

	// The values of the target Convertible Reserved Instances.
	TargetConfigurationValueSet []TargetReservationValue `locationName:"targetConfigurationValueSet" locationNameList:"item" type:"list"`

	// Describes the reason why the exchange cannot be completed.
	ValidationFailureReason *string `locationName:"validationFailureReason" type:"string"`
}

// String returns the string representation
func (s GetReservedInstancesExchangeQuoteOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetReservedInstancesExchangeQuote = "GetReservedInstancesExchangeQuote"

// GetReservedInstancesExchangeQuoteRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Returns a quote and exchange information for exchanging one or more specified
// Convertible Reserved Instances for a new Convertible Reserved Instance. If
// the exchange cannot be performed, the reason is returned in the response.
// Use AcceptReservedInstancesExchangeQuote to perform the exchange.
//
//    // Example sending a request using GetReservedInstancesExchangeQuoteRequest.
//    req := client.GetReservedInstancesExchangeQuoteRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/GetReservedInstancesExchangeQuote
func (c *Client) GetReservedInstancesExchangeQuoteRequest(input *GetReservedInstancesExchangeQuoteInput) GetReservedInstancesExchangeQuoteRequest {
	op := &aws.Operation{
		Name:       opGetReservedInstancesExchangeQuote,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetReservedInstancesExchangeQuoteInput{}
	}

	req := c.newRequest(op, input, &GetReservedInstancesExchangeQuoteOutput{})
	return GetReservedInstancesExchangeQuoteRequest{Request: req, Input: input, Copy: c.GetReservedInstancesExchangeQuoteRequest}
}

// GetReservedInstancesExchangeQuoteRequest is the request type for the
// GetReservedInstancesExchangeQuote API operation.
type GetReservedInstancesExchangeQuoteRequest struct {
	*aws.Request
	Input *GetReservedInstancesExchangeQuoteInput
	Copy  func(*GetReservedInstancesExchangeQuoteInput) GetReservedInstancesExchangeQuoteRequest
}

// Send marshals and sends the GetReservedInstancesExchangeQuote API request.
func (r GetReservedInstancesExchangeQuoteRequest) Send(ctx context.Context) (*GetReservedInstancesExchangeQuoteResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetReservedInstancesExchangeQuoteResponse{
		GetReservedInstancesExchangeQuoteOutput: r.Request.Data.(*GetReservedInstancesExchangeQuoteOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetReservedInstancesExchangeQuoteResponse is the response type for the
// GetReservedInstancesExchangeQuote API operation.
type GetReservedInstancesExchangeQuoteResponse struct {
	*GetReservedInstancesExchangeQuoteOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetReservedInstancesExchangeQuote request.
func (r *GetReservedInstancesExchangeQuoteResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}