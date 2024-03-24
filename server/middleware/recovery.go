package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/common"
)

func Recovery() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(error); ok {
					ctx.AbortWithStatusJSON(http.StatusInternalServerError, common.ErrInternal(err))
				}
				panic(r)
			}
		}()
		ctx.Next()
	}
}
