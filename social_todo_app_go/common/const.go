package common

import "gorm.io/gorm"

const (
	KeyRequester = "requester"
	KeyGormDB    = "gorm"
	KeyJWT       = "jwt"
	KeyAWSS3     = "aws_s3"
	KeyConfig    = "config"
)

type DBContext interface {
	GetDB() *gorm.DB
}
