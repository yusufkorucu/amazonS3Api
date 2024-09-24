package amazonS3Api

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"net/url"
)

type AmazonS3ApiClient struct {
	client *s3.S3
	config *S3ApiConfig
}

func NewAmazonS3ApiClient() *AmazonS3ApiClient {

	amazonS3ApiConfigs := GetS3Configs()
	serviceURL, err := url.Parse(fmt.Sprintf("%s/%s", amazonS3ApiConfigs.CdnUrl, amazonS3ApiConfigs.BucketName))
	if err != nil {
		log.Fatal("Failed to parse service URL: %v", err)
	}

	config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(amazonS3ApiConfigs.AccessKeyId, amazonS3ApiConfigs.SecretAccessKey, ""),
		Endpoint:         aws.String(serviceURL.String()),
		Region:           aws.String("us-east-1"),
		S3ForcePathStyle: aws.Bool(true),
	}

	s3session, err := session.NewSession(config)
	if err != nil {
		log.Fatal("Failed to create session: %v", err)
	}

	client := s3.New(s3session)
	return &AmazonS3ApiClient{
		client: client,
		config: amazonS3ApiConfigs,
	}
}

func (s *AmazonS3ApiClient) UploadFile(uploadPath string, fileName string, file []byte) (error, string) {
	key := fmt.Sprintf("%s/%s", uploadPath, fileName)

	_, err := s.client.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(s.config.BucketName),
		Key:         aws.String(key),
		Body:        bytes.NewReader(file),
		ContentType: aws.String("application/octet-stream"),
	})
	if err != nil {
		return fmt.Errorf("failed to upload file to S3: %w", err), ""
	}

	downloadUrl := fmt.Sprintf("%s/%s/%s/%s", s.config.CdnUrl, s.config.BucketName, s.config.BucketName, key)
	return nil, downloadUrl
}
