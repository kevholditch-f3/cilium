// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Detaches a network interface from an instance.
func (c *Client) DetachNetworkInterface(ctx context.Context, params *DetachNetworkInterfaceInput, optFns ...func(*Options)) (*DetachNetworkInterfaceOutput, error) {
	if params == nil {
		params = &DetachNetworkInterfaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetachNetworkInterface", params, optFns, addOperationDetachNetworkInterfaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetachNetworkInterfaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DetachNetworkInterface.
type DetachNetworkInterfaceInput struct {

	// The ID of the attachment.
	//
	// This member is required.
	AttachmentId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// Specifies whether to force a detachment.
	//
	// * Use the Force parameter only as a
	// last resort to detach a network interface from a failed instance.
	//
	// * If you use
	// the Force parameter to detach a network interface, you might not be able to
	// attach a different network interface to the same index on the instance without
	// first stopping and starting the instance.
	//
	// * If you force the detachment of a
	// network interface, the instance metadata
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html)
	// might not get updated. This means that the attributes associated with the
	// detached network interface might still be visible. The instance metadata will
	// get updated when you stop and start the instance.
	Force *bool
}

type DetachNetworkInterfaceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDetachNetworkInterfaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDetachNetworkInterface{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDetachNetworkInterface{}, middleware.After)
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
	if err = addOpDetachNetworkInterfaceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDetachNetworkInterface(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDetachNetworkInterface(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DetachNetworkInterface",
	}
}
