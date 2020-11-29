package middleware

import (
	"blog-service/pkg/app"
	"blog-service/pkg/errcode"
	"blog-service/pkg/limiter"
	"github.com/gin-gonic/gin"
)

//gin中间件 传入参数 返回就是处理程序
func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response:=app.NewResponse(c)

				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
