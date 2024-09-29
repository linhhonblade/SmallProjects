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

func ListCategory(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var queryString struct {
			common.Paging
			model.Filter
		}
		if err := c.ShouldBind(&queryString); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		queryString.Paging.Process()
		store := storage.NewSQLStore(db)
		business := biz.NewListCategoryBiz(store)
		res, err := business.ListCategory(c.Request.Context(), &queryString.Filter, &queryString.Paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(res, queryString.Paging, queryString.Filter))
	}
}
