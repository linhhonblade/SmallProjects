package gincompany

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-service-go/common"
	"simple-service-go/component"
	"simple-service-go/modules/company/companybiz"
	"simple-service-go/modules/company/companystorage"
	"strconv"
)

func GetCompany(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := companystorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := companybiz.NewFindCompanyBiz(store)
		data, err := biz.GetCompany(c.Request.Context(), id)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
