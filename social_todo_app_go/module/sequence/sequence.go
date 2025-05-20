package sequence

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Sequence struct {
	Prefix string `gorm:"primaryKey"`
	Date   string `gorm:"primaryKey"`
	Value  int
}

func Next(db *gorm.DB, prefix string, padding, step int) (string, error) {
	date := time.Now().Format("060102") // YYMMDD
	var value int
	err := db.Raw(`
		INSERT INTO sequence (prefix, date, value)
		VALUES (?, ?, ?)
		ON CONFLICT (prefix, date)
		DO UPDATE SET value = sequence.value + ?
		RETURNING value
	`, prefix, date, step, step).Scan(&value).Error
	if err != nil {
		return "", err
	}
	name := fmt.Sprintf("%s_%0*d", prefix, padding, value)
	return name, nil
}

func (s *Sequence) TableName() string {
	return "sequence"
}
