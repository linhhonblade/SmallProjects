package httpservice

import (
	"github.com/gin-gonic/gin"
	sctx "github.com/linhhonblade/service-context"
	"github.com/linhhonblade/service-context/core"
	"net/http"
	"social_todo_app_go/common"
	"social_todo_app_go/middleware"
	"social_todo_app_go/module/attachment"
	"social_todo_app_go/module/user/infras/repository"
	"social_todo_app_go/module/user/usecase"
)

type service struct {
	uc         usecase.UseCase
	sctx       sctx.ServiceContext
	authClient middleware.AuthClient
}

func NewUserService(uc usecase.UseCase, sctx sctx.ServiceContext) service {
	return service{uc: uc, sctx: sctx}
}

func (s service) handleRegister() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto usecase.EmailPasswordRegistrationDTO

		if err := c.BindJSON(&dto); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithDebug(err.Error()))
			return
		}

		if err := s.uc.Register(c.Request.Context(), dto); err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}

func (s service) handleLoginEmailPassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto usecase.EmailPasswordLoginDTO
		if err := c.BindJSON(&dto); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithDebug(err.Error()))
			return
		}
		resp, err := s.uc.LoginEmailPassword(c.Request.Context(), dto)
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}
		c.JSON(http.StatusOK, core.ResponseData(resp))
	}
}

func (s service) handleRefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var bodyData struct {
			RefreshToken string `json:"refresh_token"`
		}

		if err := c.BindJSON(&bodyData); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithDebug(err.Error()))
			return
		}
		data, err := s.uc.RefreshToken(c.Request.Context(), bodyData.RefreshToken)
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}
		c.JSON(http.StatusOK, core.ResponseData(data))
	}
}

func (s service) handleChangeAvatar() gin.HandlerFunc {
	return func(c *gin.Context) {
		dto := usecase.SetSingleImageDTO{}
		if err := c.BindJSON(&dto); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithDebug(err.Error()))
			return
		}
		dto.Requester = c.MustGet(common.KeyRequester).(common.Requester)

		// Create Change avatar usecase
		dbCtx := s.sctx.MustGet(common.KeyGormDB).(common.DBContext)
		userRepo := repository.NewUserRepo(dbCtx.GetDB())
		attachmentRepo := attachment.NewRepo(dbCtx.GetDB())
		changeAvtUC := usecase.NewChangeAvtUC(userRepo, userRepo, attachmentRepo)

		if err := changeAvtUC.ChangeAvt(c.Request.Context(), dto); err != nil {
			common.WriteErrorResponse(c, err)
			return
		}
		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}

func (s service) Routes(g *gin.RouterGroup) {
	g.POST("/register", s.handleRegister())
	g.POST("/auth/login", s.handleLoginEmailPassword())
	g.POST("/auth/refresh-token", s.handleRefreshToken())
	g.PATCH("/profile/change-avatar", middleware.RequireAuth(s.authClient), s.handleChangeAvatar())
}

func (s service) SetAuthClient(ac middleware.AuthClient) service {
	s.authClient = ac
	return s
}
