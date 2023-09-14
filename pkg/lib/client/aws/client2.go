package aws

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// AWSConfig represents the JSON configuration structure.
type AWSConfig struct {
	AccessKeyId     string `json:"AccessKeyId"`
	SecretAccessKey string `json:"SecretAccessKey"`
	Region          string `json:"Region"`
}

// Client holds the AWS S3 client.
type Client struct {
	AWSAuthPath string
	ctx         context.Context
	S3          *s3.Client
	EC2         *ec2.Client
}

// LoadAWSConfigFromFile loads AWS configuration from a JSON file.
func LoadAWSConfigFromFile(configFilePath string) (*AWSConfig, error) {
	var cfg AWSConfig

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// Authenticate authenticates with AWS using the provided credentials and region.
func (c *Client) Authenticate(cfg *AWSConfig) (aws.Config, error) {
	cfgProvider := credentials.NewStaticCredentialsProvider(cfg.AccessKeyId, cfg.SecretAccessKey, "")
	return config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(cfg.Region),
		config.WithCredentialsProvider(cfgProvider),
	)
}

func (c *Client) connectComputeServiceAPI(cfg *AWSConfig) (*ec2.Client, error) {
	cfgAWS, err := c.Authenticate(cfg)
	if err != nil {
		return nil, err
	}

	return ec2.NewFromConfig(cfgAWS), nil
}

func (c *Client) connectImageServiceAPI(cfg *AWSConfig) (*ec2.Client, error) {
	cfgAWS, err := c.Authenticate(cfg)

	if err != nil {
		return nil, err
	}

	return ec2.NewFromConfig(cfgAWS), nil
}

// ConnectS3ServiceAPI connects to the AWS S3 service and returns an S3 client.
func (c *Client) connectStorageServiceAPI(cfg *AWSConfig) (*s3.Client, error) {
	cfgAWS, err := c.Authenticate(cfg)
	if err != nil {
		return nil, err
	}

	return s3.NewFromConfig(cfgAWS), nil
}

func listEC2Instance(ec2Client *ec2.Client) error {
	input := &ec2.DescribeInstancesInput{}
	result, err := ec2Client.DescribeInstances(context.TODO(), input)
	if err != nil {
		return err
	}
	log.Println("Instances:")
	var instances []types.Instance
	for _, reservation := range result.Reservations {
		instances = append(instances, reservation.Instances...)
		log.Println(instances)
	}
	return nil
}

func listEC2AMIs(ec2Client *ec2.Client) error {
	input := &ec2.DescribeImagesInput{
		Owners: []string{"self"}, // List images owned by your account
	}

	result, err := ec2Client.DescribeImages(context.TODO(), input)
	if err != nil {
		return err
	}
	log.Println("AMI:")
	for _, ami := range result.Images {
		log.Println(*ami.Name)
	}
	return nil
}

// ListS3Buckets lists all S3 buckets in the authenticated AWS account.
func ListS3Buckets(s3Client *s3.Client) error {
	input := &s3.ListBucketsInput{}
	result, err := s3Client.ListBuckets(context.TODO(), input)
	if err != nil {
		return err
	}

	log.Println("S3 Buckets:")
	for _, bucket := range result.Buckets {
		log.Println(*bucket.Name)
	}

	return nil
}

func main() {
	configFilePath := "/Users/trongminhle/Library/CloudStorage/OneDrive-ITMOUNIVERSITY/Dev/gitlab-VT/forklift/pkg/lib/client/aws/aws-cred.json" // Replace with your JSON file path
	cfg, err := LoadAWSConfigFromFile(configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	client := &Client{}

	s3Client, err := client.connectStorageServiceAPI(cfg)
	if err != nil {
		log.Fatal(err)
	}

	if err := ListS3Buckets(s3Client); err != nil {
		log.Fatal(err)
	}

	ec2Client, err := client.connectImageServiceAPI(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Now you can use s3Client to perform operations with the AWS S3 service.
	if err := listEC2AMIs(ec2Client); err != nil {
		log.Fatal(err)
	}

	ec2Clients, err := client.connectImageServiceAPI(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Now you can use s3Client to perform operations with the AWS S3 service.
	if err := listEC2Instance(ec2Clients); err != nil {
		log.Fatal(err)
	}
}
