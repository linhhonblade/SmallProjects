package ginuser

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"simple-service-go/modules/user/userbiz"
	"simple-service-go/modules/user/usermodel"
	"simple-service-go/modules/user/userstorage"
)

func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := userstorage.NewSQLStore(db)
		biz := userbiz.NewCreateUserBiz(store)
		if err := biz.CreateUser(c.Request.Context(), &data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
		}
		c.JSON(http.StatusOK, data)
	}
}
