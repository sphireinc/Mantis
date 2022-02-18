package aws

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

// NewDynamo creates and returns a new instance of dynamo db
func (a *AWS) NewDynamo() *dynamodb.Client {
	return dynamodb.NewFromConfig(a.config)
}

// NewSecretsManager creates and returns a new instance of dynamo db
func (a *AWS) NewSecretsManager(optFns ...func(*secretsmanager.Options)) *secretsmanager.Client {
	return secretsmanager.NewFromConfig(a.config, optFns...)
}
