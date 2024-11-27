package httpservice

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social_todo_app_go/module/user/usecase"
)

type service struct {
	uc usecase.UseCase
}

func NewUserService(uc usecase.UseCase) service {
	return service{uc}
}

func (s service) handleRegister() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto usecase.EmailPasswordRegistrationDTO

		if err := c.BindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := s.uc.Register(c.Request.Context(), dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}

func (s service) Routes(g *gin.RouterGroup) {
	g.POST("/register", s.handleRegister())
}
