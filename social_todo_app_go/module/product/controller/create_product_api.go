package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	productdomain "social_todo_app_go/module/product/domain"
)

func (api APIController) CreateProductAPI(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var productData productdomain.ProductCreateDTO
		if err := c.ShouldBind(&productData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := api.createUseCase.CreateProduct(c.Request.Context(), &productData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Response to client
		c.JSON(http.StatusCreated, gin.H{"data": productData.Id})
	}

}
