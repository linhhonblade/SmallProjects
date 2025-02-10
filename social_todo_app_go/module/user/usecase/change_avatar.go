package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/linhhonblade/service-context/core"
	"social_todo_app_go/common"
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
	//1. Find Image to make sure it exist in database
	im, err := uc.attachmentRepo.Find(ctx, dto.ImageId)
	if err != nil {
		return core.ErrBadRequest.WithError(err.Error())
	}

	//2. Set avatar for user
	userId := dto.Requester.UserId()
	if err = uc.userCmdRepo.Update(ctx, map[string]interface{}{"id": userId.String()}, domain.NewUserUpdate(im.Id.String())); err != nil {
		return core.ErrInternalServerError.WithDebug(err.Error())
	}

	//3. Update attachment status to active
	if err = uc.attachmentRepo.SetAttachmentStatusActive(ctx, dto.ImageId); err != nil {
		return core.ErrInternalServerError.WithDebug(err.Error())
	}
	return nil
}

type AttachmentRepository interface {
	Find(ctx context.Context, id uuid.UUID) (*common.Attachment, error)
	SetAttachmentStatusActive(ctx context.Context, id uuid.UUID) error
}
