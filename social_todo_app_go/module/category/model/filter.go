package model

type Filter struct {
	Status string `json:"status,omitempty" form:"status"`
	Name   string `json:"name,omitempty" form:"name"`
}
