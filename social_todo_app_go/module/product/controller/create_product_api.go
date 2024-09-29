package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	productdomain "social_todo_app_go/module/product/domain"
	productusecase "social_todo_app_go/module/product/domain/usecase"
	productpostgres "social_todo_app_go/module/product/repository/postgres"
)

func CreateProductAPI(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var productData productdomain.ProductCreateDTO
		if err := c.ShouldBind(&productData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		repo := productpostgres.NewPostgresRepository(db)
		useCase := productusecase.NewCreateProductUseCase(repo)
		if err := useCase.CreateProduct(c.Request.Context(), &productData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Response to client
		c.JSON(http.StatusCreated, gin.H{"data": productData.Id})
	}

}
