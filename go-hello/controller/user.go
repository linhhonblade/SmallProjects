package controller

import (
	"github.com/gin-gonic/gin"
	"go-hello/common"
	"go-hello/model"
	"net/http"
)

func ListUser(context *gin.Context) {
	var filter model.UserNameFilter
	if err := context.ShouldBind(&filter); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var paging common.Paging
	if err := context.ShouldBind(&paging); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	paging.Fulfill()
	result, err := model.ListDataByCondition(nil, &filter, &paging)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
}
