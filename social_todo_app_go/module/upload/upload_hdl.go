package upload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"social_todo_app_go/common"
	"time"
)

func Upload(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		dst := fmt.Sprintf("static/%d.%s", time.Now().UTC().UnixNano(), fileHeader.Filename)
		if err != nil {
			c.JSON(400, common.ErrorInvalidRequest(err))
			return
		}
		if err := c.SaveUploadedFile(fileHeader, dst); err != nil {
			c.JSON(400, common.ErrInternal(err))
		}
		img := common.Image{
			Id:        0,
			Url:       dst,
			Width:     100,
			Height:    100,
			CloudName: "local",
			Extension: ""}
		img.Fulfill("http://localhost:3000")
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}
