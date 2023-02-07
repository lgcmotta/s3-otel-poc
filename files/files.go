package files

import (
	"context"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-sdk-go-v2/otelaws"
)

var bucketName *string

func GetFileURLWithoutOTEL(ctx context.Context, filename string) string {
	cfg, err := config.LoadDefaultConfig(ctx)

	if err != nil {
		panic(err)
	}

	client := s3.NewFromConfig(cfg)

	presignClient := s3.NewPresignClient(client)

	presignParams := &s3.GetObjectInput{
		Bucket: bucketName,
		Key:    aws.String(filename),
	}

	presignDuration := func(presignOptions *s3.PresignOptions) {
		presignOptions.Expires = 5 * time.Minute
	}

	presignResult, err := presignClient.PresignGetObject(ctx, presignParams, presignDuration)

	if err != nil {
		panic(err)
	}

	return presignResult.URL
}

func GetFileURLWithOTEL(ctx context.Context, filename string) string {
	cfg, err := config.LoadDefaultConfig(ctx)

	otelaws.AppendMiddlewares(&cfg.APIOptions)

	if err != nil {
		panic(err)
	}

	client := s3.NewFromConfig(cfg)

	presignClient := s3.NewPresignClient(client)

	presignParams := &s3.GetObjectInput{
		Bucket: bucketName,
		Key:    aws.String(filename),
	}

	presignDuration := func(presignOptions *s3.PresignOptions) {
		presignOptions.Expires = 5 * time.Minute
	}

	presignResult, err := presignClient.PresignGetObject(ctx, presignParams, presignDuration)

	if err != nil {
		panic(err)
	}

	return presignResult.URL
}

func init() {
	bucketName = aws.String(os.Getenv("BUCKET_NAME"))
}
