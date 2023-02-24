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

func ListUser(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter usermodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill()

		store := userstorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := userbiz.NewListUserBiz(store)
		result, err := biz.ListUser(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
