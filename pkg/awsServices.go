package service

import (
	"fmt"

	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// AWSService - all available AWS services
type AWSService struct {
	S3        *s3.S3
	S3Manager *s3manager.Uploader
	DynamoDb  *dynamodb.DynamoDB
}

// SessionConfig - config required for AWS Session
type SessionConfig struct {
	Region  string
	Profile string
}

// NewSession makes aws services available
func NewSession(config *SessionConfig) (*AWSService, error) {

	if len(config.Profile) != 0 && len(config.Region) != 0 {
		fmt.Println(config)
		var sess, _ = session.NewSessionWithOptions(session.Options{
			Config:  aws.Config{Region: aws.String(config.Region)},
			Profile: config.Profile,
		})

		return &AWSService{
			S3:        s3.New(sess),
			S3Manager: s3manager.NewUploader(sess),
			DynamoDb:  dynamodb.New(sess),
		}, nil
	}
	fmt.Println("Unable to create session as config is invalid")

	return nil, errors.New("unable to initiate AWS Services")
}
