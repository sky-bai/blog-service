package middleware

import (
	"blog-service/global"
	"blog-service/pkg/logger"
	"bytes"
	"github.com/gin-gonic/gin"
	"time"
)

//访问日志写入器
type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

//读响应主体 并放入访问日志结构体中的body里面
func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err//n是p的长度  p是字节数组
	}
	return w.ResponseWriter.Write(p)
}

//访问日志的中间件   获取响应主体 和请求的相关参数
func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer}
		c.Writer = bodyWriter //如何更新一个结构体 自己设置一个特性结构体 再赋值 将gin的writer更新

		beginTime := time.Now().Unix() //unix 获取当前的系统时间
		c.Next()
		endTime := time.Now().Unix()

		//设置日志的字段
		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}
		//设置日志的内容
		global.Logger.WithFields(fields).Infof(c,"access log: method: %s, status_code: %d, begin_time: %d, end_time:%d",
			c.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime, )
	}
}
