package httpservice

import (
	"github.com/gin-gonic/gin"
	sctx "github.com/linhhonblade/service-context"
	"github.com/linhhonblade/service-context/core"
	"google.golang.org/grpc"
	"net/http"
	"social_todo_app_go/common"
	"social_todo_app_go/module/product/query"
	"social_todo_app_go/module/product/repository/grpcclient"
	"social_todo_app_go/proto/category"
)

type httpService struct {
	sctx                sctx.ServiceContext
	grpcCategClientConn grpc.ClientConnInterface
}

func NewHttpService(sctx sctx.ServiceContext) *httpService {
	return &httpService{sctx: sctx}
}

// handleListProduct handles GET /products
// @Summary Get Product List
// @Description Get list of products with optional filters and pagination
// @Tags product
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Param name query string false "Product name filter"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /v1/products [get]
func (s *httpService) handleListProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var param query.ListProductParam
		if err := c.Bind(&param); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithDebug(err.Error()))
			return
		}

		categRepo := grpcclient.NewCategGRPCClient(category.NewCategoryClient(s.grpcCategClientConn))
		result, err := query.NewListProductQuery(s.sctx, categRepo).Execute(c.Request.Context(), &param)
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}
		c.JSON(http.StatusOK, core.SuccessResponse(result, param.Paging, param.ListProductFilter))
	}
}

// handleCreateProduct handles POST /products
// @Summary Create a new product
// @Description Create a new product with the provided data
// @Tags product
// @Accept json
// @Produce json
// @Param data body query.ProductDTO true "Product data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /v1/products [post]
func (s *httpService) handleCreateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data query.ProductDTO
		if err := c.Bind(&data); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithDebug(err.Error()))
			return
		}

		result, err := query.NewCreateOneProductQuery(s.sctx).Execute(c.Request.Context(), &data)
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}
		c.JSON(http.StatusOK, core.ResponseData(result.Id))
	}
}

func (s *httpService) Routes(g *gin.RouterGroup) {
	products := g.Group("/products")
	products.GET("/", s.handleListProduct())
	products.POST("/", s.handleCreateProduct())
}

func (s *httpService) SetGRPCCategClientConn(cc grpc.ClientConnInterface) {
	s.grpcCategClientConn = cc
}
