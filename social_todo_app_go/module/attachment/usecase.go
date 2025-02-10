package attachment

import (
	"context"
	"errors"
	"fmt"
	"github.com/linhhonblade/service-context/core"
	"social_todo_app_go/common"
	"time"
)

type UseCase interface {
	UploadFile(ctx context.Context, dto UploadDTO) (*Attachment, error)
}
type useCase struct {
	uploader FileUploader
	repo     AttachmentCmdRepository
}

func NewUseCase(uploader FileUploader, repo AttachmentCmdRepository) useCase {
	return useCase{uploader: uploader, repo: repo}
}

func (uc *useCase) UploadFile(ctx context.Context, dto UploadDTO) (*Attachment, error) {
	dstFileName := fmt.Sprintf("%d_%s", time.Now().UTC().UnixNano(), dto.FileName)
	if err := uc.uploader.SaveUploadedFile(ctx, dto.FileData, dstFileName); err != nil {
		return nil, core.ErrInternalServerError.WithError(ErrCannotUploadFile.Error()).WithDebug(err.Error())
	}
	now := time.Now().UTC()
	file := Attachment{
		Id:              common.GenUUID(),
		Title:           dto.Name,
		FileName:        dstFileName,
		FileSize:        dto.FileSize,
		FileType:        dto.FileType,
		StorageProvider: uc.uploader.GetName(),
		Status:          StatusUploaded,
		CreatedAt:       now,
		UpdatedAt:       now,
	}
	if err := uc.repo.Create(ctx, &file); err != nil {
		return nil, core.ErrInternalServerError.WithError(ErrCannotUploadFile.Error()).WithDebug(err.Error())
	}
	return &file, nil
}

type FileUploader interface {
	SaveUploadedFile(ctx context.Context, data []byte, dst string) error
	GetName() string
	GetDomain() string
}

type AttachmentCmdRepository interface {
	Create(ctx context.Context, att *Attachment) error
}

type UploadDTO struct {
	Name     string
	FileName string
	FileType string
	FileSize int
	FileData []byte
}

var (
	ErrCannotUploadFile = errors.New("cannot upload file")
	ErrFileNotFound     = errors.New("file not found")
)
