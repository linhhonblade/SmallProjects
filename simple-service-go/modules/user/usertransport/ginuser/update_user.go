package ginuser

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-service-go/common"
	"simple-service-go/component"
	"simple-service-go/modules/user/userbiz"
	"simple-service-go/modules/user/usermodel"
	"simple-service-go/modules/user/userstorage"
	"strconv"
)

func UpdateUser(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		var data usermodel.UserUpdate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := userstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := userbiz.NewUpdateUserBiz(store)
		if err := biz.UpdateData(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
