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

func ListProduct(db *gorm.DB) func(ctx *gin.Context) {
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
		business := biz.NewListProductBiz(store)
		res, err := business.ListProduct(c.Request.Context(), &queryString.Filter, &queryString.Paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(res, queryString.Paging, queryString.Filter))
	}
}
