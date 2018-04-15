package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var db *dynamodb.DynamoDB

// Init initializes DynamoDB
func Init(Profile string, Region string) {
	if len(Region) != 0 && len(Profile) != 0 {
		var sess, _ = session.NewSessionWithOptions(session.Options{
			Config:  aws.Config{Region: aws.String(Region)},
			Profile: Profile,
		})

		db = dynamodb.New(sess)
	} else {

	}
}

// GetDB returns instance of DynamoDB
func GetDB() *dynamodb.DynamoDB {
	return db
}
