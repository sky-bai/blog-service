package middleware

import "github.com/gin-gonic/gin"

//这里设置app的信息
func AppInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app_name","blog_service")//为内容增加kv键值对
		c.Set("app_version","1.0.0")
		c.Next()
	}
}
