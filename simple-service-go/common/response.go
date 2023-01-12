package common

type succesRes struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func NewSuccessResponse(data, paging, filter interface{}) *succesRes {
	return &succesRes{Data: data, Paging: paging}
}

func SimpleSuccessResponse(data interface{}) *succesRes {
	return NewSuccessResponse(data, nil, nil)
}
