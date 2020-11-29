package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

func ContextTimeOut(t time.Duration) func(c *gin.Context) { //持续时间
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()

		c.Request=c.Request.WithContext(ctx)  //context更改为ctx
		c.Next()
	}
}
