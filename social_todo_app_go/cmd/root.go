package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	sctx "github.com/linhhonblade/service-context"
	"github.com/linhhonblade/service-context/component/gormc"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"os"
	"social_todo_app_go/builder"
	"social_todo_app_go/cmd/consumer"
	"social_todo_app_go/common"
	"social_todo_app_go/common/pubsub"
	"social_todo_app_go/component"
	_ "social_todo_app_go/docs"
	"social_todo_app_go/middleware"
	"social_todo_app_go/module/category/infras/grpcservice"
	categoryservice "social_todo_app_go/module/category/infras/httpservice"
	orderservice "social_todo_app_go/module/order/infras/httpservice"
	productservice "social_todo_app_go/module/product/infras/httpservice"
	"social_todo_app_go/module/upload"
	"social_todo_app_go/module/user/infras/httpservice"
	"social_todo_app_go/module/user/infras/repository"
	"social_todo_app_go/module/user/usecase"
	consumer2 "social_todo_app_go/sub"
)

// @title G11 Golang Course API
// @version 1.0
// @description API demo dùng Gin + Swagger
// @host localhost:3000
// @SecurityDefinitions.apiKey Bearer
// @in header
// @name Authorization
// @BasePath /v1

func newService() sctx.ServiceContext {
	return sctx.NewServiceContext(
		sctx.WithName("social_todo_app_go"),
		sctx.WithComponent(gormc.NewGormDB(common.KeyGormDB, "app")),
		sctx.WithComponent(component.NewJWT(common.KeyJWT)),
		sctx.WithComponent(component.NewAWSS3Provider(common.KeyAWSS3)),
		sctx.WithComponent(component.NewConfig(common.KeyConfig)),
		sctx.WithComponent(component.NewNATSComponent(common.KeyNATS)),
		sctx.WithComponent(pubsub.NewLocalPubSub(common.KeyLocalPS)),
	)
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start main service",
	Run: func(cmd *cobra.Command, args []string) {
		service := newService()
		if err := service.Load(); err != nil {
			log.Fatalln(err)
		}
		db := service.MustGet(common.KeyGormDB).(common.DBContext).GetDB()
		log.Println("DB Connection: ", db)
		/////////////////////////////////////////////

		r := gin.Default()
		// Thêm route Swagger
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		//r.Use(middleware.AllowCors())
		tokenProvider := service.MustGet(common.KeyJWT).(component.TokenProvider)

		r.Use(middleware.Recovery())
		authClient := usecase.NewIntrospectUC(repository.NewUserRepo(db), repository.NewUserSessionPostgresRepo(db), tokenProvider)
		r.Static("/static", "./static")

		v1 := r.Group("/v1")
		{
			// @Summary Upload file
			// @Description Upload a file to the server
			// @Tags upload
			// @Accept multipart/form-data
			// @Produce json
			// @Param file formData file true "File to upload"
			// @Success 200 {object} map[string]interface{}
			// @Failure 400 {object} map[string]string
			// @Router /v1/upload [put]
			v1.PUT("/upload", upload.Upload(db))

			userUC := usecase.NewUCWithBuilder(builder.NewComplexBuilder(builder.NewSimpleBuilder(db, tokenProvider)))
			httpservice.NewUserService(userUC, service).SetAuthClient(authClient).Routes(v1)
		}

		// Health check service
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		// Revoke token
		r.DELETE("/v1/revoke-token", middleware.RequireAuth(authClient), func(c *gin.Context) {
			repo := repository.NewUserSessionPostgresRepo(db)
			requester := c.MustGet(common.KeyRequester).(common.Requester) // cast from any to Requester
			if err := repo.Delete(c.Request.Context(), requester.TokenId()); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"data": true})
		})

		// Run gRPC category server
		go func() {
			_ = grpcservice.NewCategoryGRPCService(service.MustGet(common.KeyConfig).(interface{ GetPortGRPCCategory() int }).GetPortGRPCCategory(), service).Start()
		}()
		opts := grpc.WithTransportCredentials(insecure.NewCredentials())
		grpcCategServerUrl := service.MustGet(common.KeyConfig).(interface{ GetUrlGRPCCategoryServer() string }).GetUrlGRPCCategoryServer()
		cc, err := grpc.NewClient(grpcCategServerUrl, opts)
		if err != nil {
			log.Fatalln(err)
		}

		// Product http service
		productService := productservice.NewHttpService(service)
		productService.SetGRPCCategClientConn(cc)
		productService.Routes(v1)

		// Product category http service
		categoryService := categoryservice.NewHttpService(service)
		categoryService.Routes(v1)

		// Order http service
		orderService := orderservice.NewHttpService(service).SetAuthClient(authClient)
		orderService.Routes(v1)

		// consumer pub sub
		go consumer2.NewTopicUserChangeAvt(service, service.MustGet(common.KeyLocalPS).(pubsub.PubSub)).Start()

		// Start http service
		if err := r.Run(":3000"); err != nil {
			log.Fatalln(err)
		}
	},
}

func Execute() {
	rootCmd.AddCommand(outEnvCmd)
	consumerCmd := &cobra.Command{Use: "consumer", Short: "Start consumer"}
	consumerCmd.AddCommand(consumer.SetImgActiveAfterChangeAvtCmd)
	rootCmd.AddCommand(consumerCmd)

	rootCmd.AddCommand(loadTestCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
