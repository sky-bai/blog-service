package middleware

import (
	"blog-service/global"
	"blog-service/pkg/app"
	"blog-service/pkg/email"
	"blog-service/pkg/errcode"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Recovery() gin.HandlerFunc {
	// 创建一个默认的拨号器
	defaultMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})
	return func(c *gin.Context) {
		if err := recover(); err != nil { //如果recovery有错误的话 就处理错误
			global.Logger.WithCallersFrames().Errorf(c, "panic recover err : %v", err)
			err = defaultMailer.SendMail(
				global.EmailSetting.To,
				fmt.Sprintf("异常操作,发生时间: %d", time.Now().Unix()),
				fmt.Sprintf("错误信息: %v", err), )
			if err != nil {
				global.Logger.Panicf(c, "mail.SendMail err: %v", err)
			}
			app.NewResponse(c).ToErrorResponse(errcode.ServerError)
			c.Abort() //阻止挂起程序
		}
		c.Next()
	}
}
