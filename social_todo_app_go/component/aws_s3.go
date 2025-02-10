package component

import (
	"bytes"
	"context"
	"flag"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	sctx "github.com/linhhonblade/service-context"
	"net/http"
)

type UploadProvider interface {
	SaveUploadedFile(ctx context.Context, data []byte, dst string) error
}

type s3Provider struct {
	id         string
	bucketName string
	region     string
	apiKey     string
	secret     string
	domain     string
	session    *session.Session
}

func NewAWSS3Provider(id string) *s3Provider {
	return &s3Provider{id: id}
}

func (p *s3Provider) ID() string {
	return p.id
}

func (p *s3Provider) InitFlags() {
	flag.StringVar(
		&p.secret,
		"aws-s3-secret",
		"",
		"AWS S3 Secret Key",
	)
	flag.StringVar(
		&p.apiKey,
		"aws-s3-api-key",
		"",
		"AWS S3 API Key",
	)
	flag.StringVar(
		&p.bucketName,
		"aws-s3-bucket-name",
		"",
		"AWS S3 Bucket Name",
	)
	flag.StringVar(
		&p.region,
		"aws-s3-region",
		"ap-southeast-1",
		"AWS S3 Region Name",
	)
	flag.StringVar(
		&p.domain,
		"cdn-domain",
		"",
		"CDN Domain",
	)
}

func (p *s3Provider) Activate(_ sctx.ServiceContext) error {
	s3Session, err := session.NewSession(&aws.Config{
		Region:      aws.String(p.region),
		Credentials: credentials.NewStaticCredentials(p.apiKey, p.secret, ""),
	})
	if err != nil {
		return err
	}
	p.session = s3Session
	return nil
}

func (p *s3Provider) Stop() error {
	return nil
}

func (p *s3Provider) SaveUploadedFile(ctx context.Context, data []byte, dst string) error {
	fileBytes := bytes.NewReader(data)
	fileType := http.DetectContentType(data)
	_, err := s3.New(p.session).PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(p.bucketName),
		Key:         aws.String(dst),
		ACL:         aws.String("private"),
		ContentType: aws.String(fileType),
		Body:        fileBytes})
	if err != nil {
		return err
	}
	return nil
}

func (p *s3Provider) GetDomain() string {
	return p.domain
}

func (p *s3Provider) GetName() string {
	return "aws_s3"
}
