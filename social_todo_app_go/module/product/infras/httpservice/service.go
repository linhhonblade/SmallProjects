package httpservice

import (
	"github.com/gin-gonic/gin"
	sctx "github.com/linhhonblade/service-context"
	"github.com/linhhonblade/service-context/core"
	"net/http"
	"social_todo_app_go/common"
	"social_todo_app_go/module/product/query"
)

type httpService struct {
	sctx sctx.ServiceContext
}

func NewHttpService(sctx sctx.ServiceContext) *httpService {
	return &httpService{sctx: sctx}
}

func (s *httpService) handleListProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var param query.ListProductParam
		if err := c.Bind(&param); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithDebug(err.Error()))
			return
		}
		result, err := query.NewListProductQuery(s.sctx).Execute(c.Request.Context(), &param)
		if err != nil {
			common.WriteErrorResponse(c, err)
		}
		c.JSON(http.StatusOK, core.SuccessResponse(result, param.Paging, param.ListProductFilter))
	}
}
func (s httpService) Routes(g *gin.RouterGroup) {
	products := g.Group("/products")
	products.GET("/", s.handleListProduct())
}
