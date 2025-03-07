package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/linhhonblade/service-context/core"
	log "github.com/sirupsen/logrus"
	"social_todo_app_go/common"
	"social_todo_app_go/common/pubsub"
	"social_todo_app_go/module/user/domain"
)

type changeAvtUC struct {
	userQueryRepo  UserQueryRepository
	userCmdRepo    UserCommandRepository
	attachmentRepo AttachmentRepository
}

func NewChangeAvtUC(userQueryRepo UserQueryRepository, userCmdRepo UserCommandRepository, attachmentRepo AttachmentRepository) *changeAvtUC {
	return &changeAvtUC{userQueryRepo: userQueryRepo, userCmdRepo: userCmdRepo, attachmentRepo: attachmentRepo}
}

func (uc *changeAvtUC) ChangeAvt(ctx context.Context, dto SetSingleImageDTO) error {
	//0. Find User
	userEntity, err := uc.userQueryRepo.FindByID(ctx, dto.Requester.UserId())
	if err != nil {
		return core.ErrBadRequest.WithError(domain.ErrCannotChangeAvatar.Error()).WithDebug(err.Error())
	}
	//1. Find Image to make sure it exist in database
	im, err := uc.attachmentRepo.Find(ctx, dto.ImageId)
	if err != nil {
		return core.ErrBadRequest.WithError(domain.ErrCannotChangeAvatar.Error()).WithDebug(err.Error())
	}
	if err := userEntity.SetAvatar(im.FileName); err != nil {
		return core.ErrBadRequest.WithError(domain.ErrCannotChangeAvatar.Error()).WithDebug(err.Error())
	}
	if err := uc.userCmdRepo.Update(ctx, map[string]interface{}{"id": userEntity.Id().String()}, userEntity); err != nil {
		return core.ErrBadRequest.WithError(domain.ErrCannotChangeAvatar.Error()).WithDebug(err.Error())
	}

	//3. Update attachment status to active
	go func() {
		defer common.Recover()
		ps := ctx.Value("pubsub").(pubsub.PubSub)
		if err := ps.Publish(ctx, common.TopicUserAvtChanged, pubsub.NewMessage(map[string]interface{}{
			"user_id": dto.Requester.UserId().String(),
			"img_id":  dto.ImageId.String(),
		})); err != nil {
			log.Println(err)
		}
	}()
	return nil
}

type AttachmentRepository interface {
	Find(ctx context.Context, id uuid.UUID) (*common.Attachment, error)
	SetAttachmentStatusActive(ctx context.Context, id uuid.UUID) error
}
