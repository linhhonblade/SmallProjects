package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	sctx "github.com/linhhonblade/service-context"
	"github.com/linhhonblade/service-context/component/gormc"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"os"
	"social_todo_app_go/builder"
	"social_todo_app_go/cmd/consumer"
	"social_todo_app_go/common"
	"social_todo_app_go/common/pubsub"
	"social_todo_app_go/component"
	"social_todo_app_go/middleware"
	"social_todo_app_go/module/category/infra/grpcservice"
	productservice "social_todo_app_go/module/product/infras/httpservice"
	"social_todo_app_go/module/upload"
	"social_todo_app_go/module/user/infras/httpservice"
	"social_todo_app_go/module/user/infras/repository"
	"social_todo_app_go/module/user/usecase"
	consumer2 "social_todo_app_go/sub"
)

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
		r.Use(middleware.AllowCors())
		tokenProvider := service.MustGet(common.KeyJWT).(component.TokenProvider)

		r.Use(middleware.Recovery())
		authClient := usecase.NewIntrospectUC(repository.NewUserRepo(db), repository.NewUserSessionPostgresRepo(db), tokenProvider)
		r.Static("/static", "./static")

		v1 := r.Group("/v1")
		{
			v1.PUT("/upload", upload.Upload(db))
			userUC := usecase.NewUCWithBuilder(builder.NewComplexBuilder(builder.NewSimpleBuilder(db, tokenProvider)))
			httpservice.NewUserService(userUC, service).SetAuthClient(authClient).Routes(v1)
		}

		r.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
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

		productService := productservice.NewHttpService(service)
		productService.SetGRPCCategClientConn(cc)
		productService.Routes(v1)

		go consumer2.NewTopicUserChangeAvt(service, service.MustGet(common.KeyLocalPS).(pubsub.PubSub)).Start()

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
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
