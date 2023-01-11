package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"simple-service-go/modules/user/usertransport/ginuser"
	"strconv"
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
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	users := r.Group("/res_users")
	{
		users.POST("", ginuser.CreateUser(db))

		users.GET("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(401, gin.H{
					"error": err.Error(),
				})
				return
			}

			var data User
			if err := db.Where("id = ?", id).First(&data).Error; err != nil {
				c.JSON(401, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, data)
		})

		users.GET("", func(c *gin.Context) {
			var data []User
			type Filter struct {
				Lang string `json:"lang" form:"lang"`
			}
			var filter Filter

			if err := c.ShouldBind(&filter); err != nil {
				c.JSON(401, gin.H{
					"error": err.Error(),
				})
				return
			}

			newDb := db
			if filter.Lang != "" {
				newDb = db.Where("lang = ?", filter.Lang)
			}

			if err := newDb.Find(&data).Error; err != nil {
				c.JSON(401, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, data)
		})

		users.PATCH("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))

			if err != nil {
				c.JSON(401, gin.H{
					"error": err.Error(),
				})
				return
			}

			var data UserUpdate

			if err := c.ShouldBind(&data); err != nil {
				c.JSON(401, gin.H{
					"error": err.Error(),
				})
				return
			}

			if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
				c.JSON(401, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{"ok": 1})
		})
	}
	return r.Run()
}
