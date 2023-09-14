package suit

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"log"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func TestListS3Buckets(t *testing.T) {
	// Load AWS credentials from the JSON file
	cfg, err := loadAWSConfigFromFile("aws-cred.json")
	if err != nil {
		t.Errorf("Error loading AWS config: %v", err)
		return
	}

	// Create an S3 client
	client := s3.NewFromConfig(cfg)

	// List S3 buckets
	buckets, err := listS3Buckets(context.Background(), client)
	if err != nil {
		t.Errorf("Error listing S3 buckets: %v", err)
		return
	}

	// Check if the number of buckets is greater than zero
	if len(buckets) == 0 {
		t.Errorf("Expected at least one S3 bucket, got 0")
		return
	}

	// Print the list of S3 bucket names for testing purposes
	fmt.Println("S3 Buckets for Testing:")
	for _, bucket := range buckets {
		fmt.Println(bucket)
	}
}

func loadAWSConfigFromFile(configFilePath string) (aws.Config, error) {
	var creds map[string]string

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &creds)
	//if err != nil {
	//	return creds, err
	//}

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			creds["AccessKeyId"],
			creds["SecretAccessKey"],
			"",
		)),
		config.WithRegion(creds["Region"]),
	)

	//if err != nil {
	//	return nil, err
	//}

	return cfg, nil
}

func listS3Buckets(ctx context.Context, client *s3.Client) ([]string, error) {
	// List S3 buckets
	listBucketsOutput, err := client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		return nil, err
	}

	// Extract and return the list of S3 bucket names
	var bucketNames []string
	for _, bucket := range listBucketsOutput.Buckets {
		bucketNames = append(bucketNames, *bucket.Name)
	}

	return bucketNames, nil
}
