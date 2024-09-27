package ginproduct

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"social_todo_app_go/common"
	"social_todo_app_go/module/product/biz"
	"social_todo_app_go/module/product/model"
	"social_todo_app_go/module/product/storage"
)

func CreateProduct(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var productData model.ProductCreation
		if err := c.ShouldBind(&productData); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		store := storage.NewSQLStore(db)
		business := biz.NewCreateProductBiz(store)
		if err := business.CreateNewProduct(c.Request.Context(), &productData); err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(productData.Id))
	}
}
