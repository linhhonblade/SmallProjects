package productusecase

import (
	"context"
	"social_todo_app_go/common"
	"social_todo_app_go/module/product/domain"
	"strings"
)

// Use Case Interface

type CreateProductUseCase interface {
	CreateProduct(ctx context.Context, productData *productdomain.ProductCreateDTO) error
}

// Dependencies Injection

func NewCreateProductUseCase(repo CreateProductRepository) CreateNewProductUseCase {
	return CreateNewProductUseCase{
		repo: repo,
	}
}

// Use Case

type CreateNewProductUseCase struct {
	repo CreateProductRepository
}

// Implements Use Case Interface

func (uc CreateNewProductUseCase) CreateProduct(ctx context.Context, productData *productdomain.ProductCreateDTO) error {
	productData.Name = strings.TrimSpace(productData.Name)
	if productData.Name == "" {
		return productdomain.ErrProductNameRequired
	}
	productData.SQLModel = common.GenNewModel()
	err := uc.repo.CreateProduct(ctx, productData)
	if err != nil {
		return err
	}
	return nil
}

// Repository Interface

type CreateProductRepository interface {
	CreateProduct(ctx context.Context, p *productdomain.ProductCreateDTO) error
}
