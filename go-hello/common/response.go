package common

type successResponse struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func NewSuccessResponse(data, paging, filter interface{}) *successResponse {
	return &successResponse{Data: data, Paging: paging}
}

func SimpleSuccessResponse(data interface{}) *successResponse {
	return NewSuccessResponse(data, nil, nil)
}
