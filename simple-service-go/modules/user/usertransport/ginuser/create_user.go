package ginuser

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-service-go/common"
	"simple-service-go/component"
	"simple-service-go/modules/user/userbiz"
	"simple-service-go/modules/user/usermodel"
	"simple-service-go/modules/user/userstorage"
)

func CreateUser(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := userstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := userbiz.NewCreateUserBiz(store)
		if err := biz.CreateUser(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
