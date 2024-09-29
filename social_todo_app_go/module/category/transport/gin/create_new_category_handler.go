package gincategory

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"social_todo_app_go/common"
	"social_todo_app_go/module/category/biz"
	"social_todo_app_go/module/category/model"
	"social_todo_app_go/module/category/storage"
)

func CreateCategory(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var categoryData model.CategoryCreation
		if err := c.ShouldBind(&categoryData); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		store := storage.NewSQLStore(db)
		business := biz.NewCreateCategoryBiz(store)
		if err := business.CreateNewCategory(c.Request.Context(), &categoryData); err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(categoryData.Id))
	}
}
