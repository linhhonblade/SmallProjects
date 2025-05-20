package httpservice

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	sctx "github.com/linhhonblade/service-context"
	"github.com/linhhonblade/service-context/core"
	"net/http"
	"social_todo_app_go/common"
	"social_todo_app_go/middleware"
	"social_todo_app_go/module/order/query"
)

type httpService struct {
	sctx       sctx.ServiceContext
	authClient middleware.AuthClient
}

func NewHttpService(sctx sctx.ServiceContext) *httpService {
	return &httpService{sctx: sctx}
}

// handleCreateOrder handles POST /orders
// @Summary Create a new order
// @Description Create a new order with the provided data
// @Tags order
// @Accept json
// @Produce json
// @Param data body query.OrderDTO true "Order data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Security  Bearer
// @Router /v1/orders [post]
func (s *httpService) handleCreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implement the logic to create an order
		var data query.OrderDTO
		requester := c.MustGet(common.KeyRequester).(common.Requester)
		if err := c.ShouldBindJSON(&data); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithDebug(err.Error()))
			return
		}
		result, err := query.NewCreateOneOrderQuery(s.sctx, requester).Execute(c.Request.Context(), &data)
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}
		c.JSON(http.StatusOK, core.ResponseData(result.Id))
	}
}

// handleGetOrder handles GET /orders/:id
// @Summary Get order by ID
// @Description Retrieve an order by its UUID
// @Tags order
// @Accept json
// @Produce json
// @Param id path string true "Order ID (UUID)"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Security  Bearer
// @Router /v1/orders/{id} [get]
func (s *httpService) handleGetOrderById() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implement the logic to get an order
		orderIDStr := c.Param("id")
		orderID, err := uuid.Parse(orderIDStr)
		requester := c.MustGet(common.KeyRequester).(common.Requester)
		if err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithDebug("invalid order id"))
			return
		}
		result, err := query.NewGetOrderByIdQuery(s.sctx, requester).Execute(c.Request.Context(), orderID)
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}
		c.JSON(http.StatusOK, core.ResponseData(result))
	}
}

// handleListOrder handles GET /orders
// @Summary List orders
// @Description Get a list of orders for the authenticated user, with optional filters and pagination
// @Tags order
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Param status query string false "Order status filter"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Security  Bearer
// @Router /v1/orders [get]
func (s *httpService) handleListOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implement the logic to list orders
		var param query.ListOrderParam
		if err := c.Bind(&param); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithDebug(err.Error()))
			return
		}
		requester := c.MustGet(common.KeyRequester).(common.Requester)
		result, err := query.NewListOrderQuery(s.sctx, requester).Execute(c.Request.Context(), &param)
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}
		c.JSON(http.StatusOK, core.SuccessResponse(result, param.Paging, param.ListOrderFilter))
	}
}

func (s *httpService) Routes(g *gin.RouterGroup) {
	orders := g.Group("/orders")
	orders.POST("/", middleware.RequireAuth(s.authClient), s.handleCreateOrder())
	orders.GET("/:id", middleware.RequireAuth(s.authClient), s.handleGetOrderById())
	orders.GET("/", middleware.RequireAuth(s.authClient), s.handleListOrder())
}

func (s *httpService) SetAuthClient(ac middleware.AuthClient) *httpService {
	s.authClient = ac
	return s
}
