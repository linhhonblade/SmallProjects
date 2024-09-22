package ginitem

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"social_todo_app_go/common"
	"social_todo_app_go/module/item/biz"
	"social_todo_app_go/module/item/model"
	"social_todo_app_go/module/item/storage"
	"strconv"
)

func UpdateItemById(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var dataUpdate model.TodoItemUpdate
		if err := c.ShouldBind(&dataUpdate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := storage.NewSQLStore(db)
		business := biz.NewUpdateItemBiz(store)
		err = business.UpdateItemById(c.Request.Context(), id, &dataUpdate)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
