package suit

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"os"
	"testing"
)

func TestListS3Bucket(t *testing.T) {
	credsFile := "aws-cred.json"

	buckets, err := listS3Bucket(credsFile)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	fmt.Println("S3 Buckets:")
	for _, bucket := range buckets {
		fmt.Println(bucket)
	}
}

func TestListEC2Instances(t *testing.T) {
	credsFile := "aws-cred.json"

	instances, err := listEC2Instances(credsFile)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	fmt.Println("EC2 Instances and Their Status:")
	for _, instance := range instances {
		fmt.Printf("Instance ID: %s, Status: %s\n", *instance.InstanceId, instance.State.Name)
	}
}

func TestListEC2AMIs(t *testing.T) {
	credsFile := "aws-cred.json"

	amis, err := listEC2AMIs(credsFile)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	fmt.Println("EC2 AMIs:")
	for _, ami := range amis {
		fmt.Printf("AMI ID: %s, Name: %s\n", *ami.ImageId, *ami.Name)
	}
}

func listS3Bucket(credsFile string) ([]string, error) {
	// Read AWS credentials from JSON file
	creds, err := loadAWSCredentials(credsFile)
	if err != nil {
		return nil, err
	}

	// Configure AWS SDK with loaded credentials
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			creds["AccessKeyId"],
			creds["SecretAccessKey"],
			"",
		)),
		config.WithRegion(creds["Region"]),
	)
	if err != nil {
		return nil, err
	}

	// Create an S3 client
	client := s3.NewFromConfig(cfg)

	// List all S3 buckets
	listBucketsInput := &s3.ListBucketsInput{}
	listBucketsOutput, err := client.ListBuckets(context.TODO(), listBucketsInput)
	if err != nil {
		return nil, err
	}

	// Extract bucket names
	var bucketNames []string
	for _, bucket := range listBucketsOutput.Buckets {
		bucketNames = append(bucketNames, *bucket.Name)
	}

	return bucketNames, nil
}

func listEC2Instances(credsFile string) ([]types.Instance, error) {
	// Read AWS credentials from JSON file
	creds, err := loadAWSCredentials(credsFile)
	if err != nil {
		return nil, err
	}

	// Configure AWS SDK with loaded credentials
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			creds["AccessKeyId"],
			creds["SecretAccessKey"],
			"",
		)),
		config.WithRegion(creds["Region"]),
	)
	if err != nil {
		return nil, err
	}

	// Create an EC2 client
	client := ec2.NewFromConfig(cfg)

	// Describe all EC2 instances
	describeInstancesInput := &ec2.DescribeInstancesInput{}
	describeInstancesOutput, err := client.DescribeInstances(context.TODO(), describeInstancesInput)
	if err != nil {
		return nil, err
	}

	// Extract instance information
	var instances []types.Instance
	for _, reservation := range describeInstancesOutput.Reservations {
		instances = append(instances, reservation.Instances...)
	}

	return instances, nil
}

func listEC2AMIs(credsFile string) ([]types.Image, error) {
	// Read AWS credentials from JSON file
	creds, err := loadAWSCredentials(credsFile)
	if err != nil {
		return nil, err
	}

	// Configure AWS SDK with loaded credentials
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			creds["AccessKeyId"],
			creds["SecretAccessKey"],
			"",
		)),
		config.WithRegion(creds["Region"]),
	)
	if err != nil {
		return nil, err
	}

	// Create an EC2 client
	client := ec2.NewFromConfig(cfg)

	// Describe all EC2 AMIs
	describeImagesInput := &ec2.DescribeImagesInput{
		Owners: []string{"self"}, // List images owned by your account
	}
	describeImagesOutput, err := client.DescribeImages(context.TODO(), describeImagesInput)
	if err != nil {
		return nil, err
	}

	return describeImagesOutput.Images, nil
}

func loadAWSCredentials(filename string) (map[string]string, error) {
	var creds map[string]string

	data, err := os.ReadFile(filename)
	if err != nil {
		return creds, err
	}

	err = json.Unmarshal(data, &creds)
	if err != nil {
		return creds, err
	}

	return creds, nil
}
