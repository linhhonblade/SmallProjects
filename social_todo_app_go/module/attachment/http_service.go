package attachment

import (
	"context"
	"github.com/gin-gonic/gin"
	sctx "github.com/linhhonblade/service-context"
	"github.com/linhhonblade/service-context/core"
	"net/http"
	"social_todo_app_go/common"
)

type httpService struct {
	sctx sctx.ServiceContext
}

func NewHttpService(sctx sctx.ServiceContext) httpService {
	return httpService{sctx: sctx}
}

func (s *httpService) handleUploadFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		f, err := c.FormFile("file")
		if err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}
		file, err := f.Open()
		if err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}
		defer file.Close()
		fileData := make([]byte, f.Size)

		if _, err := file.Read(fileData); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}
		dto := UploadDTO{
			Name:     c.PostForm("name"),
			FileName: f.Filename,
			FileType: http.DetectContentType(fileData), // OR use c.Header.Get("Content-Type") but less security
			FileSize: int(f.Size),
			FileData: fileData,
		}

		uploader := s.sctx.MustGet(common.KeyAWSS3).(FileUploader)
		dbContext := s.sctx.MustGet(common.KeyGormDB).(common.DBContext)

		uc := NewUseCase(uploader, NewRepo(dbContext.GetDB()))
		media, err := uc.UploadFile(c.Request.Context(), dto)
		media.SetCDNDomain(uploader.GetDomain())
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}
		c.JSON(http.StatusOK, core.ResponseData(media))

	}
}

func (s httpService) Routes(g *gin.RouterGroup) {
	g.POST("/upload", s.handleUploadFile())
}

type mockAttachmentRepo struct{}

func (mockAttachmentRepo) Create(ctx context.Context, att *Attachment) error {
	//TODO implement me
	return nil
}
