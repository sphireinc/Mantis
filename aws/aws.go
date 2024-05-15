package aws

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

// AWS primary struct with context and aws.config
type AWS struct {
	context context.Context
	config  aws.Config
}

// New creates a new AWS config using the context and region specified
//
//	If you want an empty context, set ctx to context.TODO()
//	region should be an AWS region, like "us-west-2"
func New(ctx context.Context, region string) *AWS {
	// Using the SDK's default configuration, loading additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	return &AWS{
		context: ctx,
		config:  cfg,
	}
}
