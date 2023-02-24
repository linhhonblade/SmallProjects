package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"simple-service-go/component"
	"simple-service-go/middleware"
	"simple-service-go/modules/company/companytransport/gincompany"
	"simple-service-go/modules/user/usertransport/ginuser"
)

func main() {
	dsn := os.Getenv("DBConnectionStr")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	if err := runService(db); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB) error {
	appCtx := component.NewAppContext(db)
	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	users := r.Group("/users")
	{
		users.POST("", ginuser.CreateUser(appCtx))
		users.GET("/:id", ginuser.GetUser(appCtx))
		users.GET("", ginuser.ListUser(appCtx))
		users.PATCH("/:id", ginuser.UpdateUser(appCtx))
		users.DELETE("/:id", ginuser.DeleteUser(appCtx))
	}

	company := r.Group("/company")
	{
		company.POST("", gincompany.CreateCompany(appCtx))
		company.GET("/:id", gincompany.GetCompany(appCtx))
		company.PATCH(":id", gincompany.UpdateCompany(appCtx))
	}
	return r.Run()
}
