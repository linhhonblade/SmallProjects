package main

import (
	"github.com/gin-gonic/gin"
	sctx "github.com/linhhonblade/service-context"
	"github.com/linhhonblade/service-context/component/gormc"
	"log"
	"net/http"
	"social_todo_app_go/builder"
	"social_todo_app_go/common"
	"social_todo_app_go/component"
	"social_todo_app_go/middleware"
	gincategory "social_todo_app_go/module/category/transport/gin"
	"social_todo_app_go/module/product/controller"
	productusecase "social_todo_app_go/module/product/domain/usecase"
	productservice "social_todo_app_go/module/product/infras/httpservice"
	productpostgres "social_todo_app_go/module/product/repository/postgres"
	ginproduct "social_todo_app_go/module/product/transport/gin"
	"social_todo_app_go/module/upload"
	"social_todo_app_go/module/user/infras/httpservice"
	"social_todo_app_go/module/user/infras/repository"
	"social_todo_app_go/module/user/usecase"
)

func newService() sctx.ServiceContext {
	serviceCtx := sctx.NewServiceContext(
		sctx.WithName("social_todo_app_go"),
		sctx.WithComponent(gormc.NewGormDB(common.KeyGormDB, "app")),
		sctx.WithComponent(component.NewJWT(common.KeyJWT)),
		sctx.WithComponent(component.NewAWSS3Provider(common.KeyAWSS3)),
	)
	return serviceCtx
}

func main() {

	service := newService()
	service.OutEnv()

	if err := service.Load(); err != nil {
		log.Fatalln(err)
	}
	db := service.MustGet(common.KeyGormDB).(common.DBContext).GetDB()
	log.Println("DB Connection: ", db)
	/////////////////////////////////////////////

	r := gin.Default()
	tokenProvider := service.MustGet(common.KeyJWT).(component.TokenProvider)

	r.Use(middleware.Recovery())
	authClient := usecase.NewIntrospectUC(repository.NewUserRepo(db), repository.NewUserSessionPostgresRepo(db), tokenProvider)
	r.Static("/static", "./static")

	// Setup dependencies
	repo := productpostgres.NewPostgresRepository(db)
	useCase := productusecase.NewCreateProductUseCase(repo)
	api := controller.NewAPIController(useCase)
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
			//products.GET("", ginproduct.ListProduct(db))
			products.GET("/:id", ginproduct.GetProductById(db))
			products.POST("", api.CreateProductAPI(db))
			products.PATCH("/:id", ginproduct.UpdateProductById(db))
			products.DELETE("/:id", ginproduct.DeleteProductById(db))
		}

		//userUC := usecase.NewUseCase(repository.NewUserRepo(db), &common.Hasher{}, tokenProvider, repository.NewUserSessionPostgresRepo(db))
		//userUC := usecase.NewUCWithBuilder(builder.NewSimpleBuilder(db, tokenProvider))
		userUC := usecase.NewUCWithBuilder(builder.NewComplexBuilder(builder.NewSimpleBuilder(db, tokenProvider)))
		httpservice.NewUserService(userUC, service).SetAuthClient(authClient).Routes(v1)
	}
	productservice.NewHttpService(service).Routes(v1)

	r.GET("/ping", middleware.RequireAuth(authClient), func(c *gin.Context) {
		requester := c.MustGet(common.KeyRequester).(common.Requester) // cast from any to Requester

		c.JSON(http.StatusOK, gin.H{
			"message":      "pong",
			"requester_id": requester.UserId(),
		})
	})

	r.DELETE("/v1/revoke-token", middleware.RequireAuth(authClient), func(c *gin.Context) {
		repo := repository.NewUserSessionPostgresRepo(db)
		requester := c.MustGet(common.KeyRequester).(common.Requester) // cast from any to Requester
		if err := repo.Delete(c.Request.Context(), requester.TokenId()); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	})

	if err := r.Run(":3000"); err != nil {
		log.Fatalln(err)
	}
}
