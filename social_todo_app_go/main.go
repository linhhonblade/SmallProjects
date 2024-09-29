package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"social_todo_app_go/middleware"
	gincategory "social_todo_app_go/module/category/transport/gin"
	ginproduct "social_todo_app_go/module/product/transport/gin"
	"social_todo_app_go/module/upload"
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
	r.Use(middleware.Recover())
	r.Static("/static", "./static")
	v1 := r.Group("/v1")
	{
		v1.PUT("/upload", upload.Upload(db))
		categories := v1.Group("/categories")
		{
			categories.GET("", gincategory.ListCategory(db))
			categories.POST("", gincategory.CreateCategory(db))
			categories.GET("/:id", gincategory.GetCategoryById(db))
			categories.PATCH("/:id", gincategory.UpdateCategoryById(db))
			categories.DELETE("/:id", gincategory.DeleteCategoryById(db))
		}
		products := v1.Group("/products")
		{
			products.GET("", ginproduct.ListProduct(db))
			products.GET("/:id", ginproduct.GetProductById(db))
			products.POST("", ginproduct.CreateProduct(db))
			products.PATCH("/:id", ginproduct.UpdateProductById(db))
			products.DELETE("/:id", ginproduct.DeleteProductById(db))
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
