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

// handleRpcListCategory handles GET /category/rpc/query-category-by-id
// @Summary Get categories by IDs
// @Description Get a list of categories by their UUIDs
// @Tags category
// @Accept json
// @Produce json
// @Param ids query []string true "List of category UUIDs"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /v1/category/rpc/query-category-by-id [get]
func (s *httpService) handleRpcListCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		ids := c.QueryArray("ids")
		var uuidIds []uuid.UUID
		for _, id := range ids {
			uid, err := uuid.Parse(id)
			if err != nil {
				common.WriteErrorResponse(c, core.ErrBadRequest.WithDebug("Invalid UUID: "+id))
				return
			}
			uuidIds = append(uuidIds, uid)
			result, err := query.NewCategoryByIdQuery(s.sctx).Execute(c.Request.Context(), uuidIds)
			if err != nil {
				common.WriteErrorResponse(c, err)
			}
			c.JSON(http.StatusOK, core.ResponseData(result))
		}
	}
}

// handleListCategory handles GET /category
// @Summary Get Category List
// @Description Get list of categories with optional filters and pagination
// @Tags category
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Param name query string false "Category name filter"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /v1/category/ [get]
func (s *httpService) handleListCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		var param query.ListCategoryParam
		if err := c.Bind(&param); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithDebug(err.Error()))
			return
		}
		result, err := query.NewListCategoryQuery(s.sctx).Execute(c.Request.Context(), &param)
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}
		c.JSON(http.StatusOK, core.SuccessResponse(result, param.Paging, param.ListCategoryFilter))
	}
}

// handleCreateCategory handles POST /category
// @Summary Create a new category
// @Description Create a new category with the provided data
// @Tags category
// @Accept json
// @Produce json
// @Param data body query.CategoryDTO true "Category data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /v1/category/ [post]
func (s *httpService) handleCreateCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data query.CategoryDTO
		if err := c.Bind(&data); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithDebug(err.Error()))
			return
		}
		result, err := query.NewCreateOneProductCategoryQuery(s.sctx).Execute(c.Request.Context(), &data)
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}
		c.JSON(http.StatusOK, core.ResponseData(result.Id))
	}
}

func (s httpService) Routes(g *gin.RouterGroup) {
	category := g.Group("/category")
	rpc := category.Group("/rpc")
	{
		rpc.GET("/query-category-by-id", s.handleRpcListCategory())
	}
	{
		category.GET("/", s.handleListCategory())
		category.POST("/", s.handleCreateCategory())
	}

}
