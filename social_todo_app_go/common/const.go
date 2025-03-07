package common

import "gorm.io/gorm"

const (
	KeyRequester = "requester"
	KeyGormDB    = "gorm"
	KeyJWT       = "jwt"
	KeyAWSS3     = "aws_s3"
	KeyConfig    = "config"
	KeyNATS      = "nats"
	KeyLocalPS   = "local_ps"

	TopicUserAvtChanged = "user.avatar.changed"
)

type DBContext interface {
	GetDB() *gorm.DB
}
