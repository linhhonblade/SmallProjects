package common

import "time"

type SQLModel struct {
	Id         int        `json:"id" gorm:"column:id;"`
	CreateDate *time.Time `json:"create_date" gorm:"column:create_date;"`
	WriteDate  *time.Time `json:"write_date" gorm:"column:write_date;"`
}
