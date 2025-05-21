package grpcclient

import (
	"context"
	"github.com/google/uuid"
	"social_todo_app_go/module/order/query"
	"social_todo_app_go/proto/product"
)

type productGRPCClient struct {
	client product.ProductClient
}

func NewProductGRPCClient(client product.ProductClient) *productGRPCClient {
	return &productGRPCClient{client: client}
}

func (p productGRPCClient) FindWithIds(ctx context.Context, ids []uuid.UUID) ([]query.ProductDTO, error) {
	idsStr := make([]string, len(ids))
	for i := range ids {
		idsStr[i] = ids[i].String()
	}
	result, err := p.client.GetProductById(ctx, &product.GetProductByIdRequest{Ids: idsStr})
	if err != nil {
		return nil, err
	}
	products := make([]query.ProductDTO, len(result.Data))
	for i := range result.Data {
		products[i] = query.ProductDTO{
			Id:   uuid.MustParse(result.Data[i].Id),
			Name: result.Data[i].Name,
		}
	}
	return products, nil
}
