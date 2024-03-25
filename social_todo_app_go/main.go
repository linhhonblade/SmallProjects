package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"social_todo_app_go/common"
	"social_todo_app_go/module/item/model"
	ginitem "social_todo_app_go/module/item/transport/gin"
)

func main() {
	dsn := os.Getenv("DB_CONN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if os.Getenv("DB_DEBUG") == "true" {
		db = db.Debug()
	}
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("DB Connection: ", db)
	/////////////////////////////////////////////

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.GET("", ginitem.ListItem(db))
			items.POST("", ginitem.CreateItem(db))
			items.GET("/:id", ginitem.GetItemById(db))
			items.PATCH("/:id", ginitem.UpdateItemById(db))
			items.DELETE("/:id", ginitem.DeleteItemById(db))
		}
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	if err := r.Run(":3000"); err != nil {
		log.Fatalln(err)
	}
}

func ListItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		paging.Process()

		var result []model.TodoItem
		// Prevent count(*)
		// Instead, use count("id")

		db = db.Where("active = ?", true)
		if err := db.Table(model.TodoItem{}.TableName()).Select("id").Count(&paging.Total).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := db.Table(model.TodoItem{}.TableName()).
			Select("id, title, description, status").
			Offset((paging.Page - 1) * paging.Limit).
			Limit(paging.Limit).
			Order("id desc").
			Find(&result).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
