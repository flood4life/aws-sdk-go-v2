// Code generated by smithy-go-codegen DO NOT EDIT.

package types

// Information that shows whether a resource is compliant with the effective tag
// policy, including details on any noncompliant tag keys.
type ComplianceDetails struct {

	// Whether a resource is compliant with the effective tag policy.
	ComplianceStatus *bool

	// These are keys defined in the effective policy that are on the resource with
	// either incorrect case treatment or noncompliant values.
	KeysWithNoncompliantValues []*string

	// These tag keys on the resource are noncompliant with the effective tag policy.
	NoncompliantKeys []*string
}

// Information about the errors that are returned for each failed resource. This
// information can include InternalServiceException and InvalidParameterException
// errors. It can also include any valid error code returned by the AWS service
// that hosts the resource that the ARN key represents. The following are common
// error codes that you might receive from other AWS services:
//
// *
// InternalServiceException – This can mean that the Resource Groups Tagging API
// didn't receive a response from another AWS service. It can also mean the the
// resource type in the request is not supported by the Resource Groups Tagging
// API. In these cases, it's safe to retry the request and then call GetResources
// (http://docs.aws.amazon.com/resourcegroupstagging/latest/APIReference/API_GetResources.html)
// to verify the changes.
//
// * AccessDeniedException – This can mean that you need
// permission to calling tagging operations in the AWS service that contains the
// resource. For example, to use the Resource Groups Tagging API to tag a
// CloudWatch alarm resource, you need permission to call TagResources
// (http://docs.aws.amazon.com/resourcegroupstagging/latest/APIReference/API_TagResources.html)
// and TagResource
// (http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_TagResource.html)
// in the CloudWatch API.
//
// For more information on errors that are generated from
// other AWS services, see the documentation for that service.
type FailureInfo struct {

	// The code of the common error. Valid values include InternalServiceException,
	// InvalidParameterException, and any valid error code returned by the AWS service
	// that hosts the resource that you want to tag.
	ErrorCode ErrorCode

	// The message of the common error.
	ErrorMessage *string

	// The HTTP status code of the common error.
	StatusCode *int32
}

// A list of resource ARNs and the tags (keys and values) that are associated with
// each.
type ResourceTagMapping struct {

	// Information that shows whether a resource is compliant with the effective tag
	// policy, including details on any noncompliant tag keys.
	ComplianceDetails *ComplianceDetails

	// The ARN of the resource.
	ResourceARN *string

	// The tags that have been applied to one or more AWS resources.
	Tags []*Tag
}

// A count of noncompliant resources.
type Summary struct {

	// The timestamp that shows when this summary was generated in this Region.
	LastUpdated *string

	// The count of noncompliant resources.
	NonCompliantResources *int64

	// The AWS Region that the summary applies to.
	Region *string

	// The AWS resource type.
	ResourceType *string

	// The account identifier or the root identifier of the organization. If you don't
	// know the root ID, you can call the AWS Organizations ListRoots
	// (http://docs.aws.amazon.com/organizations/latest/APIReference/API_ListRoots.html)
	// API.
	TargetId *string

	// Whether the target is an account, an OU, or the organization root.
	TargetIdType TargetIdType
}

// The metadata that you apply to AWS resources to help you categorize and organize
// them. Each tag consists of a key and a value, both of which you define. For more
// information, see Tagging AWS Resources
// (http://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the AWS
// General Reference.
type Tag struct {

	// One part of a key-value pair that makes up a tag. A key is a general label that
	// acts like a category for more specific tag values.
	//
	// This member is required.
	Key *string

	// One part of a key-value pair that make up a tag. A value acts as a descriptor
	// within a tag category (key). The value can be empty or null.
	//
	// This member is required.
	Value *string
}

// A list of tags (keys and values) that are used to specify the associated
// resources.
type TagFilter struct {

	// One part of a key-value pair that makes up a tag. A key is a general label that
	// acts like a category for more specific tag values.
	Key *string

	// One part of a key-value pair that make up a tag. A value acts as a descriptor
	// within a tag category (key). The value can be empty or null.
	Values []*string
}