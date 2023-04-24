// Code generated by smithy-go-codegen DO NOT EDIT.

package ecr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Applies a repository policy to the specified repository to control access
// permissions. For more information, see Amazon ECR Repository policies (https://docs.aws.amazon.com/AmazonECR/latest/userguide/repository-policies.html)
// in the Amazon Elastic Container Registry User Guide.
func (c *Client) SetRepositoryPolicy(ctx context.Context, params *SetRepositoryPolicyInput, optFns ...func(*Options)) (*SetRepositoryPolicyOutput, error) {
	if params == nil {
		params = &SetRepositoryPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetRepositoryPolicy", params, optFns, c.addOperationSetRepositoryPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetRepositoryPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetRepositoryPolicyInput struct {

	// The JSON repository policy text to apply to the repository. For more
	// information, see Amazon ECR repository policies (https://docs.aws.amazon.com/AmazonECR/latest/userguide/repository-policy-examples.html)
	// in the Amazon Elastic Container Registry User Guide.
	//
	// This member is required.
	PolicyText *string

	// The name of the repository to receive the policy.
	//
	// This member is required.
	RepositoryName *string

	// If the policy you are attempting to set on a repository policy would prevent
	// you from setting another policy in the future, you must force the
	// SetRepositoryPolicy operation. This is intended to prevent accidental repository
	// lock outs.
	Force bool

	// The Amazon Web Services account ID associated with the registry that contains
	// the repository. If you do not specify a registry, the default registry is
	// assumed.
	RegistryId *string

	noSmithyDocumentSerde
}

type SetRepositoryPolicyOutput struct {

	// The JSON repository policy text applied to the repository.
	PolicyText *string

	// The registry ID associated with the request.
	RegistryId *string

	// The repository name associated with the request.
	RepositoryName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetRepositoryPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSetRepositoryPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSetRepositoryPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpSetRepositoryPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetRepositoryPolicy(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opSetRepositoryPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecr",
		OperationName: "SetRepositoryPolicy",
	}
}
