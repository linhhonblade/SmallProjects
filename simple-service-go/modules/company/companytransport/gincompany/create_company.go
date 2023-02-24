package gincompany

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-service-go/common"
	"simple-service-go/component"
	"simple-service-go/modules/company/companybiz"
	"simple-service-go/modules/company/companymodel"
	"simple-service-go/modules/company/companystorage"
)

func CreateCompany(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data companymodel.CompanyCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := companystorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := companybiz.NewCreateCompanyBiz(store)
		if err := biz.CreateCompany(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
