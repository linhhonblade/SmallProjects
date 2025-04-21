package middleware

import (
	"github.com/gin-gonic/gin"
	"sync/atomic"
)

var counter int64

func CountRequest() func(ctx *gin.Context) {
	return func(context *gin.Context) {
		atomic.AddInt64(&counter, 1)

		context.Next()
	}
}
