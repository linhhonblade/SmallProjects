package ginuser

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-service-go/common"
	"simple-service-go/component"
	"simple-service-go/modules/user/userbiz"
	"simple-service-go/modules/user/userstorage"
	"strconv"
)

func GetUser(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := userstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := userbiz.NewFindUserBiz(store)
		data, err := biz.GetUser(c.Request.Context(), id)
		if err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
