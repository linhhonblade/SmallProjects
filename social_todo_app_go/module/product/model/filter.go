package model

type Filter struct {
	Status     string `json:"status,omitempty" form:"status"`
	CategoryId int    `json:"category_id,omitempty" form:"category_id"`
	Type       string `json:"type,omitempty" form:"type"`
}
