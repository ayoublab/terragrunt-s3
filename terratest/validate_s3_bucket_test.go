package test

import (
    "context"
    "testing"

    
    "github.com/gruntwork-io/terratest/modules/aws"
    "github.com/gruntwork-io/terratest/modules/terraform"

    
    awsv2 "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/s3"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestValidateS3Bucket(t *testing.T) {
	
	t.Parallel()

	
	tgOptions := &terraform.Options{
		TerraformDir: "../envs/dev",
		
		TerraformBinary: "terragrunt",
	}

	
	bucketName := terraform.Output(t, tgOptions, "bucket_name")
	region := terraform.Output(t, tgOptions, "region")

	
	aws.AssertS3BucketExists(t, region, bucketName)
	
	
	aws.AssertS3BucketVersioningExists(t, region, bucketName)
	
	
	actualTags := aws.GetS3BucketTags(t, region, bucketName)
	assert.True(t, actualTags["environment"] == "dev")
	assert.True(t, actualTags["owner"] == "myself")
	
	assertBucketPublicAccessBlock(t, region, bucketName)
	
}


func assertBucketPublicAccessBlock(t *testing.T, region, bucket string) {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(region),
	)
	require.NoError(t, err)

	s3c := s3.NewFromConfig(cfg)

	out, err := s3c.GetPublicAccessBlock(
		context.TODO(),
		&s3.GetPublicAccessBlockInput{
			Bucket: awsv2.String(bucket),
		},
	)
	require.NoError(t, err, "Public-access-block config is missing on bucket %s", bucket)

	pab := out.PublicAccessBlockConfiguration
	assert.True(t, awsv2.ToBool(pab.BlockPublicAcls),       "BlockPublicAcls should be true")
	assert.True(t, awsv2.ToBool(pab.IgnorePublicAcls),      "IgnorePublicAcls should be true")
	assert.True(t, awsv2.ToBool(pab.BlockPublicPolicy),     "BlockPublicPolicy should be true")
	assert.True(t, awsv2.ToBool(pab.RestrictPublicBuckets), "RestrictPublicBuckets should be true")
}
