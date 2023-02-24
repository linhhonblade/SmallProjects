package companymodel

import "simple-service-go/common"

const EntityName = "Company"

type Company struct {
	common.SQLModel `json:",inline"`
	Name            string              `json:"name" gorm:"column:name"`
	Fullname        string              `json:"fullname" gorm:"column:fullname"`
	Logo            *common.Attachment  `json:"logo" gorm:"column:logo"`
	Cover           *common.Attachments `json:"cover" gorm:"column:cover"`
	Favicon         *common.Attachment  `json:"favicon" gorm:"column:favicon"`
}

func (Company) TableName() string {
	return "company"
}

type CompanyUpdate struct {
	Name     *string             `json:"name" gorm:"column:name"`
	Fullname *string             `json:"fullname" gorm:"column:fullname"`
	Logo     *common.Attachment  `json:"logo" gorm:"column:logo"`
	Cover    *common.Attachments `json:"cover" gorm:"column:cover"`
	Favicon  *common.Attachment  `json:"favicon" gorm:"column:favicon"`
}

func (CompanyUpdate) TableName() string {
	return Company{}.TableName()
}

type CompanyCreate struct {
	Name     *string             `json:"name" gorm:"column:name"`
	Fullname *string             `json:"fullname" gorm:"column:fullname"`
	Logo     *common.Attachment  `json:"logo" gorm:"column:logo"`
	Cover    *common.Attachments `json:"cover" gorm:"column:cover"`
	Favicon  *common.Attachment  `json:"favicon" gorm:"column:favicon"`
}

func (CompanyCreate) TableName() string {
	return Company{}.TableName()
}
