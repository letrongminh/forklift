package suit

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"os"
	"testing"
)

func TestListS3Bucket(t *testing.T) {
	credsFile := "tests/suit/aws-cred.json"

	buckets, err := listS3Bucket(credsFile)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	fmt.Println("S3 Buckets:")
	for _, bucket := range buckets {
		fmt.Println(bucket)
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
