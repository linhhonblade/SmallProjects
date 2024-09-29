package common

import (
	"github.com/google/uuid"
	"time"
)

type SQLModel struct {
	Id         uuid.UUID `json:"id" gorm:"column:id;"`
	CreateDate time.Time `json:"create_date" gorm:"column:create_date;"`
	WriteDate  time.Time `json:"write_date" gorm:"column:write_date;"`
	Status     string    `json:"status" gorm:"column:status;"`
}

func GenNewModel() SQLModel {
	now := time.Now().UTC()
	return SQLModel{
		Id:         uuid.New(),
		CreateDate: now,
		WriteDate:  now,
		Status:     "active",
	}
}
