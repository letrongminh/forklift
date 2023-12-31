// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes resource discovery association with an Amazon VPC IPAM. An associated
// resource discovery is a resource discovery that has been associated with an
// IPAM..
func (c *Client) DescribeIpamResourceDiscoveryAssociations(ctx context.Context, params *DescribeIpamResourceDiscoveryAssociationsInput, optFns ...func(*Options)) (*DescribeIpamResourceDiscoveryAssociationsOutput, error) {
	if params == nil {
		params = &DescribeIpamResourceDiscoveryAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeIpamResourceDiscoveryAssociations", params, optFns, c.addOperationDescribeIpamResourceDiscoveryAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeIpamResourceDiscoveryAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeIpamResourceDiscoveryAssociationsInput struct {

	// A check for whether you have the required permissions for the action without
	// actually making the request and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The resource discovery association filters.
	Filters []types.Filter

	// The resource discovery association IDs.
	IpamResourceDiscoveryAssociationIds []string

	// The maximum number of resource discovery associations to return in one page of
	// results.
	MaxResults *int32

	// Specify the pagination token from a previous request to retrieve the next page
	// of results.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeIpamResourceDiscoveryAssociationsOutput struct {

	// The resource discovery associations.
	IpamResourceDiscoveryAssociations []types.IpamResourceDiscoveryAssociation

	// Specify the pagination token from a previous request to retrieve the next page
	// of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeIpamResourceDiscoveryAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeIpamResourceDiscoveryAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeIpamResourceDiscoveryAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addDescribeIpamResourceDiscoveryAssociationsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeIpamResourceDiscoveryAssociations(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

// DescribeIpamResourceDiscoveryAssociationsAPIClient is a client that implements
// the DescribeIpamResourceDiscoveryAssociations operation.
type DescribeIpamResourceDiscoveryAssociationsAPIClient interface {
	DescribeIpamResourceDiscoveryAssociations(context.Context, *DescribeIpamResourceDiscoveryAssociationsInput, ...func(*Options)) (*DescribeIpamResourceDiscoveryAssociationsOutput, error)
}

var _ DescribeIpamResourceDiscoveryAssociationsAPIClient = (*Client)(nil)

// DescribeIpamResourceDiscoveryAssociationsPaginatorOptions is the paginator
// options for DescribeIpamResourceDiscoveryAssociations
type DescribeIpamResourceDiscoveryAssociationsPaginatorOptions struct {
	// The maximum number of resource discovery associations to return in one page of
	// results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeIpamResourceDiscoveryAssociationsPaginator is a paginator for
// DescribeIpamResourceDiscoveryAssociations
type DescribeIpamResourceDiscoveryAssociationsPaginator struct {
	options   DescribeIpamResourceDiscoveryAssociationsPaginatorOptions
	client    DescribeIpamResourceDiscoveryAssociationsAPIClient
	params    *DescribeIpamResourceDiscoveryAssociationsInput
	nextToken *string
	firstPage bool
}

// NewDescribeIpamResourceDiscoveryAssociationsPaginator returns a new
// DescribeIpamResourceDiscoveryAssociationsPaginator
func NewDescribeIpamResourceDiscoveryAssociationsPaginator(client DescribeIpamResourceDiscoveryAssociationsAPIClient, params *DescribeIpamResourceDiscoveryAssociationsInput, optFns ...func(*DescribeIpamResourceDiscoveryAssociationsPaginatorOptions)) *DescribeIpamResourceDiscoveryAssociationsPaginator {
	if params == nil {
		params = &DescribeIpamResourceDiscoveryAssociationsInput{}
	}

	options := DescribeIpamResourceDiscoveryAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeIpamResourceDiscoveryAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeIpamResourceDiscoveryAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeIpamResourceDiscoveryAssociations page.
func (p *DescribeIpamResourceDiscoveryAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeIpamResourceDiscoveryAssociationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.DescribeIpamResourceDiscoveryAssociations(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeIpamResourceDiscoveryAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeIpamResourceDiscoveryAssociations",
	}
}

type opDescribeIpamResourceDiscoveryAssociationsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDescribeIpamResourceDiscoveryAssociationsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeIpamResourceDiscoveryAssociationsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "ec2"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "ec2"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("ec2")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addDescribeIpamResourceDiscoveryAssociationsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDescribeIpamResourceDiscoveryAssociationsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
