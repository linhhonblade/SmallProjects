package common

import "gorm.io/gorm"

const (
	KeyRequester = "requester"
	KeyGormDB    = "gorm"
	KeyJWT       = "jwt"
)

type DBContext interface {
	GetDB() *gorm.DB
}
