package aws

import (
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3 struct {
	sess                        *session.Session
	uploader                    *s3manager.Uploader
	downloader                  *s3manager.Downloader
	svc                         *s3.S3
	S3AdminProfilePictureBucket string
}

func NewS3(sess *session.Session, env string) *S3 {
	return &S3{
		sess:                        sess,
		uploader:                    s3manager.NewUploader(sess),
		downloader:                  s3manager.NewDownloader(sess),
		svc:                         s3.New(sess),
		S3AdminProfilePictureBucket: fmt.Sprintf("%s-admin-profile-pick", env),
	}
}

func (s *S3) UploadToBucket(content io.Reader, bucket, key, contentType string) (string, error) {
	res, err := s.uploader.Upload(&s3manager.UploadInput{
		Body:        content,
		Bucket:      aws.String(bucket),
		ContentType: aws.String(contentType),
		Key:         aws.String(key),
	})

	fmt.Println(res.Location)

	if err != nil {
		return "", fmt.Errorf("error uploading file: %v", err)
	}

	return res.Location, nil
}

func (s *S3) UpdateObject(content io.Reader, bucket, key, contentType string) error {
	_, err := s.UploadToBucket(content, bucket, key, contentType)
	return err
}

func (s *S3) DeleteObject(bucket, key string) error {
	if _, err := s.svc.DeleteObject(&s3.DeleteObjectInput{
		Key:    aws.String(key),
		Bucket: aws.String(bucket),
	}); err != nil {
		return fmt.Errorf("failed to delete object: %v", err)
	}

	return nil
}

func (s *S3) GeneratePresignedUrl(bucket, key string) (string, error) {
	req, _ := s.svc.GetObjectRequest(&s3.GetObjectInput{
		Key:    aws.String(key),
		Bucket: aws.String(bucket),
	})

	url, err := req.Presign(15 * time.Minute)
	if err != nil {
		return "", fmt.Errorf("failed to generate url %s for key %s: %v", bucket, key, err)
	}

	return url, nil
}
