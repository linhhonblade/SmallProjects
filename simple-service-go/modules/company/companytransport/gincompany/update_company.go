package gincompany

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-service-go/common"
	"simple-service-go/component"
	"simple-service-go/modules/company/companybiz"
	"simple-service-go/modules/company/companymodel"
	"simple-service-go/modules/company/companystorage"
	"strconv"
)

func UpdateCompany(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		var data companymodel.CompanyUpdate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := companystorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := companybiz.NewUpdateCompanyBiz(store)
		if err := biz.UpdateData(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
