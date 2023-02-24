package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Attachment struct {
	Id        int    `json:"id" gorm:"column:id;"`
	Url       string `json:"url" gorm:"column:url;"`
	CloudName string `json:"cloud_name,omitempty" gorm:"-"`
	Extension string `json:"extension,omitempty" gorm:"-"`
}

func (Attachment) TableName() string { return "attachment" }

func (a *Attachment) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}
	return json.Marshal(a)
}

func (a *Attachment) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	var att Attachment
	if err := json.Unmarshal(bytes, &att); err != nil {
		return err
	}
	*a = att
	return nil
}
