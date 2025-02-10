package common

import (
	"fmt"
	"github.com/google/uuid"
)

type Attachment struct {
	Id              uuid.UUID `json:"id"`
	Title           string    `json:"title"`
	FileName        string    `json:"file_name"`
	FileUrl         string    `json:"file_url" gorm:"-"`
	FileSize        int       `json:"file_size"`
	FileType        string    `json:"file_type"`
	StorageProvider string    `json:"storage_provider"`
	Status          string    `json:"status"`
}

func (att *Attachment) SetCDNDomain(domain string) {
	att.FileUrl = fmt.Sprintf("%s/%s", domain, att.FileName)
}
