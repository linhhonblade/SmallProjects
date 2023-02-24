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

func DeleteUser(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := userstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := userbiz.NewDeleteUserBiz(store)
		if err := biz.DeleteUser(c.Request.Context(), id); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
