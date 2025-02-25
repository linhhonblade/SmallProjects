package httpservice

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	sctx "github.com/linhhonblade/service-context"
	"github.com/linhhonblade/service-context/core"
	"net/http"
	"social_todo_app_go/common"
	"social_todo_app_go/module/category/query"
)

type httpService struct {
	sctx sctx.ServiceContext
}

func NewHttpService(sctx sctx.ServiceContext) *httpService {
	return &httpService{sctx: sctx}
}

func (s *httpService) handleRpcListCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		var param struct {
			Ids []uuid.UUID `json:"ids"`
		}

		if err := c.Bind(&param); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithDebug(err.Error()))
			return
		}
		result, err := query.NewCategoryByIdQuery(s.sctx).Execute(c.Request.Context(), param.Ids)
		if err != nil {
			common.WriteErrorResponse(c, err)
		}
		c.JSON(http.StatusOK, core.ResponseData(result))
	}
}
func (s httpService) Routes(g *gin.RouterGroup) {
	category := g.Group("/category")
	rpc := category.Group("/rpc")
	{
		rpc.GET("/query-category-by-id", s.handleRpcListCategory())
	}

}
