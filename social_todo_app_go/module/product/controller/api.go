package controller

import (
	"context"
	productdomain "social_todo_app_go/module/product/domain"
)

type CreateProductUseCase interface {
	CreateProduct(ctx context.Context, prod *productdomain.ProductCreateDTO) error
}

type APIController struct {
	createUseCase CreateProductUseCase
}

func NewAPIController(createUseCase CreateProductUseCase) APIController {
	return APIController{createUseCase: createUseCase}
}
