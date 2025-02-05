package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social_todo_app_go/common"
)

type CanGetStatusCode interface {
	GetStatusCode() int
}

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")
				if appErr, ok := err.(*common.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					return
				}
				appErr := common.ErrInternal(err.(error))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err)
				return
			}
		}()
		c.Next()
	}
}

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")
				if appErr, ok := err.(CanGetStatusCode); ok {
					c.AbortWithStatusJSON(appErr.GetStatusCode(), appErr)
				} else {
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
						"code":    http.StatusInternalServerError,
						"status":  "Internal Server Error",
						"message": "Something went wrong"})
				}
				//serviceCtx.Logger("service").Errorf("Panic: %+v \n", err)

				// Must go with gin recovery
				if gin.IsDebugging() {
					panic(err)
				}
			}
		}()
		c.Next()
	}
}
