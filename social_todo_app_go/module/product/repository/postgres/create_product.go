package productpostgres

import (
	"context"
	"social_todo_app_go/module/product/domain"
)

func (repo PostgresRepository) CreateProduct(ctx context.Context, p *productdomain.ProductCreateDTO) error {
	if err := repo.db.Create(&p).Error; err != nil {
		return err
	}
	return nil
}
