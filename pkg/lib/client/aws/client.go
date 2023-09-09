package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

type client struct {
	ec2Client *ec2.Client
	stsClient *sts.Client
	s3Client  *s3.Client
	region    string
}

const (
	ExportImageFormatTypeKey = "image-format"
	OrigAmiTagKey            = "original-ami"
)

// New AWS client with the specified region.

func NewClient(region string) (*client, error) {

	// Load the SDK's configuration from environment and shared config, and
	// create the ec2Client with this.
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	if region == "" {
		region = cfg.Region
	}

	// Create EC2, STS, and S3 clients using the AWS SDK configuration
	ec2Client := ec2.NewFromConfig(cfg)
	stsClient := sts.NewFromConfig(cfg)
	s3Client := s3.NewFromConfig(cfg)

	return &client{
		ec2Client: ec2Client,
		stsClient: stsClient,
		s3Client:  s3Client,
		region:    region,
	}, nil
}
