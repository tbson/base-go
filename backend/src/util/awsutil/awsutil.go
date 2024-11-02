package awsutil

import (
	"context"
	"src/common/setting"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func S3Client() *s3.Client {
	region := setting.S3_REGION
	cfg, _ := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	return s3.NewFromConfig(cfg)
}
