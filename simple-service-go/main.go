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
	"simple-service-go/modules/user/usertransport/ginuser"
)

// CREATE TABLE "public"."res_users" (
// "id" int4 NOT NULL DEFAULT nextval('res_users_id_seq'::regclass),
// "login" bpchar NOT NULL,
// "password" bpchar,
// PRIMARY KEY ("id")
// );

type User struct {
	Id       int    `json:"id,omitempty" gorm:"column:id"`
	Login    string `json:"login" gorm:"column:login"`
	Password string `json:"password" gorm:"column:password"`
	Lang     string `json:"lang" gorm:"column:lang"`
}

func (User) TableName() string {
	return "res_users"
}

type UserUpdate struct {
	Login    *string `json:"login" gorm:"column:login"`
	Password *string `json:"password" gorm:"column:password"`
	Lang     *string `json:"lang" gorm:"column:lang"`
}

func (UserUpdate) TableName() string {
	return User{}.TableName()
}

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

	users := r.Group("/res_users")
	{
		users.POST("", ginuser.CreateUser(appCtx))

		users.GET("/:id", ginuser.GetUser(appCtx))

		users.GET("", ginuser.ListUser(appCtx))

		users.PATCH("/:id", ginuser.UpdateUser(appCtx))

		users.DELETE("/:id", ginuser.DeleteUser(appCtx))
	}
	return r.Run()
}
