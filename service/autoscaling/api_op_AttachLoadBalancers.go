// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/AttachLoadBalancersType
type AttachLoadBalancersInput struct {
	_ struct{} `type:"structure"`

	// The name of the Auto Scaling group.
	//
	// AutoScalingGroupName is a required field
	AutoScalingGroupName *string `min:"1" type:"string" required:"true"`

	// The names of the load balancers. You can specify up to 10 load balancers.
	//
	// LoadBalancerNames is a required field
	LoadBalancerNames []string `type:"list" required:"true"`
}

// String returns the string representation
func (s AttachLoadBalancersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AttachLoadBalancersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AttachLoadBalancersInput"}

	if s.AutoScalingGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AutoScalingGroupName"))
	}
	if s.AutoScalingGroupName != nil && len(*s.AutoScalingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AutoScalingGroupName", 1))
	}

	if s.LoadBalancerNames == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerNames"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/AttachLoadBalancersResultType
type AttachLoadBalancersOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AttachLoadBalancersOutput) String() string {
	return awsutil.Prettify(s)
}

const opAttachLoadBalancers = "AttachLoadBalancers"

// AttachLoadBalancersRequest returns a request value for making API operation for
// Auto Scaling.
//
// Attaches one or more Classic Load Balancers to the specified Auto Scaling
// group.
//
// To attach an Application Load Balancer or a Network Load Balancer instead,
// see AttachLoadBalancerTargetGroups.
//
// To describe the load balancers for an Auto Scaling group, use DescribeLoadBalancers.
// To detach the load balancer from the Auto Scaling group, use DetachLoadBalancers.
//
// For more information, see Attaching a Load Balancer to Your Auto Scaling
// Group (https://docs.aws.amazon.com/autoscaling/ec2/userguide/attach-load-balancer-asg.html)
// in the Amazon EC2 Auto Scaling User Guide.
//
//    // Example sending a request using AttachLoadBalancersRequest.
//    req := client.AttachLoadBalancersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/AttachLoadBalancers
func (c *Client) AttachLoadBalancersRequest(input *AttachLoadBalancersInput) AttachLoadBalancersRequest {
	op := &aws.Operation{
		Name:       opAttachLoadBalancers,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AttachLoadBalancersInput{}
	}

	req := c.newRequest(op, input, &AttachLoadBalancersOutput{})
	return AttachLoadBalancersRequest{Request: req, Input: input, Copy: c.AttachLoadBalancersRequest}
}

// AttachLoadBalancersRequest is the request type for the
// AttachLoadBalancers API operation.
type AttachLoadBalancersRequest struct {
	*aws.Request
	Input *AttachLoadBalancersInput
	Copy  func(*AttachLoadBalancersInput) AttachLoadBalancersRequest
}

// Send marshals and sends the AttachLoadBalancers API request.
func (r AttachLoadBalancersRequest) Send(ctx context.Context) (*AttachLoadBalancersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AttachLoadBalancersResponse{
		AttachLoadBalancersOutput: r.Request.Data.(*AttachLoadBalancersOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AttachLoadBalancersResponse is the response type for the
// AttachLoadBalancers API operation.
type AttachLoadBalancersResponse struct {
	*AttachLoadBalancersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AttachLoadBalancers request.
func (r *AttachLoadBalancersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}