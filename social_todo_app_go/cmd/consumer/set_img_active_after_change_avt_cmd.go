package consumer

import (
	"github.com/google/uuid"
	sctx "github.com/linhhonblade/service-context"
	"github.com/linhhonblade/service-context/component/gormc"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"social_todo_app_go/common"
	"social_todo_app_go/common/pubsub"
	"social_todo_app_go/component"
	"social_todo_app_go/module/attachment"
)

var SetImgActiveAfterChangeAvtCmd = &cobra.Command{
	Use:   "SetImgActiveAfterChangeAvt",
	Short: "Start consumer: SetImgActiveAfterChangeAvt",
	Run: func(cmd *cobra.Command, args []string) {
		service := sctx.NewServiceContext(
			sctx.WithName("SetImgActiveAfterChangeAvt"),
			sctx.WithComponent(gormc.NewGormDB(common.KeyGormDB, "app")),
			sctx.WithComponent(component.NewNATSComponent(common.KeyNATS)),
		)
		if err := service.Load(); err != nil {
			log.Fatalln(err)
		}
		ps := service.MustGet(common.KeyNATS).(pubsub.PubSub)

		ch, _ := ps.Subscribe(context.Background(), common.TopicUserAvtChanged)
		dbCtx := service.MustGet(common.KeyGormDB).(common.DBContext)
		for msg := range ch {
			mapData := msg.Data()
			imgId := uuid.MustParse(mapData["img_id"].(string))
			repo := attachment.NewRepo(dbCtx.GetDB())
			if err := repo.SetAttachmentStatusActive(context.Background(), imgId); err != nil {
				log.Println(err)
			}
		}

	},
}
