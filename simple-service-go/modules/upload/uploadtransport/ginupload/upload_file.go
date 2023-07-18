package ginupload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-service-go/common"
	"simple-service-go/component"
)

func Upload(appCtx component.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		c.SaveUploadedFile(fileHeader, fmt.Sprintf("./static/filestore/%s", fileHeader.Filename))
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
