package consumer

import (
	"github.com/google/uuid"
	sctx "github.com/linhhonblade/service-context"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"social_todo_app_go/common"
	"social_todo_app_go/common/asyncjob"
	"social_todo_app_go/common/pubsub"
	"social_todo_app_go/module/attachment"
)

type topicUserChangeAvt struct {
	sctx sctx.ServiceContext
	ps   pubsub.PubSub
}

func NewTopicUserChangeAvt(sctx sctx.ServiceContext, ps pubsub.PubSub) *topicUserChangeAvt {
	return &topicUserChangeAvt{sctx: sctx, ps: ps}
}

func (t *topicUserChangeAvt) handlerSetImgActiveAfterChangeAvt(msg *pubsub.Message) error {
	dbCtx := t.sctx.MustGet(common.KeyGormDB).(common.DBContext)
	mapData := msg.Data()
	imgId := uuid.MustParse(mapData["img_id"].(string))
	repo := attachment.NewRepo(dbCtx.GetDB())
	return repo.SetAttachmentStatusActive(context.Background(), imgId)
}

func (t *topicUserChangeAvt) doAnotherThing(msg *pubsub.Message) error {
	mapData := msg.Data()
	imgId := mapData["img_id"].(string)
	userId := mapData["user_id"].(string)
	log.Println("User %s has changed avatar with image %s", userId, imgId)
	return nil

}

func (t *topicUserChangeAvt) Start() {
	ctx := context.Background()
	ch, _ := t.ps.Subscribe(ctx, common.TopicUserAvtChanged)
	for msg := range ch {
		job1 := asyncjob.NewJob(func(ctx context.Context) error {
			return t.handlerSetImgActiveAfterChangeAvt(msg)
		})
		job2 := asyncjob.NewJob(func(ctx context.Context) error {
			return t.doAnotherThing(msg)
		})
		asyncjob.NewGroup(true, job1, job2).Run(ctx)
	}
}
