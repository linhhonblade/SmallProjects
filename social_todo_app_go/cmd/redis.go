package cmd

import (
	sctx "github.com/linhhonblade/service-context"
	"github.com/spf13/cobra"
	"log"
	"social_todo_app_go/common"
	"social_todo_app_go/component"
)

var redisCmd = &cobra.Command{
	Use:   "redis",
	Short: "Connect to Redis service",
	Run: func(cmd *cobra.Command, args []string) {
		service := sctx.NewServiceContext(
			sctx.WithName("redis"),
			sctx.WithComponent(component.NewSRedisComponent(common.KeyRedis)),
		)
		if err := service.Load(); err != nil {
			log.Fatalln(err)
		}
		redis := service.MustGet(common.KeyRedis).(component.RedisClient).GetRedisClient()
		if redis == nil {
			log.Fatalln("Failed to connect to Redis service")
		}
	},
}
