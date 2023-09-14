package err

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"os"
	"testing"
)

type Client struct {
	AWSConfigPath string
	ctx           context.Context
	ec2Client     *ec2.Client
	s3Client      *s3.Client
}

func NewClient(ctx context.Context, awsConfigPath string) *Client {
	return &Client{
		ctx:           ctx,
		AWSConfigPath: awsConfigPath,
	}
}

func (c *Client) SetAWSConfigPath() error {
	err := os.Setenv("AWS_SHARED_CREDENTIALS_FILE", c.AWSConfigPath)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Authenticate() (err error) {
	err = c.SetAWSConfigPath()
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) ConnectS3API() (err error) {
	err = c.Authenticate()
	if err != nil {
		return err
	}

	if c.s3Client == nil {
		cfg, err := config.LoadDefaultConfig(c.ctx)
		if err != nil {
			return err
		}
		c.s3Client = s3.NewFromConfig(cfg)
	}
	return nil
}

func (c *Client) ListS3Buckets() error {
	err := c.ConnectS3API()
	if err != nil {
		return err
	}

	// List S3 buckets
	input := &s3.ListBucketsInput{}
	resp, err := c.s3Client.ListBuckets(c.ctx, input)
	if err != nil {
		return err
	}

	fmt.Println("List of S3 Buckets:")
	for _, bucket := range resp.Buckets {
		fmt.Println(*bucket.Name)
	}

	return nil
}

func TestListS3Buckets(t *testing.T) {
	ctx := context.TODO()
	awsConfigPath := "aws-cred"
	client := NewClient(ctx, awsConfigPath)

	err := client.ListS3Buckets()
	if err != nil {
		t.Errorf("Error listing S3 buckets: %v", err)
	}
}

func (c *Client) ConnectEC2API() (err error) {
	err = c.Authenticate()
	if err != nil {
		return err
	}

	if c.ec2Client == nil {
		cfg, err := config.LoadDefaultConfig(c.ctx)
		if err != nil {
			return err
		}
		c.ec2Client = ec2.NewFromConfig(cfg)
	}
	return nil
}

func TestConnectEC2Service(t *testing.T) {
	// Initialize the AWS client for testing

	ctx := context.TODO()
	awsConfigPath := "aws-cred.json"
	client := NewClient(ctx, awsConfigPath)

	err := client.ConnectEC2API()
	if err != nil {
		fmt.Println("Error connecting to EC2 service:", err)
	}

	fmt.Println("Connected to EC2 service successfully!")
}

func (c *Client) VMStart(instanceID string) error {
	err := c.Authenticate()
	if err != nil {
		return err
	}

	input := &ec2.StartInstancesInput{
		InstanceIds: []string{instanceID},
	}

	_, err = c.ec2Client.StartInstances(c.ctx, input)
	if err != nil {
		return err
	}

	return nil
}
