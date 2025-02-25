package grpcclient

import (
	"context"
	"github.com/google/uuid"
	"social_todo_app_go/module/product/query"
	"social_todo_app_go/proto/category"
)

type categGRPCClient struct {
	client category.CategoryClient
}

func NewCategGRPCClient(client category.CategoryClient) *categGRPCClient {
	return &categGRPCClient{client: client}
}

func (c *categGRPCClient) FindWithIds(ctx context.Context, ids []uuid.UUID) ([]query.CategoryDTO, error) {
	idsStr := make([]string, len(ids))
	for i := range ids {
		idsStr[i] = ids[i].String()
	}
	result, err := c.client.GetCategoryById(ctx, &category.GetCategoryByIdRequest{Ids: idsStr})
	if err != nil {
		return nil, err
	}
	categories := make([]query.CategoryDTO, len(result.Data))
	for i := range result.Data {
		categories[i] = query.CategoryDTO{
			Id:   uuid.MustParse(result.Data[i].Id),
			Name: result.Data[i].Name,
		}
	}
	return categories, nil
}
