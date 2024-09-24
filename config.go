package amazonS3Api

var s3ApiConfig S3ApiConfig

type S3ApiConfig struct {
	CdnUrl          string `json:"cdnUrl"`
	AccessKeyId     string `json:"accessKeyId"`
	SecretAccessKey string `json:"secretAccessKey"`
	BucketName      string `json:"bucketName"`
}

func SetS3ApiConfig(config S3ApiConfig) {
	s3ApiConfig = config
}

func GetS3Configs() *S3ApiConfig {
	return &s3ApiConfig
}
