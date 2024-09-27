package model

type Filter struct {
	Status  string `json:"status,omitempty" form:"status"`
	CategId int    `json:"categ_id,omitempty" form:"categ_id"`
	Type    string `json:"type,omitempty" form:"type"`
}
